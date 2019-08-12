#include <stdio.h>
#include <string.h>
#include <syslog.h>
#include <openssl/des.h>
#include <openssl/ec.h>
#include <openssl/ecdsa.h>
#include <openssl/evp.h>
#include <openssl/bn.h>
#include <openssl/bio.h>
#include <openssl/obj_mac.h>
#include <openssl/pem.h>
#include <openssl/err.h>
#include <openssl/rand.h>

#include "sm/sm2/sm2.h"
#include "sm/sm3/sm3.h"
#include "sm/sm4/sms4.h"
#include "sm/smcryptokit.h"
#include "sm/smcrypto_err.h"

#define SM4_BLOCK_SIZE 16

void __attribute__ ((constructor)) smcrypto_init(void)
{
    ERR_load_crypto_strings();
    ERR_clear_error();
    syslog(LOG_USER|LOG_INFO, "smcrypto_init\n");
}

void __attribute__ ((destructor)) smcrypto_fini(void)
{
    ERR_free_strings();
    syslog(LOG_USER|LOG_INFO, "smcrypto_fini\n");
}

EC_KEY * SM2NewEcKey()
{
    return SM2_gen_key();
}

void SM2FreeEcKey(EC_KEY * ecKey)
{
    EC_KEY_free(ecKey);
}

EC_KEY* LoadSM2PrivKeyFromFile(void* path, int len)
{
    int retCode = 0;
    BIO *key = NULL;
    EVP_PKEY *pkey = NULL;
    EC_KEY *ec_key = NULL;
    char *ppath = NULL;

    SM_ERROR_ESCAPE(path == NULL, SM_F_LOAD_SM2_PRIV_KEY_FROM_FILE, SM_R_INVALID_PARAMETERS, 0);

    ppath = (char*)malloc(len + 1);
    memcpy(ppath, path, len);
    ppath[len] = 0;

    key = BIO_new(BIO_s_file());
    SM_ERROR_ESCAPE(key == NULL, SM_F_LOAD_SM2_PRIV_KEY_FROM_FILE, SM_R_BIO_NEW_FAILED, 0);

    retCode = BIO_read_filename(key, ppath);
    SM_ERROR_ESCAPE(retCode <= 0, SM_F_LOAD_SM2_PRIV_KEY_FROM_FILE, SM_R_BIO_READ_FILENAME_FAILED, retCode);

    pkey = PEM_read_bio_PrivateKey(key, NULL, NULL, NULL);
    SM_ERROR_ESCAPE(pkey == NULL, SM_F_LOAD_SM2_PRIV_KEY_FROM_FILE, SM_R_PEM_READ_BIO_PRIVATEKEY_FAILED, 0);

    ec_key = EVP_PKEY_get1_EC_KEY(pkey);
    SM_ERROR_ESCAPE(ec_key == NULL, SM_F_LOAD_SM2_PRIV_KEY_FROM_FILE, SM_R_EVP_PKEY_GET1_EC_KEY_FAILED, 0);

_err:
    SM_RESOURCE_FREE(ppath, free);
    SM_RESOURCE_FREE(key, BIO_free);
    SM_RESOURCE_FREE(pkey, EVP_PKEY_free);

    return ec_key;
}

EC_KEY * LoadSM2PrivKeyFromBytes(void* keybytes,int len)
{
    int retCode = 0;
    BIO *key = NULL;
    EVP_PKEY *pkey = NULL;
    EC_KEY *ec_key = NULL;

    SM_ERROR_ESCAPE(keybytes == NULL || len <= 135, SM_F_LOAD_SM2_PRIV_KEY_FROM_BYTES, SM_R_INVALID_PARAMETERS, 0);

    key = BIO_new_mem_buf(keybytes, len);
    SM_ERROR_ESCAPE(key == NULL, SM_F_LOAD_SM2_PRIV_KEY_FROM_BYTES, SM_R_BIO_NEW_MEM_FAILED, 0);

    pkey = PEM_read_bio_PrivateKey(key, NULL, NULL, NULL);
    SM_ERROR_ESCAPE(pkey == NULL, SM_F_LOAD_SM2_PRIV_KEY_FROM_BYTES, SM_R_PEM_READ_BIO_PRIVATEKEY_FAILED, 0);

    ec_key = EVP_PKEY_get1_EC_KEY(pkey);
    SM_ERROR_ESCAPE(ec_key == NULL, SM_F_LOAD_SM2_PRIV_KEY_FROM_BYTES, SM_R_EVP_PKEY_GET1_EC_KEY_FAILED, 0);

_err:
    SM_RESOURCE_FREE(key, BIO_free);
    SM_RESOURCE_FREE(pkey, EVP_PKEY_free);

    return ec_key;
}

