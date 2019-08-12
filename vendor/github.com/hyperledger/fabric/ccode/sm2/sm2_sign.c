/* ====================================================================
 * Copyright (c) 2015 - 2016 The GmSSL Project.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in
 *    the documentation and/or other materials provided with the
 *    distribution.
 *
 * 3. All advertising materials mentioning features or use of this
 *    software must display the following acknowledgment:
 *    "This product includes software developed by the GmSSL Project.
 *    (http://gmssl.org/)"
 *
 * 4. The name "GmSSL Project" must not be used to endorse or promote
 *    products derived from this software without prior written
 *    permission. For written permission, please contact
 *    guanzhi1980@gmail.com.
 *
 * 5. Products derived from this software may not be called "GmSSL"
 *    nor may "GmSSL" appear in their names without prior written
 *    permission of the GmSSL Project.
 *
 * 6. Redistributions of any form whatsoever must retain the following
 *    acknowledgment:
 *    "This product includes software developed by the GmSSL Project
 *    (http://gmssl.org/)"
 *
 * THIS SOFTWARE IS PROVIDED BY THE GmSSL PROJECT ``AS IS'' AND ANY
 * EXPRESSED OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE GmSSL PROJECT OR
 * ITS CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
 * NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
 * OF THE POSSIBILITY OF SUCH DAMAGE.
 * ====================================================================
 */

#include <string.h>
#include <openssl/bn.h>
#include <openssl/ec.h>
#include <openssl/err.h>
#include <openssl/evp.h>
#include <openssl/rand.h>
#include <openssl/obj_mac.h>

//#include "../smcrypto_err.h"
//#include "sm2.h"
#include "sm/smcrypto_err.h"
#include "sm/sm2/sm2.h"


static int sm2_sign_setup(EC_KEY *ec_key, BN_CTX *ctx_in, BIGNUM **kp, BIGNUM **xp)
{
    int ret = 0;
    const EC_GROUP *ec_group;
    BN_CTX *ctx = NULL;
    BIGNUM *k = NULL;
    BIGNUM *x = NULL;
    BIGNUM *order = NULL;
    EC_POINT *point = NULL;

    SM_ERROR_ESCAPE(ec_key == NULL, SM_F_SM2_SIGN_SETUP, SM_R_INVALID_PARAMETERS, 0);

    ec_group = EC_KEY_get0_group(ec_key);
    SM_ERROR_ESCAPE(ec_group == NULL, SM_F_SM2_SIGN_SETUP, SM_R_EC_KEY_GET0_GROUP_FAILED, 0);

    if (ctx_in == NULL)  {
        ctx = BN_CTX_new();
        SM_ERROR_ESCAPE(ctx == NULL, SM_F_SM2_SIGN_SETUP, SM_R_BN_CTX_NEW_FAILED, 0);
    }
    else {
        ctx = ctx_in;
    }

    k = BN_new();
    x = BN_new();
    order = BN_new();
    SM_ERROR_ESCAPE(k == NULL || x == NULL || order == NULL, SM_F_SM2_SIGN_SETUP, SM_R_BN_NEW_FAILED, 0);

    ret = EC_GROUP_get_order(ec_group, order, ctx);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_SIGN_SETUP, SM_R_EC_GROUP_GET_ORDER_FAILED, 0);
    
    point = EC_POINT_new(ec_group);
    SM_ERROR_ESCAPE(point == NULL, SM_F_SM2_SIGN_SETUP, SM_R_EC_POINT_NEW_FAILED, 0);

    do {
        /* get random k */
        do {
            ret = BN_rand_range(k, order);
            SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_SIGN_SETUP, SM_R_BN_RAND_RANGE_FAILED, 0);
        } while (BN_is_zero(k));

        /* compute r the x-coordinate of generator * k */
        ret = EC_POINT_mul(ec_group, point, k, NULL, NULL, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_SIGN_SETUP, SM_R_EC_POINT_MUL_FAILED, 0);

        if (EC_METHOD_get_field_type(EC_GROUP_method_of(ec_group)) == NID_X9_62_prime_field) {
            ret = EC_POINT_get_affine_coordinates_GFp(ec_group, point, x, NULL, ctx);
            SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_SIGN_SETUP, SM_R_EC_POINT_GET_AFFINE_COORDINATES_GFP_FAILED, 0);
        } else /* NID_X9_62_characteristic_two_field */ {
            ret = EC_POINT_get_affine_coordinates_GF2m(ec_group, point, x, NULL, ctx);
            SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_SIGN_SETUP, SM_R_EC_POINT_GET_AFFINE_COORDINATES_GF2M_FAILED, 0);
        }

        ret = BN_nnmod(x, x, order, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_SIGN_SETUP, SM_R_BN_NNMOD_FAILED, 0);
    } while (BN_is_zero(x));

    /* clear old values if necessary */
    BN_clear_free(*kp);
    BN_clear_free(*xp);

    /* save the pre-computed values  */
    *kp = k;
    *xp = x;
    ret = 1;

