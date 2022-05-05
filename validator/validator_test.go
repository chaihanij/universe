package validator_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/gridwhizth/universe/utils"
	"gitlab.com/gridwhizth/universe/validator"
	"testing"
	"time"
)

func TestUtils_IsValidEmail(t *testing.T) {
	t.Run("Email is valid", func(t *testing.T) {
		assert.True(t, validator.IsValidEmail("example@gmail.com"))
		assert.True(t, validator.IsValidEmail("example123@gmail.com"))
		assert.True(t, validator.IsValidEmail("example-ex@gmail.com"))
		assert.True(t, validator.IsValidEmail("example_ex@gmail.com"))
	})

	t.Run("Email is invalid", func(t *testing.T) {
		assert.False(t, validator.IsValidEmail("email"))
		assert.False(t, validator.IsValidEmail("example@"))
		assert.False(t, validator.IsValidEmail("example@ex"))
		assert.False(t, validator.IsValidEmail("example@example"))
	})
}

func TestUtils_IsValidPhoneNumber(t *testing.T) {
	t.Run("Phone number is valid", func(t *testing.T) {
		assert.True(t, validator.IsValidPhoneNumber("+0833454345"))
		assert.True(t, validator.IsValidPhoneNumber("089-1234567"))
		assert.True(t, validator.IsValidPhoneNumber("091-933-0998"))
	})

	t.Run("Phone number is invalid", func(t *testing.T) {
		assert.False(t, validator.IsValidPhoneNumber("f+898448331"))
		assert.False(t, validator.IsValidPhoneNumber("++09334454433"))
		assert.False(t, validator.IsValidPhoneNumber("0987687768f"))
		assert.False(t, validator.IsValidPhoneNumber("abcdefghijklkmnop"))
	})
}

func TestUtils_IsValidUuid(t *testing.T) {
	t.Run("Happy UUID is valid", func(t *testing.T) {
		assert.True(t, validator.IsValidUUID("fbd3036f-0f1c-4e98-b71c-d4cd61213f90"))
	})

	t.Run("Happy UUID is invalid", func(t *testing.T) {
		assert.False(t, validator.IsValidUUID("bd3036f-0f1c-4e98-b71c-d4cd61213f90"))
		assert.False(t, validator.IsValidUUID("cccc-ccc-ccc-ccc-ccccc"))
		assert.False(t, validator.IsValidUUID(""))
		assert.False(t, validator.IsValidUUID("2323423423423423"))
	})
}

func TestUtils_IsValidSlug(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		assert.True(t, validator.IsValidSlug("ndigital-test"))
		assert.False(t, validator.IsValidSlug("Ndigital-test"))
		assert.False(t, validator.IsValidSlug("*778#$$"))
	})
}

func Test_IsValidCurrency(t *testing.T) {

	validCurrencyList := []string{"USD", "THB", "ARS"}
	invalidCurrencyList := []string{"ABC", "USDA", "usd", "123"}

	t.Run("Happy", func(t *testing.T) {
		for _, currency := range validCurrencyList {
			assert.True(t, validator.IsValidCurrency(currency), "This currency should pass validation: %s", currency)
		}
	})
	t.Run("Fail", func(t *testing.T) {
		for _, currency := range invalidCurrencyList {
			assert.False(t, validator.IsValidCurrency(currency), "This currency should fail validation: %s", currency)
		}
	})
}