void SM2FreeX509(X509* x)
{
    X509_free(x);
}

X509* LoadSM2CertFromFile(void* path, int len)
{
    int retCode = 0;
    X509 *x=NULL;
    BIO *cert = NULL;
    char *ppath = NULL;

    SM_ERROR_ESCAPE(path == NULL, SM_F_LOAD_SM2_CERT_FROM_FILE, SM_R_INVALID_PARAMETERS, 0);

    ppath = (char*)malloc(len + 1);
    memcpy(ppath, path, len);
    ppath[len] = 0;

    cert = BIO_new(BIO_s_file());
    SM_ERROR_ESCAPE(cert == NULL, SM_F_LOAD_SM2_CERT_FROM_FILE, SM_R_BIO_NEW_FAILED, 0);

    retCode = BIO_read_filename(cert, ppath);
    SM_ERROR_ESCAPE(retCode <= 0, SM_F_LOAD_SM2_CERT_FROM_FILE, SM_R_BIO_READ_FILENAME_FAILED, retCode);

    x = PEM_read_bio_X509_AUX(cert, NULL, NULL, NULL);
    SM_ERROR_ESCAPE(x == NULL, SM_F_LOAD_SM2_CERT_FROM_FILE, SM_R_PEM_READ_BIO_X509_AUX_FAILED, 0);

_err:
    SM_RESOURCE_FREE(ppath, free);
    SM_RESOURCE_FREE(cert, BIO_free);
    return x;
}

EC_KEY* LoadSM2PubKeyFromBytes(void* keybytes,int len)
{
    EC_KEY *ret = NULL;
    unsigned char *key = (unsigned char*)keybytes;
    
    SM_ERROR_ESCAPE(keybytes == NULL || len <= 65, SM_F_LOAD_SM2_PUB_KEY_FROM_BYTES, SM_R_INVALID_PARAMETERS, 0);

    d2i_EC_PUBKEY(&ret, (const unsigned char**)&key, len);
    SM_ERROR_ESCAPE(ret == NULL, SM_F_LOAD_SM2_PUB_KEY_FROM_BYTES, SM_R_D2I_EC_PUBKEY_FAILED, 0);

_err:
    key = NULL;
    return ret;
}

int SM2Sign(int type, void *dgst, int dLen, void *sig,
            unsigned int *sigLen, void *eckey)
{
    int retCode = 0;

    SM_ERROR_ESCAPE(dgst == NULL || eckey == NULL || sigLen == NULL, SM_F_SM2_SIGN, SM_R_INVALID_PARAMETERS, 0);

    if (sig == NULL) {
        *sigLen = ECDSA_size(eckey);
        return 1;
    }

    retCode = SM2_sign(type, dgst, dLen, sig, sigLen, eckey);
    SM_ERROR_ESCAPE(retCode != 1, SM_F_SM2_SIGN, SM_R_SM2_SIGN_FAILED, retCode);

    retCode = 1;

_err: 
    return retCode;
}

int	SM2SignDirect(int type, void *dgst, int dLen, void *r, int *rLen,
                  void *s, int *sLen, void *eckey)
{
    int retCode = 0;
    ECDSA_SIG *sig = NULL;

    SM_ERROR_ESCAPE(dgst == NULL || eckey == NULL || rLen == NULL || sLen == NULL, SM_F_SM2_SIGN_DIRECT, SM_R_INVALID_PARAMETERS, 0);

    if (r == NULL || s == NULL) {
        *rLen = 32;
        *sLen = 32;
        return 1;
    }

    RAND_seed(dgst, dLen);

    sig = SM2_do_sign(dgst, dLen, eckey);
    SM_ERROR_ESCAPE(sig == NULL, SM_F_SM2_SIGN, SM_R_SM2_SIGN_FAILED, 0);

    *rLen = BN_bn2bin(sig->r, r);
    SM_ERROR_ESCAPE(*rLen <= 0 || *rLen > 32, SM_F_SM2_SIGN_DIRECT, SM_R_SM2_GETBNBYTES_FAILED, 0);
    
    *sLen = BN_bn2bin(sig->s, s);
    SM_ERROR_ESCAPE(*sLen <= 0 || *sLen > 32, SM_F_SM2_SIGN_DIRECT, SM_R_SM2_GETBNBYTES_FAILED, 0);

    retCode = 1;
_err:
    SM_RESOURCE_FREE(sig, ECDSA_SIG_free);

    return retCode;
}

