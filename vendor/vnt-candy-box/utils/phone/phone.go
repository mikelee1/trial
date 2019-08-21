package phone

import (
	"fmt"
	"github.com/ttacon/libphonenumber"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 13912345678, CN => +8613912345678
func ParseToInternational(phone string, country string) (string, error) {
	if phone == "" || country == "" {
		return "", fmt.Errorf("empty phonenumber [%s] or country code [%s]", phone, country)
	}
	num, err := libphonenumber.Parse(phone, country)
	if err != nil {
		return "", fmt.Errorf("error parsing phone [%s] with country code [%s]", phone, country)
	}
	return strings.Replace(libphonenumber.Format(num, libphonenumber.INTERNATIONAL), " ", "", -1), nil
}

// +8613912345678, CN => 13912345678
func ParseToLocal(phone string, country string) (string, error) {
	if phone == "" || country == "" {
		return "", fmt.Errorf("empty international phonenumber [%s] or country code [%s]", phone, country)
	}
	pn := &libphonenumber.PhoneNumber{}
	if err := libphonenumber.ParseToNumber(phone, country, pn); err != nil {
		return "", fmt.Errorf("error parsing international number [%s] to the local with country code [%s]", phone, country)
	}
	return strconv.FormatUint(*pn.NationalNumber, 10), nil
}
func GenerateVerifyCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(100000))
}