_err:
    if (!ret) {
        SM_RESOURCE_FREE(k, BN_clear_free);
        SM_RESOURCE_FREE(x, BN_clear_free);
    }
    if (!ctx_in) {
        SM_RESOURCE_FREE(ctx,BN_CTX_free);
    }
    SM_RESOURCE_FREE(order, BN_free);
    SM_RESOURCE_FREE(point, EC_POINT_free);

    return(ret);
}

static ECDSA_SIG *sm2_do_sign(const unsigned char *dgst, int dgstlen,
    const BIGNUM *in_k, const BIGNUM *in_x, EC_KEY *ec_key)
{
    int ret = 0;
    ECDSA_SIG *ec_sig = NULL;
    const EC_GROUP *ec_group;
    const BIGNUM *priv_key;
    const BIGNUM *ck;
    BIGNUM *k = NULL;
    BN_CTX *ctx = NULL;
    BIGNUM *order = NULL;
    BIGNUM *e = NULL;
    BIGNUM *bn = NULL;
    BIGNUM *bn_ret = NULL;

    SM_ERROR_ESCAPE(ec_key == NULL, SM_F_SM2_DO_SIGN, SM_R_INVALID_PARAMETERS, 0);

    ec_group = EC_KEY_get0_group(ec_key);
    SM_ERROR_ESCAPE(ec_group == NULL, SM_F_SM2_DO_SIGN, SM_R_EC_KEY_GET0_GROUP_FAILED, 0);

    priv_key = EC_KEY_get0_private_key(ec_key);
    SM_ERROR_ESCAPE(priv_key == NULL, SM_F_SM2_DO_SIGN, SM_R_EC_KEY_GET0_PRIVATE_KEY_FAILED, 0);

    ec_sig = ECDSA_SIG_new();
    SM_ERROR_ESCAPE(ec_sig == NULL, SM_F_SM2_DO_SIGN, SM_R_ECDSA_SIG_NEW_FAILED, 0);

#ifdef OPENSSL_VER_1_1_0
    ec_sig->r = BN_new();
    ec_sig->s = BN_new();
    SM_ERROR_ESCAPE(ec_sig->r == NULL || ec_sig->s == NULL, SM_F_SM2_DO_SIGN, SM_R_BN_NEW_FAILED, 0);
#endif
    ctx = BN_CTX_new();
    SM_ERROR_ESCAPE(ctx == NULL, SM_F_SM2_DO_SIGN, SM_R_BN_CTX_NEW_FAILED, 0);

    bn = BN_new();
    order = BN_new();
    SM_ERROR_ESCAPE(order == NULL || bn == NULL, SM_F_SM2_DO_SIGN, SM_R_BN_NEW_FAILED, 0);

    ret = EC_GROUP_get_order(ec_group, order, ctx);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_EC_GROUP_GET_ORDER_FAILED, 0);

    /* convert dgst to e */
    e = BN_bin2bn(dgst, dgstlen, NULL);
    SM_ERROR_ESCAPE(e == NULL, SM_F_SM2_DO_SIGN, SM_R_BN_BIN2BN_FAILED, 0);

    do {
        /* use or compute k and (kG).x */
        if (!in_k || !in_x) {
            ret = sm2_sign_setup(ec_key, ctx, &k, &ec_sig->r);
            SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_SM2_SIGN_SETUP_FAILED, 0);

            ck = k;
        } else {
            ck = in_k;
            bn_ret = BN_copy(ec_sig->r, in_x);
            SM_ERROR_ESCAPE(bn_ret == NULL, SM_F_SM2_DO_SIGN, SM_R_BN_COPY_FAILED, 0);
        }

        /* r = e + x (mod n) */
        ret = BN_mod_add(ec_sig->r, ec_sig->r, e, order, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_BN_MOD_ADD_FAILED, 0);

        ret = BN_mod_add(bn, ec_sig->r, ck, order, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_BN_MOD_ADD_FAILED, 0);

        /* check r != 0 && r + k != n */
        if (BN_is_zero(ec_sig->r) || BN_is_zero(bn)) {
            SM_ERROR_ESCAPE( in_k && in_x, SM_F_SM2_DO_SIGN, SM_R_NEED_NEW_SETUP_VALUES, 0);

            continue;
        }

        /* s = ((1 + d)^-1 * (k - rd)) mod n */
        ret = BN_one(bn);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_BN_ONE_FAILED, 0);

        ret = BN_mod_add(ec_sig->s, priv_key, bn, order, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_BN_MOD_ADD_FAILED, 0);

        bn_ret = BN_mod_inverse(ec_sig->s, ec_sig->s, order, ctx);
        SM_ERROR_ESCAPE(bn_ret == NULL, SM_F_SM2_DO_SIGN, SM_R_BN_MOD_INVERSE_FAILED, 0);

        ret = BN_mod_mul(bn, ec_sig->r, priv_key, order, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_BN_MOD_MUL_FAILED, 0);

        ret = BN_mod_sub(bn, ck, bn, order, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_BN_MOD_SUB_FAILED, 0);

        ret = BN_mod_mul(ec_sig->s, ec_sig->s, bn, order, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_SIGN, SM_R_BN_MOD_MUL_FAILED, 0);

        /* check s != 0 */
        if (BN_is_zero(ec_sig->s)) {
            SM_ERROR_ESCAPE( in_k && in_x, SM_F_SM2_DO_SIGN, SM_R_NEED_NEW_SETUP_VALUES, 0);
        } else {
            break;
        }

    } while (1);

    ret = 1;