int SM2Verify(int type, void * dgst, int dLen, void *sig, int sigLen, void *eckey)
{
    int retCode = 0;

    SM_ERROR_ESCAPE(dgst == NULL || sig == NULL || eckey == NULL || sigLen <= 0 || sigLen > 72, SM_F_SM2_VERIFY, SM_R_INVALID_PARAMETERS, 0);

    retCode = SM2_verify(type, dgst, dLen, sig, sigLen, eckey);
    SM_ERROR_ESCAPE(retCode != 1, SM_F_SM2_VERIFY, SM_R_SM2_VERIFY_FAILED, retCode);

    retCode = 1;

_err:    
    return retCode;
}

int SM2VerifyDirect(int type, void * dgst, int dLen,
                    void *r, int rLen, void *s, int sLen, void *eckey)
{
    int retCode = 0;
    int derLen = 0;
    ECDSA_SIG *sig = NULL;
    unsigned char *der = NULL;

    SM_ERROR_ESCAPE(dgst == NULL || r == NULL || s == NULL || eckey == NULL || rLen <= 0 || rLen > 32 || sLen <=0 || sLen > 32, SM_F_SM2_VERIFY_DIRECT, SM_R_INVALID_PARAMETERS, 0);

    sig = ECDSA_SIG_new();
    SM_ERROR_ESCAPE(sig == NULL, SM_F_SM2_VERIFY_DIRECT, SM_R_ECDSA_SIG_NEW_FAILED, 0);

    if ((BN_bin2bn(r, rLen, sig->r) == NULL) || (BN_bin2bn(s, sLen, sig->s) == NULL))
    {
        SM_ERROR_ESCAPE(1, SM_F_SM2_VERIFY_DIRECT, SM_R_BN_BIN2BN_FAILED, 0);
    }

    derLen = i2d_ECDSA_SIG(sig, &der);
    SM_ERROR_ESCAPE(derLen <= 0 || derLen > 72, SM_F_SM2_VERIFY_DIRECT, SM_R_I2D_ECDSA_SIG_FAILED, 0);

    retCode = SM2_verify(type, dgst, dLen, der, derLen, eckey);
    SM_ERROR_ESCAPE(retCode != 1, SM_F_SM2_VERIFY_DIRECT, SM_R_SM2_VERIFY_FAILED, retCode);

    retCode = 1;

_err:
    SM_RESOURCE_FREE(der, OPENSSL_free);
    SM_RESOURCE_FREE(sig, ECDSA_SIG_free);
    
    return retCode;
}

//state == sm3_ctx_t.digest   total[0] == sm3_ctx_t.num total[1] == sm3_ctx_t.nblocks
void Sm3Starts(void *total, int totalLen, void *state, int stateLen)
{
    unsigned int *digest = (unsigned int*)state;
    unsigned int *num = (unsigned int*)total;

    SM_ERROR_ESCAPE(totalLen != 2 || stateLen != 8, SM_F_SM3STARTS, SM_R_INVALID_PARAMETERS, 0);

    digest[0] = 0x7380166F;
    digest[1] = 0x4914B2B9;
    digest[2] = 0x172442D7;
    digest[3] = 0xDA8A0600;
    digest[4] = 0xA96F30BC;
    digest[5] = 0x163138AA;
    digest[6] = 0xE38DEE4D;
    digest[7] = 0xB0FB0E4E;

    num[0] = 0;
    num[1] = 0;
_err:
    return;
}

