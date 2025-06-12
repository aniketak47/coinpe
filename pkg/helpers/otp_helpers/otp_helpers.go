package otphelpers

import (
	"coinpe/pkg/constants"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateOTP(secret string, shouldMock bool) (string, error) {
	otp, err := totp.GenerateCode(secret, time.Now())
	if shouldMock {
		otp = constants.MockOTP
	}

	return otp, err
}

func ValidateOTP(secret string, receivedOTP string, shouldMock bool) (valid bool, err error) {
	if shouldMock {
		valid, err = (receivedOTP == constants.MockOTP), nil
	} else {
		valid, err = totp.ValidateCustom(receivedOTP, secret, time.Now(), totp.ValidateOpts{Skew: 15, Digits: otp.DigitsSix})
	}
	return
}
