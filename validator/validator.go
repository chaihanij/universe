package validator

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pariz/gountries"
	"github.com/shopspring/decimal"
	"gitlab.com/gridwhizth/universe/constants"
	"gitlab.com/gridwhizth/universe/utils"
	_currency "golang.org/x/text/currency"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func IsValidEmail(email string) bool {
	return regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`).
		MatchString(email)
}

func IsValidPhoneNumber(phoneNumber string) bool {
	return regexp.MustCompile(`^[+]{0,1}[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`).
		MatchString(phoneNumber)
}

func IsValidUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

func IsValidSlug(slug string) bool {
	return regexp.
		MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`). // lowercase kebab case string
		MatchString(slug)
}

func IsValidCurrency(currency string) bool {
	if _, err := _currency.ParseISO(currency); err != nil {
		return false
	}
	return regexp.
		MustCompile(`^[A-Z]+(?:-[A-Z]+)*$`).
		MatchString(currency) && len(currency) == 3
}

func IsValidCountry(country string) bool {
	ok := regexp.
		MustCompile(`^[A-Z]+(?:-[A-Z]+)*$`).
		MatchString(country) && len(country) == 3
	if !ok {
		return false
	}

	_, err := gountries.New().FindCountryByAlpha(country)
	if err != nil {
		return false
	}
	return true
}

func IsValidNumericFromString(decimalOfString string) bool {
	_, err := decimal.NewFromString(decimalOfString)
	if err != nil {
		return false
	}
	return true
}

func ValidatePasswordComplexity(password string) bool {
	// Must have at least one lower case
	isValid := regexp.
		MustCompile(`[a-z]`).
		MatchString(password)
	if !isValid {
		return false
	}

	// Must have at least one upper case
	isValid = regexp.
		MustCompile(`[A-Z]`).
		MatchString(password)
	if !isValid {
		return false
	}

	// Must have at least one number
	isValid = regexp.
		MustCompile(`\d`).
		MatchString(password)
	if !isValid {
		return false
	}

	// Must have at least one symbol
	isValid = regexp.
		MustCompile(`[-+_!@#$%^&*.,?]`).
		MatchString(password)
	if !isValid {
		return false
	}

	return true
}

func IsValidBoolFromString(boolOfString string) bool {
	_, err := strconv.ParseBool(boolOfString)
	if err != nil {
		return false
	}
	return true
}

func IsValidDateTimeFromString(layoutOfDateTime string, dateTimeOfString string) bool {
	_, err := time.Parse(layoutOfDateTime, dateTimeOfString)
	if err != nil {
		return false
	}
	return true
}

func IsWeakPin6Digit(pin string) bool {
	arrWeakPin := utils.GetWeakPin6Digit()
	for _, code := range arrWeakPin {
		if pin == code {
			return true
		}
	}
	return false
}

func IsValidUsername(username string) bool {
	// Must more than 6 digit
	if len(username) < 6 {
		return false
	}

	isValid := regexp.
		MustCompile(`^[a-z0-9]+(?:[a-z0-9]+)*$`).
		MatchString(username)
	if !isValid {
		return false
	}

	// check duplicate characters in string
	runes := []rune(username)
	for _, text := range runes {
		textCheck := string(text) + string(text) + string(text) + string(text) + string(text)
		if strings.Count(username, textCheck) >= 1 {
			return false
		}
	}

	return true
}

func IsJSON(str string) bool {
	var js interface{}
	if err := json.Unmarshal([]byte(str), &js); err != nil {
		return false
	}
	switch js.(type) {
	case map[string]interface{}:
		return true
	case []interface{}:
		return true
	default:
		return false
	}
}

func IsBase64(str string) bool {
	return constants.RxBase64.MatchString(str)
}

func IsBase64DataType(str string) bool {
	return constants.RxBase64DataType.MatchString(str)
}

func IsURL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return constants.RxURL.MatchString(str)

}