_err:
    if (!ret) {
        SM_RESOURCE_FREE(ec_sig, ECDSA_SIG_free);
    }
    SM_RESOURCE_FREE(k, BN_free);
    SM_RESOURCE_FREE(e, BN_free);
    SM_RESOURCE_FREE(bn, BN_free);
    SM_RESOURCE_FREE(order, BN_free);
    SM_RESOURCE_FREE(ctx, BN_CTX_free);

    return ec_sig;
}

int sm2_do_verify(const unsigned char *dgst, int dgstlen,
    const ECDSA_SIG *sig, EC_KEY *ec_key)
{
    int ret = -1;
    const EC_GROUP *ec_group;
    const EC_POINT *pub_key;
    EC_POINT *point = NULL;
    BN_CTX *ctx = NULL;
    BIGNUM *order = NULL;
    BIGNUM *e = NULL;
    BIGNUM *t = NULL;

    SM_ERROR_ESCAPE(sig == NULL || ec_key == NULL, SM_F_SM2_DO_VERIFY, SM_R_INVALID_PARAMETERS, 0);

    ec_group = EC_KEY_get0_group(ec_key);
    SM_ERROR_ESCAPE(ec_group == NULL, SM_F_SM2_DO_VERIFY, SM_R_EC_KEY_GET0_GROUP_FAILED, 0);

    pub_key  = EC_KEY_get0_public_key(ec_key);
    SM_ERROR_ESCAPE(pub_key == NULL, SM_F_SM2_DO_VERIFY, SM_R_EC_KEY_GET0_PUBLIC_KEY_FAILED, 0);

    ctx = BN_CTX_new();
    SM_ERROR_ESCAPE(ctx == NULL, SM_F_SM2_DO_VERIFY, SM_R_BN_CTX_NEW_FAILED, 0);

    order = BN_new();
    t = BN_new();
    SM_ERROR_ESCAPE(order == NULL ||  t == NULL, SM_F_SM2_DO_VERIFY, SM_R_BN_NEW_FAILED, 0);

    ret = EC_GROUP_get_order(ec_group, order, ctx);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_VERIFY, SM_R_EC_GROUP_GET_ORDER_FAILED, 0);

    /* check r, s in [1, n-1] and r + s != 0 (mod n) */
    if (BN_is_zero(sig->r) ||
        BN_is_negative(sig->r) ||
        BN_ucmp(sig->r, order) >= 0 ||
        BN_is_zero(sig->s) ||
        BN_is_negative(sig->s) ||
        BN_ucmp(sig->s, order) >= 0) {
            SM_ERROR_ESCAPE(1, SM_F_SM2_DO_VERIFY, SM_R_BAD_SIGNATURE, 0);
    }

    /* check t = r + s != 0 */
    ret = BN_mod_add(t, sig->r, sig->s, order, ctx);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_VERIFY, SM_R_BN_MOD_ADD_FAILED, 0);

    if (BN_is_zero(t)) {
        SM_ERROR_ESCAPE(1, SM_F_SM2_DO_VERIFY, SM_R_BN_IS_ZERO_FAILED, 0);
    }

    /* convert digest to e */
    e = BN_bin2bn(dgst, dgstlen, NULL);
    SM_ERROR_ESCAPE(e == NULL, SM_F_SM2_DO_VERIFY, SM_R_BN_BIN2BN_FAILED, 0);

    /* compute (x, y) = sG + tP, P is pub_key */
    point = EC_POINT_new(ec_group);
    SM_ERROR_ESCAPE(point == NULL, SM_F_SM2_DO_VERIFY, SM_R_EC_POINT_NEW_FAILED, 0);

    ret = EC_POINT_mul(ec_group, point, sig->s, pub_key, t, ctx);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_VERIFY, SM_R_EC_POINT_MUL_FAILED, 0);

    if (EC_METHOD_get_field_type(EC_GROUP_method_of(ec_group)) == NID_X9_62_prime_field) {
        ret = EC_POINT_get_affine_coordinates_GFp(ec_group, point, t, NULL, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_VERIFY, SM_R_EC_POINT_GET_AFFINE_COORDINATES_GFP_FAILED, 0);
    } else /* NID_X9_62_characteristic_two_field */ {
        ret = EC_POINT_get_affine_coordinates_GF2m(ec_group, point, t, NULL, ctx);
        SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_VERIFY, SM_R_EC_POINT_GET_AFFINE_COORDINATES_GF2M_FAILED, 0);
    }
    ret = BN_nnmod(t, t, order, ctx);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_VERIFY, SM_R_BN_NNMOD_FAILED, 0);

    /* check (sG + tP).x + e  == sig.r */
    ret = BN_mod_add(t, t, e, order, ctx);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_DO_VERIFY, SM_R_BN_MOD_ADD_FAILED, 0);

    if (BN_ucmp(t, sig->r) == 0) {
        ret = 1;
    } else {
        ret = 0;
    }

