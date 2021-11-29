package utils

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"github.com/pkg/errors"
)

//AddCountryCodePhoneNumber replace First Digit Phone Number With Country Code
func AddCountryCodePhoneNumber(countryCode string, phoneNumber string) (*string, error) {
	if countryCode == "" || phoneNumber == "" {
		return nil, errors.New("parameters error")
	}

	num, err := phonenumbers.Parse(phoneNumber, countryCode)
	if err != nil {
		return nil, err
	}
	if num.CountryCode == nil || num.NationalNumber == nil {
		return nil, errors.New("invalid phone number")
	}

	strCountryCode := fmt.Sprint(*num.CountryCode)
	strNationalNumber := fmt.Sprint(*num.NationalNumber)

	result := fmt.Sprintf("%s%s%s", "+", strCountryCode, strNationalNumber)

	return &result, nil
}