func TestIsValidCountry(t *testing.T) {
	validCountryList := []string{"THA", "USA", "KHM"}
	invalidCountryList := []string{"THB", "usa", "usd", "kHm", "456"}

	t.Run("Happy", func(t *testing.T) {
		for _, country := range validCountryList {
			assert.True(t, validator.IsValidCountry(country), "This country should pass validation: %s", country)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		for _, country := range invalidCountryList {
			assert.False(t, validator.IsValidCountry(country), "This country should fail validation: %s", country)
		}
	})

}

func TestIsValidNumericFromString(t *testing.T) {
	badDecimalCase1 := ""
	badDecimalCase2 := "-"
	badDecimalCase3 := "abcd"
	badDecimalCase4 := "1. 0"

	mockedDecimalCase1 := "1"
	mockedDecimalCase2 := "10.00"
	mockedDecimalCase3 := "0.1"
	mockedDecimalCase4 := "-1"
	mockedDecimalCase5 := "-1.00"

	t.Run("Happy", func(t *testing.T) {
		isDecimalCase1 := validator.IsValidNumericFromString(mockedDecimalCase1)
		isDecimalCase2 := validator.IsValidNumericFromString(mockedDecimalCase2)
		isDecimalCase3 := validator.IsValidNumericFromString(mockedDecimalCase3)
		isDecimalCase4 := validator.IsValidNumericFromString(mockedDecimalCase4)
		isDecimalCase5 := validator.IsValidNumericFromString(mockedDecimalCase5)

		assert.Equal(t, true, isDecimalCase1)
		assert.Equal(t, true, isDecimalCase2)
		assert.Equal(t, true, isDecimalCase3)
		assert.Equal(t, true, isDecimalCase4)
		assert.Equal(t, true, isDecimalCase5)
	})

	t.Run("Fail", func(t *testing.T) {
		isDecimalCase1 := validator.IsValidNumericFromString(badDecimalCase1)
		isDecimalCase2 := validator.IsValidNumericFromString(badDecimalCase2)
		isDecimalCase3 := validator.IsValidNumericFromString(badDecimalCase3)
		isDecimalCase4 := validator.IsValidNumericFromString(badDecimalCase4)

		assert.Equal(t, false, isDecimalCase1)
		assert.Equal(t, false, isDecimalCase2)
		assert.Equal(t, false, isDecimalCase3)
		assert.Equal(t, false, isDecimalCase4)
	})
}

func TestValidatePasswordComplexity(t *testing.T) {
	var (
		goodPasswords = []string{
			"P@ssw0rd#1234",
			"Aa1.",
			"Aa1.Bb2_",
			"Aa1.อิอิครุคริ", // non ascii char
		}

		badPasswords = []string{
			"aa1.เอนเวนเจอจงเจริญ", // no uppercase
			"AA1.เอนเวนเจอจงเจริญ", // no lowercase
			"Aa$.เอนเวนเจอจงเจริญ", // no number
			"Aa12เอนเวนเจอจงเจริญ", // no symbol
		}
	)

	t.Run("Good passwords", func(t *testing.T) {
		for _, pw := range goodPasswords {
			assert.True(t, validator.ValidatePasswordComplexity(pw))
		}
	})

	t.Run("Bad passwords", func(t *testing.T) {
		for _, pw := range badPasswords {
			assert.False(t, validator.ValidatePasswordComplexity(pw))
		}
	})
}

func TestIsValidBoolFromString(t *testing.T) {
	badBoolCase1 := ""
	badBoolCase2 := "-"
	badBoolCase3 := "TRue"
	badBoolCase4 := "FalSe"
	badBoolCase5 := "a"
	badBoolCase6 := "o"
	badBoolCase7 := "-1"
	badBoolCase8 := "2"

	mockedBoolCase1 := "1"
	mockedBoolCase2 := "t"
	mockedBoolCase3 := "T"
	mockedBoolCase4 := "True"
	mockedBoolCase5 := "TRUE"
	mockedBoolCase6 := "0"
	mockedBoolCase7 := "f"
	mockedBoolCase8 := "F"
	mockedBoolCase9 := "False"
	mockedBoolCase10 := "FALSE"

	t.Run("Happy", func(t *testing.T) {
		isBoolCase1 := validator.IsValidBoolFromString(mockedBoolCase1)
		isBoolCase2 := validator.IsValidBoolFromString(mockedBoolCase2)
		isBoolCase3 := validator.IsValidBoolFromString(mockedBoolCase3)
		isBoolCase4 := validator.IsValidBoolFromString(mockedBoolCase4)
		isBoolCase5 := validator.IsValidBoolFromString(mockedBoolCase5)
		isBoolCase6 := validator.IsValidBoolFromString(mockedBoolCase6)
		isBoolCase7 := validator.IsValidBoolFromString(mockedBoolCase7)
		isBoolCase8 := validator.IsValidBoolFromString(mockedBoolCase8)
		isBoolCase9 := validator.IsValidBoolFromString(mockedBoolCase9)
		isBoolCase10 := validator.IsValidBoolFromString(mockedBoolCase10)

		assert.Equal(t, true, isBoolCase1)
		assert.Equal(t, true, isBoolCase2)
		assert.Equal(t, true, isBoolCase3)
		assert.Equal(t, true, isBoolCase4)
		assert.Equal(t, true, isBoolCase5)
		assert.Equal(t, true, isBoolCase6)
		assert.Equal(t, true, isBoolCase7)
		assert.Equal(t, true, isBoolCase8)
		assert.Equal(t, true, isBoolCase9)
		assert.Equal(t, true, isBoolCase10)
	})

	t.Run("Fail", func(t *testing.T) {
		isBoolCase1 := validator.IsValidBoolFromString(badBoolCase1)
		isBoolCase2 := validator.IsValidBoolFromString(badBoolCase2)
		isBoolCase3 := validator.IsValidBoolFromString(badBoolCase3)
		isBoolCase4 := validator.IsValidBoolFromString(badBoolCase4)
		isBoolCase5 := validator.IsValidBoolFromString(badBoolCase5)
		isBoolCase6 := validator.IsValidBoolFromString(badBoolCase6)
		isBoolCase7 := validator.IsValidBoolFromString(badBoolCase7)
		isBoolCase8 := validator.IsValidBoolFromString(badBoolCase8)

		assert.Equal(t, false, isBoolCase1)
		assert.Equal(t, false, isBoolCase2)
		assert.Equal(t, false, isBoolCase3)
		assert.Equal(t, false, isBoolCase4)
		assert.Equal(t, false, isBoolCase5)
		assert.Equal(t, false, isBoolCase6)
		assert.Equal(t, false, isBoolCase7)
		assert.Equal(t, false, isBoolCase8)
	})
}

func TestIsValidDateTimeFromString(t *testing.T) {
	dateTimeLayout1 := time.RFC3339
	goodDateTimeLayout1Case1 := "2012-12-09T11:45:26.000Z"
	goodDateTimeLayout1Case2 := "2012-12-09T11:45:26.0Z"
	goodDateTimeLayout1Case3 := "2012-12-09T11:45:26Z"

	badDateTimeLayout1Case1 := ""
	badDateTimeLayout1Case2 := "2234"
	badDateTimeLayout1Case3 := "2012-99-23T11:45:26.000Z"
	badDateTimeLayout1Case4 := "2012-02-30T11:45:26.000Z"
	badDateTimeLayout1Case5 := "2012-02-30T11:45.000Z"

	dateTimeLayout2 := "15:04:05"
	goodDateTimeLayout2Case1 := "12:13:14"
	goodDateTimeLayout2Case2 := "13:14:15"

	badDateTimeLayout2Case1 := ""
	badDateTimeLayout2Case2 := "142323523423"
	badDateTimeLayout2Case3 := "13:14:99"
	badDateTimeLayout2Case4 := "13:14"
	badDateTimeLayout2Case5 := "13:14 AM"
	badDateTimeLayout2Case6 := "2012-02-01T11:45:26.000Z"

	t.Run("Happy", func(t *testing.T) {
		isDateTimeLayout1Case1 := validator.IsValidDateTimeFromString(dateTimeLayout1, goodDateTimeLayout1Case1)
		isDateTimeLayout1Case2 := validator.IsValidDateTimeFromString(dateTimeLayout1, goodDateTimeLayout1Case2)
		isDateTimeLayout1Case3 := validator.IsValidDateTimeFromString(dateTimeLayout1, goodDateTimeLayout1Case3)

		isDateTimeLayout2Case1 := validator.IsValidDateTimeFromString(dateTimeLayout2, goodDateTimeLayout2Case1)
		isDateTimeLayout2Case2 := validator.IsValidDateTimeFromString(dateTimeLayout2, goodDateTimeLayout2Case2)

		assert.Equal(t, true, isDateTimeLayout1Case1)
		assert.Equal(t, true, isDateTimeLayout1Case2)
		assert.Equal(t, true, isDateTimeLayout1Case3)

		assert.Equal(t, true, isDateTimeLayout2Case1)
		assert.Equal(t, true, isDateTimeLayout2Case2)
	})

	t.Run("Fail", func(t *testing.T) {
		isDateTimeLayout1Case1 := validator.IsValidDateTimeFromString(dateTimeLayout1, badDateTimeLayout1Case1)
		isDateTimeLayout1Case2 := validator.IsValidDateTimeFromString(dateTimeLayout1, badDateTimeLayout1Case2)
		isDateTimeLayout1Case3 := validator.IsValidDateTimeFromString(dateTimeLayout1, badDateTimeLayout1Case3)
		isDateTimeLayout1Case4 := validator.IsValidDateTimeFromString(dateTimeLayout1, badDateTimeLayout1Case4)
		isDateTimeLayout1Case5 := validator.IsValidDateTimeFromString(dateTimeLayout1, badDateTimeLayout1Case5)

		isDateTimeLayout2Case1 := validator.IsValidDateTimeFromString(dateTimeLayout2, badDateTimeLayout2Case1)
		isDateTimeLayout2Case2 := validator.IsValidDateTimeFromString(dateTimeLayout2, badDateTimeLayout2Case2)
		isDateTimeLayout2Case3 := validator.IsValidDateTimeFromString(dateTimeLayout2, badDateTimeLayout2Case3)
		isDateTimeLayout2Case4 := validator.IsValidDateTimeFromString(dateTimeLayout2, badDateTimeLayout2Case4)
		isDateTimeLayout2Case5 := validator.IsValidDateTimeFromString(dateTimeLayout2, badDateTimeLayout2Case5)
		isDateTimeLayout2Case6 := validator.IsValidDateTimeFromString(dateTimeLayout2, badDateTimeLayout2Case6)

		assert.Equal(t, false, isDateTimeLayout1Case1)
		assert.Equal(t, false, isDateTimeLayout1Case2)
		assert.Equal(t, false, isDateTimeLayout1Case3)
		assert.Equal(t, false, isDateTimeLayout1Case4)
		assert.Equal(t, false, isDateTimeLayout1Case5)

		assert.Equal(t, false, isDateTimeLayout2Case1)
		assert.Equal(t, false, isDateTimeLayout2Case2)
		assert.Equal(t, false, isDateTimeLayout2Case3)
		assert.Equal(t, false, isDateTimeLayout2Case4)
		assert.Equal(t, false, isDateTimeLayout2Case5)
		assert.Equal(t, false, isDateTimeLayout2Case6)
	})
}

func TestIsIsWeakPin(t *testing.T) {
	t.Run("Happy Weak", func(t *testing.T) {
		arrWeakPin := utils.GetWeakPin6Digit()
		for _, pin := range arrWeakPin {
			isWeekPinCase := validator.IsWeakPin6Digit(pin)
			assert.Equal(t, true, isWeekPinCase)
		}
	})

	t.Run("Happy Strong", func(t *testing.T) {
		Strong1 := validator.IsWeakPin6Digit("234153")
		Strong2 := validator.IsWeakPin6Digit("278146")
		Strong3 := validator.IsWeakPin6Digit("947351")
		Strong4 := validator.IsWeakPin6Digit("472389")

		assert.Equal(t, false, Strong1)
		assert.Equal(t, false, Strong2)
		assert.Equal(t, false, Strong3)
		assert.Equal(t, false, Strong4)
	})
}

func TestIsValidUsername(t *testing.T) {
	t.Run("Happy True", func(t *testing.T) {
		assert.True(t, validator.IsValidUsername("thailand"))
		assert.True(t, validator.IsValidUsername("username01"))
		assert.True(t, validator.IsValidUsername("inwza007"))
		assert.True(t, validator.IsValidUsername("0891234567"))
		assert.True(t, validator.IsValidUsername("007inwza"))
	})

	t.Run("Happy False less than 6 digit", func(t *testing.T) {
		assert.False(t, validator.IsValidUsername("thai"))
		assert.False(t, validator.IsValidUsername("aa"))
		assert.False(t, validator.IsValidUsername("bb"))
	})

	t.Run("Happy False miss regexp", func(t *testing.T) {
		assert.False(t, validator.IsValidUsername("Thailand"))
		assert.False(t, validator.IsValidUsername("InwZa007"))
		assert.False(t, validator.IsValidUsername("thailand-test"))
		assert.False(t, validator.IsValidUsername("thailand-ทดสอบ"))
		assert.False(t, validator.IsValidUsername("thailand-'''"))
		assert.False(t, validator.IsValidUsername("*778#$$"))
	})

	t.Run("Happy False duplicate characters", func(t *testing.T) {
		assert.False(t, validator.IsValidUsername("aaaaan"))
		assert.False(t, validator.IsValidUsername("inwza007aaaaa"))
		assert.False(t, validator.IsValidUsername("aaaaainwza007"))
		assert.False(t, validator.IsValidUsername("00000inwza007"))
		assert.False(t, validator.IsValidUsername("inwza007iiiii"))
		assert.False(t, validator.IsValidUsername("0899999999"))
	})

	t.Run("Json is valid", func(t *testing.T) {
		assert.True(t, validator.IsJSON(`{ "id": "62531347b3ed0e9865d72833" }`))
		assert.True(t, validator.IsJSON(`[ { "id": "62531347b3ed0e9865d72833" }, { "id": "62531347b3ed0e9865d72833" } ]`))
	})
	t.Run("Json is not valid", func(t *testing.T) {
		assert.False(t, validator.IsJSON(`xxx`))
		assert.False(t, validator.IsJSON(`11118374`))
	})

	t.Run("Base64 is valid", func(t *testing.T) {
		assert.True(t, validator.IsBase64("ZjQxMzkzYzEtNWFlNC00ZjI0LWFjZTktZDU0YjMxYjJjNmNm"))
		assert.True(t, validator.IsBase64(`PD94bWwgdmVyc2lvbj0iMS4wIiA/Pjxzdmcgd2lkdGg9IjQ4cHgiIGhlaWdodD0iNDhweCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBkYXRhLW5hbWU9IkxheWVyIDEiIGlkPSJMYXllcl8xIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjx0aXRsZS8+PHBhdGggZD0iTTgsMUEyLDIsMCwwLDAsNiwzVjQ1YTIsMiwwLDAsMCw0LDBWM0EyLDIsMCwwLDAsOCwxWiIvPjxwYXRoIGQ9Ik00My41NSwxMy43NEMzOC4yMiw3LjE4LDMyLjcxLDcuNjIsMjcuODQsOGMtNC42My4zNy04LjI5LjY2LTEyLjI5LTQuMjdBMiwyLDAsMCwwLDEyLDVWMjJhMiwyLDAsMCwwLC45NCwxLjcsOS4wOSw5LjA5LDAsMCwwLDQuOTEsMS40NmM0LDAsNy44LTIuNjIsMTEuMjgtNSw1LjE0LTMuNTMsOC40OS01LjUyLDExLjgxLTMuNDVhMiwyLDAsMCwwLDIuNjEtM1pNMjYuODcsMTYuODVDMjIuMjIsMjAsMTksMjIsMTYsMjAuNzhWOS42NmM0LjE4LDMsOC4zNywyLjYzLDEyLjE2LDIuMzMsMi41NC0uMiw0Ljc5LS4zOCw3LC4zMUMzMi4yMywxMy4xNywyOS40NiwxNS4wNywyNi44NywxNi44NVoiLz48L3N2Zz4=`))
	})

	t.Run("Base64 is not valid", func(t *testing.T) {
		assert.False(t, validator.IsBase64("xxx"))
		assert.False(t, validator.IsBase64(`1234567`))
	})
}