_err:
    SM_RESOURCE_FREE(e, BN_free);
    SM_RESOURCE_FREE(t, BN_free);
    SM_RESOURCE_FREE(order, BN_free);
    SM_RESOURCE_FREE(ctx, BN_CTX_free);
    SM_RESOURCE_FREE(point, EC_POINT_free);

    return ret;
}

int SM2_sign_setup(EC_KEY *ec_key, BN_CTX *ctx_in, BIGNUM **kp, BIGNUM **xp)
{
    return sm2_sign_setup(ec_key, ctx_in, kp, xp);
}

ECDSA_SIG *SM2_do_sign_ex(const unsigned char *dgst, int dgstlen,
    const BIGNUM *kp, const BIGNUM *xp, EC_KEY *ec_key)
{
    return sm2_do_sign(dgst, dgstlen, kp, xp, ec_key);
}

ECDSA_SIG *SM2_do_sign(const unsigned char *dgst, int dgstlen, EC_KEY *ec_key)
{
    return SM2_do_sign_ex(dgst, dgstlen, NULL, NULL, ec_key);
}

int SM2_do_verify(const unsigned char *dgst, int dgstlen,
    const ECDSA_SIG *sig, EC_KEY *ec_key)
{
    return sm2_do_verify(dgst, dgstlen, sig, ec_key);
}