void Sm3Update(void *total, int totalLen, void *state, int stateLen,
               void *buffer, int bufferLen, void *input, int iLen)
{
    unsigned int left = 0;
    unsigned int *digest = (unsigned int*)state;
    unsigned int *num = (unsigned int*)total;
    unsigned char *block = (unsigned char*)buffer;
    unsigned char *data = (unsigned char*)input;

    SM_ERROR_ESCAPE(total == NULL || totalLen !=2 || state == NULL || stateLen != 8 || buffer == NULL || bufferLen != 64 || input == NULL, SM_F_SM3UPDATE, SM_R_INVALID_PARAMETERS, 0);

    if (num[0]) {
        left = SM3_BLOCK_SIZE - num[0];
        if ((unsigned int)iLen < left) {
            memcpy(block + num[0], data, iLen);
            num[0] += iLen;
            return;
        } else {
            memcpy(block + num[0], data, left);
            sm3_compress(digest, block);
            num[1]++;
            data += left;
            iLen -= left;
        }
    }
    while (iLen >= SM3_BLOCK_SIZE) {
        sm3_compress(digest, data);
        num[1]++;
        data += SM3_BLOCK_SIZE;
        iLen -= SM3_BLOCK_SIZE;
    }
    num[0] = iLen;
    if (iLen) {
        memcpy(block, data, iLen);
    }

_err:
    return;
}

void Sm3Finish(void *total, int totalLen, void *state, int stateLen,
               void *buffer, int bufferLen, void *output, int oLen)
{
    int i;
    unsigned int *digest = (unsigned int*)state;
    unsigned int *num = (unsigned int*)total;
    unsigned char *block = (unsigned char*)buffer;
    unsigned int *pdigest = (uint32_t *)output;
    unsigned int *count = NULL;

    SM_ERROR_ESCAPE(total == NULL || totalLen !=2 || state == NULL || stateLen != 8 || buffer == NULL || bufferLen != 64 || output == NULL || oLen != 32, SM_F_SM3FINISH, SM_R_INVALID_PARAMETERS, 0);

    count = (unsigned int*)(block + SM3_BLOCK_SIZE - 8);;

    block[num[0]] = 0x80;

    if (num[0] + 9 <= SM3_BLOCK_SIZE) {
        memset(block + num[0] + 1, 0, SM3_BLOCK_SIZE - num[0] - 9);
    } else {
        memset(block + num[0] + 1, 0, SM3_BLOCK_SIZE - num[0] - 1);
        sm3_compress(digest, block);
        memset(block, 0, SM3_BLOCK_SIZE - 8);
    }

    count[0] = cpu_to_be32((num[1]) >> 23);
    count[1] = cpu_to_be32((num[1] << 9) + (num[0] << 3));

    sm3_compress(digest, block);
    for (i = 0; i < stateLen; i++) {
        pdigest[i] = cpu_to_be32(digest[i]);
    }

_err:
    return;
}

void sm4_setkey_enc(void * encKey, void * key, int keyLen)
{
    SM_ERROR_ESCAPE(encKey == NULL || key == NULL || keyLen != 16, SM_F_SM4_SETKEY_ENC, SM_R_INVALID_PARAMETERS, 0);

    sms4_set_encrypt_key(encKey, key);

_err:
    return;
}

void sm4_setkey_dec(void * decKey, void * key, int keyLen)
{
    SM_ERROR_ESCAPE(decKey == NULL || key == NULL || keyLen != 16, SM_F_SM4_SETKEY_DEC, SM_R_INVALID_PARAMETERS, 0);

    sms4_set_decrypt_key(decKey, key);

_err:
    return;
}

void sm4_crypt_ecb(void * seckey, int mode, int length, void *in, void *out)
{
    int blLen = 0;
    unsigned char *input = NULL;
    unsigned char *output = NULL;

    SM_ERROR_ESCAPE(seckey == NULL || (length % 16) != 0 || in == NULL || out == NULL, SM_F_SM4_CRYPT_ECB, SM_R_INVALID_PARAMETERS, 0);

    blLen = length;
    input = in;
    output = out;

    while(blLen > 0)
    {
        sms4_encrypt(input, output, (sms4_key_t *)seckey);

        blLen -= SM4_BLOCK_SIZE;
        input += SM4_BLOCK_SIZE;
        output += SM4_BLOCK_SIZE;
    }

_err:
    return;
}

void sm4_crypt_cbc(void * seckey, int mode, int length, void * initVec, void * in, void * out)
{
    SM_ERROR_ESCAPE(seckey == NULL || (length % 16) != 0 || initVec == NULL || in == NULL || out == NULL, SM_F_SM4_CRYPT_ECB, SM_R_INVALID_PARAMETERS, 0);

    sms4_cbc_encrypt(in, out, length, seckey, initVec, mode);

_err:
    return;
}