int SM2_sign_ex(int type, const unsigned char *dgst, int dgstlen,
    unsigned char *sig, unsigned int *siglen,
    const BIGNUM *k, const BIGNUM *x, EC_KEY *ec_key)
{
    int ret = 0;
    ECDSA_SIG *s = NULL;

    RAND_seed(dgst, dgstlen);

    s = SM2_do_sign_ex(dgst, dgstlen, k, x, ec_key);
    SM_ERROR_ESCAPE(s == NULL, SM_F_SM2_SIGN_EX, SM_R_SM2_DO_SIGN_EX_FAILED, 0);

    ret = i2d_ECDSA_SIG(s, &sig);
    SM_ERROR_ESCAPE(ret <= 0, SM_F_SM2_SIGN_EX, SM_R_I2D_ECDSA_SIG_FAILED, 0);

    *siglen = ret;

    ret = 1;

_err:
    SM_RESOURCE_FREE(s, ECDSA_SIG_free);
    return ret;
}

int SM2_sign(int type, const unsigned char *dgst, int dgstlen,
    unsigned char *sig, unsigned int *siglen, EC_KEY *ec_key)
{
    return SM2_sign_ex(type, dgst, dgstlen, sig, siglen, NULL, NULL, ec_key);
}

int SM2_verify(int type, const unsigned char *dgst, int dgstlen,
    const unsigned char *sig, int siglen, EC_KEY *ec_key)
{
    ECDSA_SIG *s = NULL;
    const unsigned char *p = sig;
    unsigned char *der = NULL;
    int derlen = -1;
    int ret = -1;

    s = ECDSA_SIG_new();
    SM_ERROR_ESCAPE(s == NULL, SM_F_SM2_VERIFY, SM_R_ECDSA_SIG_NEW_FAILED, 0);

    if (!d2i_ECDSA_SIG(&s, &p, siglen)) {
        SM_ERROR_ESCAPE(1, SM_F_SM2_VERIFY, SM_R_D2I_ECDSA_SIG_FAILED, 0);
    }

    derlen = i2d_ECDSA_SIG(s, &der);
    if (derlen != siglen || memcmp(sig, der, derlen)) {
        SM_ERROR_ESCAPE(1, SM_F_SM2_VERIFY, SM_R_I2D_ECDSA_SIG_FAILED, 0);
    }

    ret = SM2_do_verify(dgst, dgstlen, s, ec_key);
    SM_ERROR_ESCAPE(ret == 0, SM_F_SM2_VERIFY, SM_R_SM2_DO_VERIFY_FAILED, 0);

    ret = 1;

_err:
    if (derlen > 0) {
        OPENSSL_cleanse(der, derlen);
        SM_RESOURCE_FREE(der, OPENSSL_free);
    }

    SM_RESOURCE_FREE(s, ECDSA_SIG_free);

    return ret;
}
