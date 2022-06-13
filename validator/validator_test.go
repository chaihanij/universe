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

	t.Run("Base64DataType is valid", func(t *testing.T) {
		assert.True(t, validator.IsBase64DataType("data:image/png;base64,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"))
	})

	t.Run("Base64DataType is not valid", func(t *testing.T) {
		assert.False(t, validator.IsBase64DataType("xxx:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAgAAAAIACAYAAAD0eNT6AAAACXBIWXMAAAsTAAALEwEAmpwYAAAgAElEQVR4nO2dC5hlV13l0x3yJBAxDxIS8ui+59yqSnf13edU3b1vJeFCCKQzIJLBVnQEZCBqRAwOahRxgEg+YQQd5CUCHxhFjRGHlwpEZCLKQ0IEeQmEdwCJmpAEDCFJO/tUuoamu7rr3PPYa+1z1vm+9XU6kN+u9f/vx7p17z3nkPPPf+jhXod5HVLlKv67Pf/94XtJPPHEE0888cRj5nXKjHjiiSeeeOKJVx4CG1w88cQTTzzxxIPyoIOLJ5544oknnnjhedDBxRNPPPHEE0+88Dzo4OKJJ5544oknHoYHHVw88cQTTzzxxMPwoIOLJ5544oknnngYHnRw8cQTTzzxxBMPw4MOLp544oknnnjiYXidMiOeeOKJJ5544pXjdcqMeOKJJ5544olXkdcpM+KJJ5544oknng5/8cQTTzzxxBOPaXDxxBNPPPHEE0+Hv3jiiSeeeOKJ1zKvU2bEE0888cQTT7xyvE6ZEU888cQTTzzxyvE6ZUY88cQTTzzxxCvH65QZ8cQTTzzxxBOvHK9TZsQTTzzxxBNPvNKA7pgRTzzxxBNPPPHKQ2CDiyeeeOKJJ554UB50cPHEE0888cQTLzwPOrh44oknnnjiiReeBx1cPPHEE0888cTD8KCDiyeeeOKJJ554GB50cPHEE0888cQTD8ODDi6eeOKJJ5544mF40MHFE0888cQTTzwMr1NmxBNPPPHEE0+8crxOmRFPPPHEE0888SryOmVGPPHEE0888cTT4S+eeOKJJ5544jENLp544oknnnji6fAXTzzxxBNPPPFa5nXKjHjiiSeeeOKJV47XKTPiiSeeeOKJJ145XqfMiCeeeOKJJ5545XidMiOeeOKJJ5544pXjdcqMeOKJJ5544olXGtAdM+KJJ5544oknXnkIbHDxxBNPPPHEEw/Kgw4unnjiiSeeeOKF50EHF0888cQTTzzxwvOgg4snnnjiiSeeeBgedHDxxBNPPPHEEw/Dgw4uHidvfjzeMjRuZ5rZn0mNe4nX6xJj/9T//Zo0G783ydx1JfShdVTmvxNPPPHE+14Z+3epse8o9qEks6/1+9Nvpbm9xP/58LmJO5V5P2XmQQcXj4M3P15Oi8XkF9if+QX272nm/lOSJCke2a8X4WCY2UvnMpsj99OYeNDBxcPx5s3kdH/gX+Z1A37xSpIkNSjjPuEDwXPTxfGZMe7PoXidMiPehtfm4Wj8Q8Wv8eELVJIkqX3dU7x1MMjtIyLYn4PzOmVGvANem4f5ZFeRigkWpCRJEkD2H4t98IwzTttEtj/z8DplRrxDEjN5SGLsR/GLT5IkCa/iQ4UL+XjCsD9T8Tplpue8heXlk/xEv9JP+N3oBSdJksQk/6LorqFxL92+4o7rwn5fm9cpMz3nJSN3fmrcv6AXmSRJErWM/eJWszIJuT/T8Tplpse86XR6n9VPvhYffEEvLEmSpAhU/Dag2Df9Froppv2+EV6nzPSYt337OQ9Is/G16MUkSZIUoxIzvmow2HlEDPt9Y7xOmekpr3i/30/gD6MXkCRJUsxKjHv3ljw/lnm/b5TXKTM95M2PbJJk7vPohSNJktQJGXv9lsXJiYz7feO8TpnpGS/Z4U5JMvsF+IKRJEnqkBJj/+n0HdPvY9rvW+F1ykyPeHPLy8f5ifpx9EKRJEnqolbfDphOj2TY71vjdcpMT3in+0mp2/lKkiS1LGOvXlmxR3Tp/NgPAhtcvEo8n0xfAV8YkiRJfVBuf6lL50fti91Ml3mr9/RHLwhJkqSeqLhPwNySm3bh/Kh9sZvpMq94tGVq3DfQC0KSJKlXMvaLax8KDLXfh+BBBxdvNp4//N8GXwiSJEk9VGLsy0Lu9yF40MHFK88bGPs49AKQJEnqse5JcjeO8fw4EA86uHjleIuLi/f1r/6/RLAAJEmS+quR/YDfkjfHdH4cjAcdXLxyV2rsZfCJL0mSJP1n8UHsmM6Pg/Ggg4u38VV85z8x7qvoSS9JkiSt6sOHrPPkQMbzYyMedHDxNr4SY59OMOElSZKkPRpkk0fHcH5sxIMOLt6G12bd61+SJIlMxv1tBOfH7LxOmYmcl4zc+fCJLkmSJO2nucXllPn8mJnXKTMd4CWZuxI9ySVJkqT9lRj7XObzQ4d/xLyzzpoek2T2dvQkX5WxX/Rh5E/8hP9t/+evrH4rwWto7LP21dr/VkXiiSeeeOvx/N//5zCzL/P70Z+lmb0Jvieu7ovus514UBD7YdhH3tzSBHrjH3/Q31nc+Wpr7rIY6yeeeOJ1lLdr16Fpbs9JzPgP/V51D3KfnMvsUnT1oxlcvAPy9qRd0OFv/2qQr2yNuX7iiSde93mDpYnx+9U/oPbKYeZ+Ieb6dWoydInnJ9fHIZPauJcUCTv2+oknnnj94A0GO49IM/v7mBBg3xp1/dDNE29/3ny+coqfXLsBE/rXu1A/8cQTr3e8TZAPTRt364FeMLXstxkeSfPE24uXLtkLAYf/Gw9Z5+5WMdZPPPHE6x+vuGuqP5DfH3rvnB/ZBOG3ER5L88T7Ls9P4meEnMCJsbedae0Du1I/8cQTr588v5fNe90Vcv/c966AIf3W5jE1T7x7ecWn7wMHgGd3qX7iiSdef3lpZl8XdP/M7TORfmvx2Jon3kOLp/+9I9gENvbuNM+P71L9xBNPvP7yhsuTccgA4PUqpN8meNDBxftenj+Urw83ee3/RfsVTzzxxGuSl2TucwEDwBvRfuvwoIOLtz/PT95PBZy8z0H7FU888cRrkpca97pge6hx70T7rcqDDi7e+rzEuK8Fm7wj9xS0X/HEE0+8Jnl+D708XACw70P7rcqDDi7e+rziU/mhJm+Sjx+F9iueeOKJ1yQvze0lwQJA5j6G9luVBx1cvPV5PlH+R7AAMHLno/2KJ5544jXJG5rJfw+2hxr7z2i/VXnQwcVb/wIEgE7VTzzxxOs3DxUAYqsfdHDx1r9CBoChcTvRfsUTTzzxmuQhAkCM9YMOLt76FzAAdKJ+4oknXr95oQMA2m9VHnRw8da/QAGgM/UTTzzx+s0DBIBu1I+heX3nAQJAp+onnnji9ZsHCADx14+leX3nAQIA1K944oknXpM8YACIs35Mzes7L/S3ANB+xRNPPPGa5IECQJz1Y2te33nsAYC9fuKJJ16/eaDPAMD8VuYxNq/vPOYAEEP9xBNPvH7zEN8CQPqtzGNsXt95rAEglvqJJ554/eYh7wSI8FuZx9i8vvMYA0BM9RNPPPH6zWMPADT1Y2xeB3ib6/DYAkAH+iGeeOL1iDeXTS5mDQBU9WNsHjPvjDNO27Sw5Ey65J6YGneFP6yvTjP74SRzn08ye7OfDHf5f/fWOj8fUwBg74d44okn3r481gBAVz/G5rHxBtvHpw6M+2nf6Kv8YX/ThpOiZAA40M/HEgBY+yGeeOKJdzAeYwCgrB9j8xh4p0+nRw7zya7iMF99VT/LpCgRAA728zEEALZ+iCeeeOKV5bEFAPb6QQdn4p1p7QOTzL4wNe7WypNigwCw0c+HDgBM/RBPPPHEm5XHFADY6wcdnIU3NCsPSjP7O40cvgcJAGV+PmQAYOmHeOKJJ15VHsu3ANjrBx2cgTedTu8zzOyltV7xlwwAZX8+VABg6Id44oknXl0eQwCIoX7QwdG8JLfWN/DDjU+KdQLALD8fIgAw9EM88cQTrwkeOgDEUj/o4EDepuJVf5K5O1uZFPsEgFl/vtABgKAf4oknnniN8ZABIKb6QQdH8OaWl49LjXtbq5NirwBQxS8gAHSmv+KJJ554qAAQW/2gg4fmpYvjM5PMfrr1SbEnAFT1GzIADI3b2ZX+iieeeOIVFyIAxFg/6OAheXNLk23+YL0xyKTwAaCOX2AAiLa/4oknnnhrF+JpgDHWDzp4KF6au7NT474RakLsCQCV/YICQLT9FU888cTb+wIEgG7Uj6F5TfKKV/7FPfqDHf73BoC31fELCADR9lc88cQTb98LEADirx9L85rirb7nb9xXgx7+q4fq9wSAmf0CAkCQfognnnjiheABA0Cc9WNqXhO8LXl+bJK5T4U+/PcJAJX8om8F3EY/xBNPPPFC8UABIM76sTWvCZ4/RN+AOPz3CgCV/bIHAIb+iieeeOId6AJ9BgDmtzKPsXl1eWlmfwZ1+K/q3gBQ2S9zAGDor3jiiSfewS7EtwCQfivzGJtXhzcYjRf8AXoHOABs+Djgg/llDQAM/RVPPPHE2+hC3wo4tN/KPMbm1eBtSoz7G+jhXzEA7O2XMQCQ9Fc88cQTb8OLPQDQ1I+xeVV5qRk/AX74VwgA+/plCwAs/RVPPPHEK8ObyyYXswYAqvoxNq8KbzAe3z817l/gh/+MAWA9v0wBgKW/4oknnnhleawBgK5+jM2rwksy9yvwg3/GAHAgvywBgKm/4oknnnhleYwBgLJ+jM2blXdynh+dZvbr8IN/hgBwML8MAYCpv+KJJ554s/DYAgB7/aCD1+X5JjwDfujPEAA28osOAGz9FU888cSbhccUANjrBx28Ad6mQeY+Bz/0SwaAMn6RAYCwv+KJJ554M/FYvgXAXj/o4E3wBiM3hR/4JQNAWb+oAMDYX/HEE0+8WXkMASCG+kEHb4KXZPY18AO/RACYxS8iALD2VzzxxBNvVh46AMRSP+jgdXmnOndUatyt+APf3ZqY8VXFLYiHxj7G//N8Hb+hAwBrf8UTTzzxqvCQASCm+kEHr8srDi/wwf+lJLM/MRjsPKJJv4AAQNlf8cQTT7wqPFQAiK1+0MHr8vxBeQXq8E8y9+LTp9Mj2/AbMgAMjdvJ2l/xxBNPvCo8RACIsX7Qwevy/Cvw9wc/+I37zsDYJ7fpFxgAqPornnjiiVeFh3gaYIz1gw5ehzc3scf6wt8VOgCkub2kbb+gAEDVX/HEE0+8qjxAAOhG/RiaV4Y3XJ6Mgx/+mXtVCL+AAEDXX/HEE0+8qjxAAIi/fizNK8PzDQ785D97y9zy8nEh/AICQO1+sM8X8cQTrz88YACIs35MzSvDG2bu8pABoHjYUCi/6FsBV+kH+3wRTzzx+sMDBYA468fWvDI8X/g3BAwAu5Md7pRQftkDQIzzRTzxxOsPD/QZAJjfyjzG5pXhFXfbCxYARvYDIf0yB4BY54t44onXHx7iWwBIv5V5jM0rw0uMe3ewAJDZl4b0yxoA4JNVPPHEE6/Ehb4VcGi/lXmMzSvDSzJ3XbAGZ+5ZIf0yBgCKySqeeOKJV+JiDwA09WNsXhleatwnwgUAe3FIv2wBgGayiieeeOKV4M1lk4tZAwBV/RibV4ZXFD1Ug4s0GdIvUwCgmqziiSeeeCV4rAGArn6MzSvDQwSAUH5ZAgDdZBVPPPHEK8FjDACU9WNsXhle6AAQ0i9DAED3F82zdvmY+eXJtnQ0/sHiMc++J5f5ej0/Ne5FSWZf6Offs4eZvbR4LsRg5KZbFicnxuxXPPG6xGMLAOz1gw5ehQcIAMH8ogMAQ39D87aPl09Ll9wTfe1fW8ytKs+Z8MHg5jQbX+v/2+cO8/HDihDB6lc88brMYwoA7PWDDl6VFzIAFJMppF9kAGDpbwjeYHT2CYlxT/cH9wdbqW3mvjk09o/nMvfoYy967KFov+KJ1xcey7cA2OsHHbwODxgAWveLCgBM/W2Tl44my977G4tHOwerc2a/4se7fKPnSbThVzzx+sZjCAAx1A86eB0eKAAE8YsIAGz9bYOXZONzfW3fEaq2BwgCt/uf4X+dae0DY6ufeOLFwkMHgFjqBx281mYePgAE8xs6ADD2t0neXJ6fnGTuSuTBv38QcN8sPiuwsLBwOHv9xBMvNh4yAMRUP+jgdXiAABDMLyAA0PW3Id5m7/EZfq7chj7wD6IPJ7m1pPUTT7woeagAEFv9oIPX4aFuBBTCb8gAMDRuJ2N/6/LSPD/ev8r+S4IDvswGclfx24ATTjhuM0v9xBMvZh4iAMRYP+jgdXjsAaDW4YULADT9rcNLzOQhxYfu0Af7zDLuL3Y4exK6fuKJFzsP8TTAGOsHHbwOjzkA1PULCgBU/a3K86/6L/L1uwN+mFcOAfaG+eWVua70QzzxEDxAAOhG/RiaV4bHGgCa8AsIAHT9rcJLzfhp3tM98EO89obivurnwGLo+oknXld4gAAQf/1YmleGxxgAmvILCAAVusE1X/wr/19EH9yNbioj++/D3G2PtR/iiYfkAQNAnPVjal4ZHlsAaNIv+lbAVfqBnC/+lf8TvJfd6EO78d5k9ivDHfaM2PohnnhoHigAxFk/tuaV4TEFgKb9sgcApvkyNPYxVe7bH4sSM/7kRncPZOqHeOIx8ECfAYD5rcxjbF4ZHksAaMMvcwBgmi9bjRukxt2KPqRb71Fm/8rb3czeD/HEY+EhvgWA9FuZx9i8MjyGANCWX9YAAJ+se12nT6dH+jpdjz6cg8nYy5j7IZ54TDz0rYBD+63MY2xeGR46ALTplzEAUEzWvS7/qviV8EM5oIq3ObaalQlrP8QTj4nHHgBo6sfYvDI8ZABo2y9bAKCZrHuuNLfnpB380F8JfSzP88PY+iGeeGy84vbtwfbQGQMAVf0Ym1eGhwoAIfwyBQCqyeqv6XR6nyRzHyE4jDEy9n8w9UM88Rh5rAGArn6MzSvDQwSAUH5ZAgDdZD2ke9/3r7DZ3LZt6ezTWfohnniMPMYAQFk/xuaV4YUOACH9MgQAdH/XuxYXF++bGPdv6EMYLmNfwtAP8cRj5bEFAPb6QQevwgMEgGB+0QGAob/rXalxvwQ/fBlk3Lfms5VT0P0QTzxWHlMAYK8fdPCqvJABoJhMIf0iAwBLf/e9iq/9rd4jH3348uj5yH6IJx4zj+VbAOz1gw5ehwcMAK37RQUApv7ueyW5exLBoUske1PxjQBUP8QTj5nHEABiqB908Do8UAAI4hcRANj6u+/lX/2/C3/ocmmQTR6N6od44jHz0AEglvpBB6/DAwSAYH5DBwDG/n5PPXa4U3xN7kYfuGxKzPgqRD/EE4+dhwwAMdUPOngdHiAABPMLCAB0/d2b53/OX0YftpQy9o6T8/zo0P0QTzx2HioAxFY/6OB1eOhbAbfpN2QAGBq3k7G/e/P8z3gN/LAlVWz3cRBPvBA8RACIsX7Qwevw2ANAHb/AAEDT3zWetcvHJJn7JvqgpZWxV4Tsh3jixcBDPA0wxvpBB6/DYw4Adf2CAgBVf9d4w2xyHvyQZZax7wvZD/HEi4EHCADdqB9D88rwWANAE34BAYCuv2usNBtfBj9kmWXsHYfs2nVoqH6IJ14MPEAAiL9+LM0rw2MMAE35BQSACt0IM198LV4LP2TJNdxhzwjVD/HEi4EHDABx1o+peWV4bAGgSb/oWwFX6Udb8yU17m/RByy7htnkgtjWr3jitckDBYA468fWvDI8pgDQtF/2ABByviTGfQ19wLLL1+jpsa1f8cRrkwf6DADMb2UeY/PK8FgCQBt+mQNA6PniD7fvoA9Ydg0z99zY1q944rXJQ3wLAOm3Mo+xeWV4DAGgLb+sASD0ZC0eAIQ+XGPQ0LgXx7Z+xROvTR76VsCh/VbmMTavDA8dANr0yxgAEJM1zfPj0YdrDBpm9vdiW7/iidcmjz0A0NSPsXlleMgA0LZftgCAmqxbl5YejD5cY5APAH8Q2/oVT7w2ecXt24PtoTMGAKr6MTavDA8VAEL4ZQoAyMk6t7x8HPpwjULGvTK29SueeG3yWAMAXf0Ym1eGhwgAofyyBAD0ZF1YWDgcfrhGoCSzLwzRD/HEi4XHGAAo68fYvDK80AEgpF+GAIDu79qVZvbb6AOWXX4tPDtUP8QTLwYeWwBgrx908Co8QAAI5hcdABj6u3b5n/HL6AOWXUlmfzJUP8QTLwYeUwBgrx908Kq8kAGgmEwh/SIDAEt/167EuHehD1h2DXNbujls/RVPvDZ4LN8CYK8fdPA6PGAAaN0vKgAw9Xft8gHg5egDll3JDndKqH6IJ14MPIYAEEP9oIPX4YECQBC/iADA1t+1y/f56egDlllJZm/3ZdoUqh/iiRcDDx0AYqkfdPA6PEAACOY3dABg7O/aleTWog9ZbtlrQvZDPPFi4CEDQEz1gw5ehwcIAMH8AgIAXX/XrmMveuyh/pC7BX/QcmqjbwCwrl/xxGuThwoAsdUPOngdHvpWwG36DRkAhsbtZOzv3jz/c74FfdDSKndnh+5HF3kn5/nRw9FkJTHjx6cj95SBcT+QLo7PZPn5xJuNhwgAMdYPOngdHnsAqOMXGABo+rs3b5iPfw5+0FLK3lLcLCl0P7rES7Lxub6Wb/Rr7o51a2zsDf7PX8vzc0/sgt++8BBPA4yxftDB6/CYA0Bdv6AAQNXfvXk7nD0pydyd+AOXTMb9HqIfXeANzcqDfIAq/ZulxLh/S3P31Fj99o0HCADdqB9D88rwWANAE34BAYCuv/vyfAB4E/zAJdNg5KaofsTMG+bjhyUj++9Vap6Y8VWnOndUTH77yAMEgPjrx9K8MjzGANCUX0AAqNCNsPMlzSePRR+4TEoy+wVfss2ofsTKGxj7uPq/TRpfOxiP7x+D377ygAEgzvoxNa8Mjy0ANOkXfSvgKv1ofb7s2nWoP/Q+jT54aWTcz0P7ESFvkI8f2dxbSfaaPM8P+D1qBr995oECQJz1Y2teGR5TAGjaL3sAQM2X4j1Y+MFLoOL96IWF6THofsTES5aWtjT9dVLfh1ew+u07D/QZAJjfyjzG5pXhsQSANvwyBwDkfFl9PLCxX0QevhTK7a8x9CMWXvFK3b/yv66dXkwey+ZXPMy3AJB+K/MYm1eGxxAA2vLLGgDgk/WQ4s6A7kmQQ5dESWa/MlxZuR9LP2Lg+b3iee31xH59S54fy+RXPPytgEP7rcxjbF4ZHjoAtOmXMQBQTNZ7r02JcX8T+uBl0SB3P0bWD2pesmNl5OfLd1rti3EvYvEr3r0XewCgqR9j88rwkAGgbb9sAYBmsu65hksrQ//K69shD14G+YPsPd7+JrZ+sPKKt4ySzH2k/b7Yu85anozQfsX7Lq+4fXu4dTlbAKCqH2PzyvBQASCEX6YAQDVZ97p8/381VI0Y5P3ettW4AWs/GHm+br8ebo9w70T7Fe+7PNYAQFc/xuaV4SECQCi/LAGAbrJ+77XZz4G3h6oTXMb+OHk/qHjD3G1v/Vf/+yqfPLYr9YudxxgAKOvH2LwyvNABIKRfhgCA7m8Z3sLSOcXtXG8MuskDlGT2tTH0g4i3Oc3G7w3eJ78nLS8vlXvGerN+xduHxxYA2OsHHbwKDxAAgvlFBwCG/pblnbXkXPHr8dCbfbjNxb17OLzgiFj6wcBLcgt8eJT9mdB+xdufxxQA2OsHHbwqL2QAKCZTSL/IAMDS31l480vjRxzwaW4Ry8/xjy4unvOA2PqB5K0+5Me4b+D6Zm9a+1pgjPXrCo/lWwDs9YMOXocHDACt+0UFAKb+zsobjsY/VHwaG7fxN9yX4rbHO9wpsfYDxWN4aJT/GV4Qa/26wmMIADHUDzp4HR4oAATxiwgAbP2twiu8dOLtAGM/tHVxcmLs/QjNY3lgVPG8gUG+sjW2+nWJhw4AsdQPOngdHiAABPMbOgAw9rcqb5CN3eqz2wkOgmqbiXtXMh7fvyv9CMU7Oc+P9gfv59H9+67sH8VUv67xkAEgpvpBB6/DAwSAYH4BAYCuv3V4xasvfxh8EH8IzKTdqXG/OZ1O74OuX4y8onYEPfyefg5Hk5VY6tc1HioAxFY/6OB1eOhbAbfpN2QAGBq3k7G/dXnFQVq8F7t6sOIPg4PLuFuLzzAw1S8m3tzSZFvw7/yX0vi9/sfbxF6/LvIQASDG+kEHr8NjDwB1/AIDAE1/m+KlI3uhr+cN+MPgQJvH+Krik+us9YuAV9wQ6u/RfTyQBsY+jrx+neQhngYYY/2gg9fhMQeAun5BAYCqv03ytkynRw4z97yQdd1w08jsp4fZ5IIY6sfM86/8fwrdy4PKuM/qPg7heYAA0I36MTSvDI81ADThFxAA6PrbBm9+PN4yzOzL/Kb8LdSBkGTuU8nIPrF4iyK2+rHx0jw/3geAf0X1coY19guM9esyDxAA4q8fS/PK8BgDQFN+AQGgQjfimi9784qv2PkaX+FfhX8lUJ3vWf10vxk//pBduw6NvX4sPB+mrgy1TmodEJm9eYezJ7HVr8s8YACIs35MzSvDYwsATfpF3wq4Sj/Y58u6PH8YD/LxI+89SOwtDdf2Hs+9zs/TZ2/Jzz6Nwm+HeMN8/LA0hg94rsnYlzDVr+s8UACIs35szSvDYwoATftlDwAxzpcNeT4MzGU2T3L7TB8G3lJ8cHCWOwve+6vo8bX+z5cP88muueXl46j9Rswr3lP3Nf8E/FCf6ZBw35kf2YShfn3ggT4DAPNbmcfYvDI8lgDQhl/mABDrfKnCy/P8sLnF5XQum1zgQ8GPpEvuiWluL/H/fElxyA9y+4h0NFku3otG/Hx95fn1cTn6QK+oNzLUrw88xLcAkH4r8xibV4bHEADa8ssaAOCTVbze87Zlbj7mBz8l2fhcZP36wkPfCji038o8xuaV4aEDQJt+GQMAxWQVr9e8lZXx4Wlm34U+xGvJ2Ou9nc2I+vWJxx4AaOrH2LwyPGQAaNsvWwCgmazi9ZqX5vYp8AO8AQ1y92OI+vWJV9y+nTUAUNWPsXlleKgAEMIvUwCgmqzi9ZY3Onvlgf7V/03ow7shfbl4eFHM/WDnsQYAuvoxNq8MDxEAQvllCQB0k1W83vKGmXs9wcHdnHL7yzH3g53HGAAo68fYvDK80AEgpF+GAIDur3jirfGG2eS8NKbv/Jc7NG5bWF4+KcZ+xMBjCwDs9YMOXoUHCADB/KIDAEN/xROv4Fi7fIyfp59EH9jtHBzuFbH1IxYeUwBgrx908Kq8kGY4ckYAACAASURBVAGgmEwh/SIDAEt/xRPv/NVX/+556IO6NRl799Ydy2fF1I9YeCzfAmCvH3TwOjxgAGjdLyoAMPVXPPG2524h5u/8l5JxfxFLP2LiMQSAGOoHHbwODxQAgvhFBAC2/oonXprZv4Yf0CG0ZC+MoR8x8dABIJb6QQevwwMEgGB+QwcAxv6K12+ef2X8ZPjBHE4f2XbhBUcy9yM2HjIAxFQ/6OB1eIAAEMwvIADQ9Ve8/vJGK5PiO/9fJziYg2lg7JNZ+xEjDxUAYqsfdPA6PPStgNv0GzIADI3bydhf8frLSzL7GvSBHFre81cWFxfvy9iPGHmIABBj/aCD1+GxB4A6foEBgKa/4vWTVzwsJ+3Yd/5n0HPY+hErD/E0wBjrBx28Do85ANT1CwoAVP0Vr3+87du3He7n5McJDmKM/LrfurT0YJZ+xMwDBIBu1I+heWV4rAGgCb+AAEDXX/H6x/Nr+tnwQxisJLOvZelHzDxAAIi/fizNK8NjDABN+QUEgArdiGu+iMfN22rcIOS8J9Y9W3OXofsROw8YAOKsH1PzyvDYAkCTftG3Aq7SD/b5Ih43L83sNQSHL4nsNeh+xM4DBYA468fWvDI8pgDQtF/2ABDjfBGPl5eM7BPxhy6Xktz+l670F8EDfQYA5rcyj7F5ZXgsAaANv8wBINb5Ih4nb255+bi+fee/3KEy/uR0Or1P7P1F8RDfAkD6rcxjbF4ZHkMAaMsvawCAT1bxOsfzc/316MOWVQPjfjr2/qJ46FsBh/ZbmcfYvDI8dABo0y9jAKCYrOJ1ijfM7UPT/n7nv4TsTdvc+IRY+4vksQcAmvoxNq8MDxkA2vbLFgBoJqt4neENBjuPCLmGY9XQ2BfG2F80r7h9e7A9dMYAQFU/xuaV4aECQAi/TAGAarKK1xmen+NXoA/XKGTsHduXVgax9RfNYw0AdPVjbF4ZHiIAhPLLEgDoJqt4neDNLU22JcZ9B364RiK/1/1hTP1l4DEGAMr6MTavDC90AAjplyEAoPsrXmd5m/3a/Xv0oRqZdqejyXIk/aXgsQUA9vpBB6/CAwSAYH7RAYChv+J1k5ea8dMIDtQINX6vL98m9v6y8JgCAHv9oINX5YUMAMVkCukXGQBY+ite93hzeX5ymtlb8IdpnEoydxFzf5l4LN8CYK8fdPA6PGAAaN0vKgAw9Ve87vH8vP5z9CEatYz7bPHtCdb+MvEYAkAM9YMOXocHCgCBfk0aPgCw9Ve8bvGSfPwo+AHaAQ0zeyljf9l46AAQS/2gg9fhAQJAyFdKQQMAY3/F6w5vuLJyPz/Xvow+PLugJLM3F7dPZuovIw8ZAGKqH3TwOjxAAAjmFxAA6PorXnd4aWZfij44u6Qkcy9m6i8jDxUAYqsfdPA6PPStgNv0GzIADI3bydhf8brBK76+5ufz3ehDs0vyAeDO+ZFNGPrLykMEgBjrBx28Do89ANTxCwwANP0VL35e8TQ7P5evRx+YXZQPAVej+8vMQzwNMMb6QQevw2MOAHX9ggIAVX/Fi5+X5vaX0QdllzW3ZB/apfnSJA8QALpRP4bmleGxBoAm/AICAF1/xYubN9xhz/CvUr+JPiQ7LWP/YWXFHtGF+dI0DxAA4q8fS/PK8BgDQFN+AQGgQjfimi/iBeVtSo17J/yA7IXsf+vAfGmcBwwAcdaPqXlleGwBoEm/6FsBV+kH+3wRLxwvze0l+IOxH0oy9/nTp9MjQ/Y3Bh4oAMRZP7bmleExBYCm/bIHgBjni3hheGk+mfOv/r+FPhh7JeN+Kdb50hYP9BkAmN/KPMbmleGxBIA2/DIHgFjni3jt8wbj8f39uvwo/EDsmXzNbzvT2gfGNl/a5CG+BYD0W5nH2LwyPIYA0JZf1gAAn6zi0fKKr/z5Nfl29GHYV/navyym+dI2D30r4NB+K/MYm1eGhw4AbfplDAAUk7XjvC352aclZvxfff8vS7Px7/o5/qdpZq8p5P/5zYlxf+B79vw0d09NRuMd2y684EgGv3meH1Z8Lx19CPZZfm58p3j7pY3+xshjDwA09WNsXhkeMgC07ZctANBM1o7xtkynRw6NfYzv9+vTCvfKTzJ7+zCzbx1m4yefZcfHI/ye6txResofi+xbmu5vrLzi9u2h6j5rAKCqH2PzyvBQASCEX6YAQDVZO8Lblrl5P39/OzXuG4310rhv+VeBLy9uERvK78Ly8kl+3PfjDz5pTcN8/LCm+tv0fAnJYw0AdPVjbF4ZHiIAhPLLEgDoJmvkvO1LK4Nh5l7d8r3x70kyd+WWxcmJbfodZpMLvI8bQx5uUil92Ldnc4j5zMxjDACU9WNsXhle6AAQ0i9DAED3t0u8beeefezQ2Beuvk8b7CCwt/g/n1G8P9+k36FZeZBn/344H9KsSnL3pJjWRxs8tgDAXj/o4FV4gAAQzC86ADD0tyu8hXw88QfmZ2CHgRl/0vf4R/cOAlX8Fge/DzC/odv7RqEvz03ssTGsj7Z4TAGAvX7QwavyQgaAYjKF9IsMACz97QJvmNuf9Yf/twkOhOIzAl/yB/jlyY6VUVm/i4uL9x3mk13FJ/yLR9DCPUil5efec9nXR5s8lm8BsNcPOngdHjAAtO4XFQCY+hszr3hAy9C4F6MPgQNvWO5rq98eMPYFPhg8w8+3HykOev/nj/uD/ln+//Mq/88f8mvsLvTPKlXscWZv3+HsgxnXRwgeQwCIoX7QwevwQAEgiF9EAGDrb6y8PYf/a9AHgCT5feS1bOsjFA8dAGKpH3TwOjxAAAjmN3QAYOxvrDz/qvpF8I1fkgoZe/cwd9uZ1kcoHjIAxFQ/6OB1eIAAEMwvIADQ9TdG3nDJ/RR805ekvZQY9y6W9RGShwoAsdUPOngdHvpWwG36DRkAhsbtZOxvbLz5fPIovWcuUcqvcfT6CM1DBIAY6wcdvA6PPQDU8QsMADT9jYk3n7vFPd+7x2/2krSvjPtE8bAm1PpA8BBPA4yxftDB6/CYA0Bdv6AAQNXfWHjb7fhk368b4Ju8JB1ESWZ/sgvrrSwPEAC6UT+G5pXhsQaAJvwCAgBdf2Pg7XjY9P5+Y30venOXpI1lvz4Yj+8f83qbhQcIAPHXj6V5ZXiMAaApv4AAUKEbcc2XpnnF1/18/d6A39glqbSeH+t6m5UHDABx1o+peWV4bAGgSb/oWwFX6Qf7fGma53v06wQbuiSVl99XtuRnnxbjepuVBwoAcdaPrXlleEwBoGm/7AEgxvnSJG9+afKjvna74Ru6JM2o4kmRsa23KjzQZwBgfivzGJtXhscSANrwyxwAYp0vTfFWH+5j3LfQG7kkVdTu+WzsYllvVXmIbwEg/VbmMTavDI8hALTllzUAwCcrmHdW5hJ/+P8LwSYuSTVkr41hvdXhoW8FHNpvZR5j88rw0AGgTb+MAYBisgJ5o3PPOc7PuY/iN29Jqq85M34c83qry2MPADT1Y2xeGR4yALTtly0A0ExWEG/xwkce5XvyNvSmLUmNydgbtm/fdjjjemuCV9y+PVQtZw0AVPVjbF4ZHioAhPDLFACoJiuIN8zsy+EbtiQ1LL+HPp1xvTXBYw0AdPVjbF4ZHiIAhPLLEgDoJiuAlxr3DPRGLUltKMnszQvOfT/TemuKxxgAKOvH2LwyvNABIKRfhgCA7i8DTw/4kTov436TZb01yWMLAOz1gw5ehQcIAMH8ogMAQ3/RvGFm9YAfqfNKMnfnVuMG6PXWNI8pALDXDzp4VV7IAFBMppB+kQGApb9I3rbxsh7wI/VGiRlf1aX1W/yd5VsA7PWDDl6HBwwArftFBQCm/qJ4iw+b3i/VA36kvim353Rh/a7xGAJADPWDDl6HBwoAQfwiAgBbfxG8lZVxcYvpP4RvxpIUWsa93y+LTTGv37156AAQS/2gg9fhAQJAML+hAwBjfxE8P6eeB9+IJQmnH455/e7NQwaAmOoHHbwODxAAgvkFBAC6/obmDfPJrlQP+JF6rEHmPjccXnBEjOt3Xx4qAMS2/0EHr8ND3wq4Tb8hA8DQuJ2M/Q3JS0eTZT3gR5IKjS+Lbf2ux0MEgBj3P+jgdXjsAaCOX2AAoOlvKN5whz1DD/iRpDXZW7bb8cmxrN8D8RBPA4xx/4MOXofHHADq+gUFAKr+huDNrazcz8+jf8JvupLEo2FmXxbD+j0YDxAAotv/1uUxNK8MjzUANOEXEADo+ts279iLHnuor/Nb0ZutJLGpuPvlYDReYF6/G/EAASCq/S/qw7/gMAaApvwCAkCFbsQ1X/bl+fnzMvRGK0msSjL3Jub1u9EFDABR7H9cg1fgsQWAJv2ibwVcpR/s8+V7Dv/c/hx6g5UkepnJwxnXb5kLFACi2P+4Bq/IYwoATftlDwAxzpc13jCbXKAH/EhSGdl/9MtpM9P6LXuBPgMA81uZx9i8MjyWANCGX+YAEOt8Kf598b5matw38BurJEUiM34Cy/qd5UJ8CwDptzKPsXlleAwBoC2/rAEAPllr8NI8P14P+JGkGWXsjSfn+dHo9Tvrhb4VcGi/lXmMzSvDQweANv0yBgCKyVqRd/p0emSajfWAH0mqIL/X/mps+wF7AKCpH2PzyvCQAaBtv2wBgGayVuNtSsxYD/iRpIpKMnv7XJ6fHNN+UNy+PVh9ZgwAVPVjbF4ZHioAhPDLFACoJmsFXmLc5egNVJI6oFfFtB+wBgC6+jE2rwwPEQBC+WUJAHSTdUae9/fDqR7wI0n1Zezdw9xtj2U/YAwAlPspY/PK8EIHgJB+GQIAur91ecPRZMXX8Q74xilJHVGS2bfHsh+wBQD2/RQ6eBUeIACEe+UKDgAM/a3DW33AT2a/jt4wJalrms8nj4phP2AKAOz7KXTwqryQAaCYTCH9IgMAS3+r8rbk+bHe18fRG6UkdVQfybIdR7DvByzfAmDfT6GD1+EBA0DrflEBgKm/lXi7dukBP5LUsvze+3j2/YAhAMSwn0IHr8MDBYAgfhEBgK2/VXh6wI8kta8ks5+eTqf3Yd4P0AEglv0UOngdHiAABPMbOgAw9ndWnh7wI0nhlOTuScz7ATIAxLSfQgevwwMEgGB+AQGArr+z8FLjduoBP5IUUCP7Adb9oLhQAYD1vDwQDzp4HR76VsBt+g0ZAIb+8GTsb1leujg+Uw/4kaTwGixNDNt+sHYhAgDzeXkgHnTwOjz2AFDHLzAA0PS3DC/P88P84f9+9EYoSX1UktlXMu0He1+IpwEyn5cH4kEHr8NjDgB1/YICAFV/y/D8BvRC9CYoSX2VX39fYNoP9r4AAYD6vCzNY2heGR5rAGjCLyAA0PV3w8N/x8qouD0pehOUpD4r2eFOYdgP9r0AAYD6vOzU4V9wGANAU34BAaBCN7DzJTHuXejNT5L6riRzFzHsB/tewABAeV526vBnDABN+kXfCrhKP0LOl3Q0WUZvfJIkrb4NcDF6P1jvAgUA2vOyU4c/WwBo2i97AEDPl9S416E3PkmS/N6Y2UvR+8F6F+gzADC/lXmMzSvDYwkAbfhlDgDo+VLcgSwx7l/RG58kScXh536W8fxAfAsA6bcyj7F5ZXgMAaAtv6wBAD5Z/ZXmkzn0pidJ0r2ay8ePZDw/0LcCDu23Mo+xeWV46ADQpl/GAEAxWf01yO0j0JueJEmr2r3djk9mPD/YAwDLfnoIY/PK8JABoG2/bAGAZrIeUvR98hCCjU+SpMz9FXo/OBCvuH17qDrMGgCY9tNDGJtXhocKACH8MgUAqsnqr61LSw8m2PgkqfeaM/YH0PvBgXisAYBtPz2EsXlleIgAEMovSwCgm6x7riRzH0RvfpLUc71lZcUewbAfrMdjDACU+ylj88rwQgeAkH4ZAgC6vwe7Btnk0QQboCT1Uj6Af2qbG5/Ash+sx2MLAMz7KXzwKjxAAAjmFx0AGPq7ES817nfQG6Ek9U/2fXMTdyrbfrAvjykAsO+n0MGr8kIGgGIyhfSLDAAs/d2IV/z60dfpOf7VyJ34TVGSOi5j7xga91t5fv79GPeDfXks3wJg30+hg9fhAQNA635RAYCpv2V588uTbcWdAZPM3g7fJCWpY1q96ZZxr1ww9swY9oM1HkMAiGE/hQ5ehwcKAEH8IgIAW39n5S0uLt43HdkLvafn+w3rbWlmP5Nk7pvoDVSS4pH9tv/zY37/+XOvK4b5+GGLFz7yqBj3A3QAiGU/hQ5ehwcIAMH8hg4AjP1tilfcOnjBue9Plpa2DJYmZi6zeaGF3NqzlpxbU/H3tf+tisQTL0beVuMGxSN9t28/5wF+uWxiW79VecgAEFP9oIPX4QECQDC/gABA11/xxBNPvKo8VACIrX7Qwevw0LcCbtNvyAAwNG4nY3/FE0888aryEAEgxvpBB6/DYw8AdfwCAwBNf8UTTzzxqvIQTwOMsX7QwevwmANAXb+gAEDVX/HEE0+8qjxAAOhG/RiaV4bHGgCa8AsIAHT9FU888cSrygMEgPjrx9K8MjzGANCUX0AAqNCNuOaLeOKJ1x8eMADEWT+m5pXhsQWAJv2ibwVcpR/s80U88cTrDw8UAOKsH1vzyvCYAkDTftkDQIzzRTzxxOsPD/QZAJjfyjzG5pXhsQSANvwyB4BY54t44onXHx7iWwBIv5V5jM0rw2MIAG35ZQ0A8MkqnnjiiVfiQt8KOLTfyjzG5pXhoQNAm34ZAwDFZBVPPPHEK3GxBwCa+jE2rwwPGQDa9ssWAGgmq3jiiSdeCV5x+3bWAEBVP8bmleGhAkAIv0wBgGqyiieeeOKV4LEGALr6MTavDA8RAEL5ZQkAdJNVPPHEE68EjzEAUNaPsXlleKEDQEi/DAEA3V/xxBNPvKo8tgDAXj/o4FV4gAAQzC86ADD0VzzxxBOvKo8pALDXDzp4VV7IAFBMppB+kQGApb/iiSeeeFV5LN8CYK8fdPA6PGAAaN0vKgAw9Vc88cQTryqPIQDEUD/o4HV4oAAQxC8iALD1VzzxxBOvKg8dAGKpH3TwOjxAAAjmN3QAYOyveOKJJ15VHjIAxFQ/6OB1eIAAEMwvIADQ9Vc88cQTryoPFQBiqx908Do89K2A2/QbMgAMjdvJ2F/xxBNPvKo8RACIsX7Qwevw2ANAHb/AAEDTX/HEE0+8qjzE0wBjrB908Do85gBQ1y8oAFD1VzzxxBOvKg8QALpRP4bmleGxBoAm/AICAF1/xRNPPPGq8gABIP76sTSvDI8xADTlFxAAKnQjrvkinnji9YcHDABx1o+peWV4bAGgSb/oWwFX6Qf7fBFPPPH6wwMFgDjrx9a8MjymANC0X/YAEON8EU888frDA30GAOa3Mo+xeWV4LAGgDb/MASDW+SKeeOL1h4f4FgDSb2UeY/PK8BgCQFt+WQMAfLKKJ5544pW40LcCDu23Mo+xeWV46ADQpl/GAEAxWcUTTzzxSlzsAYCmfozNK8NDBoC2/bIFAJrJKp544olXglfcvp01AFDVj7F5ZXioABDCL1MAoJqs4oknnngleKwBgK5+jM0rw0MEgFB+WQIA3WQVTzzxxCvBYwwAlPVjbF4ZXugAENIvQwBA91c88cQTryqPLQCw1w86eBUeIAAE84sOAAz9FU888cSrymMKAOz1gw5elRcyABSTKaRfZABg6a944oknXlUey7cA2OsHHbwODxgAWveLCgBM/RVPPPHEq8pjCAAx1A86eB0eKAAE8YsIAGz9FU888cSrykMHgFjqBx28Dg8QAIL5DR0AGPsrnnjiiVeVhwwAMdUPOngdHiAABPMLCAB0/RVPPPHEq8pDBYDY6gcdvA4PfSvgNv2GDABD43Yy9lc88cQTryoPEQBirB908Do89gBQxy8wAND0VzzxxBOvKg/xNMAY6wcdvA6POQDU9QsKAFT9FU888cSrygMEgG7Uj6F5ZXisAaAJv0nmvhnK23w+eRRjf8UTTzzxqvKSzAa7EZDXJ9B+G+GxNK8MjzEANOU3zexNwSavcU9m7K944oknXlVemttfC7WH+hds16H91uYxNa8Mjy0ANOnXH8qfDRcA7GVN9IN9vognnnj94SXGvSLgi6h3o/3W4rE1rwyPKQA07dcnyo8Em7yZe2MT/WCfL+KJJ15/eP6FzfUB99A3o/1W5jE2rwyPJQC04TfN7LWhvBWfNzjVuaPq9oN9vognnnj94G3Jzz7N7227g50Pmf2DaOvH1ryyPIYA0JZfHwBeF8rbqoz98br9qONXPPHEE68pXsj3/1flx0P6rcVja15ZHjoAtOl3aNyvhJzASeY+PxjsPKJOP+r4FU888cRrgpfm+fGpcbeG3D+H+WQXym9tHlPzZuEhA0DbfufM+HEhJ/Ae/c86/ajjVzzxxBOvCV5q7OtD753JaLwD5bc2j6l5s/BQASCE37nllSEgANyTjpZ/sGo/6vgVTzzxxKvLS3L7c8EP/8zdGfVnqFiaNysPEQBC+vXjfhkwmb85MO4HGPornnjiiVf2Km7848+Eu4Lvmca9B+G3MR5D86rwQgeA0H79YXxl6Mm8R/f4sZ+1vLx0GN1kFU888cTb6zo5z49OM/tS0F5Z3Ab4eSH9tsWDDl6FBwgAQf36RPsTqEm9Z2L/87wZP/6886ZHs01W8cQTr9+806fTI9PcPdXvk19B7pODkZvGWL/KF4uZkAFgLptcHNrvXJ6fnBp7N3Jy3xsE3L8V33P1f/5U8fZAkrvxXGbzWbSQW3vWknNrKv4+K0M88cTrL2+QjV3xGSW/7z89ydzV/uC/Hb43ZvbmhYWFw9s430LxoIPX4QEDQDC/PgC8Az3JJUmSpP3lA8Arm9zvmz4/yvCgg9fhgQJAUL/FDXrQk1ySJEnaX8PRZKXJ/b7p86MMDzp4HR4gAAT3W3zAxfu8DT3RJUmSpL1lP+O36E1N7vdNnx9leNDB6/AAAQDiNzXuRfjJLkmSJP1/mfHT2tjvQ/Ogg9fhoW8FHMrvwvLySamx/wGf8JIkSVLxweivrXfzH+bz8kA86OB1eOwBoNmw416OnvSSJEmSDwC5fWab+31IHnTwOjzmANC0361LSw8u7tKHnviSJEm9lrE3Li4u3rfN/R7Ki8UMawBoy6//OZ4Nn/ySJEn91g+H2O91+G/AYwwAbfq1dvkY/7N8kmABSJIk9VD2mlD7vQ7/DXhsASBE/eaX7Plp8dQ++EKQJEnqkYz71tzichpyv9fhfxAeUwAIWb/U2Cvgi0GSJKlPGrmnIPb71nmxmmEJAMHrt2vXoYlx74IvCEmSpH7oj2H7fdu8WM0wBABU85Id7pTiu6gEC0OSJKnL+tjCwvQY5H7fKi9WM+gAgG7eMHfb08zeQrBAJEmSuidjb5w3k9MZ9vvWeLGaQQYAluYNc/tQP0nvgC8USZKkLsm4b/i9dZFpv2+FF6sZVACgap6/imdkp5n9NnzBSJIkdUL2lq1mZcK43zfOi9UMIgDQNW/Pde9vAnxihS8cSZKkeJUY99Vkx8qIeb9vlBermdABAO13o2tuabLNJ9cb0QtIkiQpShl7wyBf2drG/szOgw5ehQcIALTNW+MtGHtmatzfwReSJElSRPLnyZu3bz/nAW3uz6w86OBVeSEDwFw2uRjttyzvvPOmR/ske3mqOwZKkiQdVP4cucvvl5f5rXRTiP2ZjQcdvA4PGABomncw3iC3j0gz+xn0ApMkSWJUkrkPbs1dhtifWXjQwevwQAGAqnkb8fI8P2yY2Uv1KGFJkqQ12VuKfbG4qypyf2bgQQevwwMEALrmleWli+Mzk8y+NjHuO/jFJ0mSFF7FCyGvF29ZnJzItD8jedDB6/AAAQDqtwlecVcrHwJerpsHSZLUG937Fennp3l+PPP+jOBBB6/DQ98KOLTfJnlb8vzYZGSfWDzf2vvbDV+gkiRJzeqeYn8r9rm1e/nHsj+H5EVrhj0AsNdvjbcwXh56j5cOjX1TktmbCRauJEnSzFp9QJqxbyge3buwvHwSYj+NjRetGeYAEEP91uOdcMJxmwdLE5Oa8RP8QroiydzVXh/x+vxqODD2bvQilySpnyo+w3TvPuQ+6/ei6/2/+2Ov5yRm/PjBaLzAtp9GyYvFDGsAiKV+VXmnOndUccOMNS0unvOALDv3xDz/roq/F/9+7/9fWYknnnji7c87/34M+1+neTGZYQwAMdVPPPHEE0888fCDV+CxBYDY6ieeeOKJJ554UZphCgAx1k888cQTT7ye82I1wxIAYq2feOKJJ554PefFaoYhAMCbJ5544oknnnhVebGaQQcAiuaJJ5544oknXlVerGaQAYCmeeKJJ5544olXlRerGVQAoGqeeOKJJ5544lXlxWoGEQDomieeeOKJJ554VXmxmgkdANB+xRNPPPHEE68NHnTwKjxAAKBtnnjiiSeeeOK1frGYCRkA5rLJxWi/4oknnnjiidckDzp4HZ4/mD8W7DcAuf1ZtF/xxBNPvCZ5iZk8JMncdRvoQ+too/9mP6XZ+Fq0X/H250EHr8NLR/YDoQJAatwVaL/iiSeeeE3yksz+ZKg91I/1FbRf8fbnQQevw0sze02w3wBk7vVov+KJJ554TfIS4y4PFgCM/We0X/H250EHr8NLRvb/BEyvXzjjjNM2If2KJ5544jXJ84fy34fbQ911aL/i7c+DDl6H5yfvq0NN3tUJvGNlhPQrnnjiidcU70xrH+j3tXvC7aH27V2qX1d40ZrxifIXgwaAzL4G6Vc88cQTrymefwH1vJD7Z2rc73Spfl3hRWtmPrMXhZ3A9u6tO5bPQvkVTzzxxGuCt2VxcqIPALeF3D/3fJMK4le8GXixmDlrvLItaAC4NwS87/Tp9EiEX/HEE0+8Bnib08y+JfjemY3PA/kVr2uHf8E577zp0aFTbKEkc1f6H2XdDwTGVD/xxBOvf7zEuN8If/i7e9I8Px7hV7wOHv5rSo37x78VdQAABoVJREFUS8BkLkLAn5yc50fHXj/xxBOvJ7xduw71+9YLEPtlmtl/DO5XvG4f/sXfk9w+EzOhV0PAB+cym8dcP/HEE6/7vOHSyjDJ7F/j9kr7WyH9iteDw7/494OliUFN6j3a7dPtH6XGPXzxwkceFVv9xBNPvM7yNiW5G6fZ+HcTY+9C7pODbPLoCOvXD17kZjb5V+KfB4eAe1Oucf/q/3zzMLMv9wvu2amxl1XR0Nhn7auqLPHEE68/PL8X+v/NvcT/ebX/+43oPXFVxt16qnNHkZ4f4sVuJuTtLCVJkqTyKnv/FIrDsI+82M1sNW6Qrv4qHj/ZJUmSpO9qMHJT5vOj97wumEmMew96okuSJEnf1SBzn/Pb82b286PXvC6YGRr7GPRklyRJkvZSbi+J4fzoNa8jZjalxn4IPuElSZKk4kPRXzvYh//Izo/e86CDN8HzAeBH0JNekiRJcsWn/38+pvOjzzzo4A3yNieZ/Qf4xJckSeqzjPvsgV79E58fveRBB2+aV9yZLzX2bvgCkCRJ6qkOdOMf9vOjjzzo4G3wfPr8XfQCkCRJ6qOSzL0p5vOjbzzo4G3w5penJ6TGfhG9ECRJkvqkJLM3b8nPPi3m86NvPOjgbfHmltyKT6J3oheEJElST7Tb77kXdeH86BMPOnibPD8Zf5FgUUiSJHVe6z3xL+bzoy+8TpnZh7fJT8w/Qy8MSZKkbmt87cLCwuEz7s/s50cveJ0ysy+vmJSpce/ELxBJkqTuKTH2o9u3n/MAhv1evAZ4nTLjr8F4fH/dJVCSJKlxfXnr0tKDmfZ78WrwOmVmr2vL4uREHwKuJ1gwkiRJ8cu4zxZPYmXc78WrwOuUmXWuhYXpMT4EvAO+cCRJkiJW8Wv/ZIc7hXm/F0+H/37XYLDziCRzf4JeQJIkSTEqyexfF2+rxrDfi6fDf11emtlLdJ8ASZKk0tqdGveSPM8P23c/Zd/vxTsIr1NmZuCdteRcauwNBAtLkiSJWPYmf/jvjHm/F+/A/1F3zMzIW15e+r7E2OfqtwGSJEn7qbi735XFh6i7sN+Lt/5/2B0zFXnDpZVh8d4WwYKTJEmCyx/8H0lzdzbD/ixei7xOmanJS0bu/DQbvxe9+CRJkkD6WDKyT5xOp/dh25/Fa4HXKTMN8VIzeXhi7NtTY+8mWJCSJEltandi3Hv2PMxnE/v+LF6DvE6ZaZg3NCsPGmb2Ut1JUJKkzsnYL/pD/wXzI5vEuD+L1xwPOngMvG1LZ58+vzT50WHmXp0a9yX44pUkSZpBibG3pZm9xh/8l81lNj/kIK/2295PxePhQQePlZeMzjkhycbnprl7qg8Ev+kX16u9/rR468Dr73yyvu4g+tA6Otj/fyOJJ554Peetfn6pePiZsVf7g/51Xv/b//unzeXjRxa/zWTeT8XD8KCDiyeeeOKJJ554GB50cPHEE0888cQTD8ODDi6eeOKJJ5544mF40MHFE0888cQTTzwMDzq4eOKJJ5544omH4XXKjHjiiSeeeOKJV47XKTPiiSeeeOKJJ15FXqfMiCeeeOKJJ554OvzFE0888cQTTzymwcUTTzzxxBNPPB3+4oknnnjiiSdey7xOmRFPPPHEE0888crxOmVGPPHEE0888cQrx+uUGfHEE0888cQTrxyvU2bEE0888cQTT7xyvE6ZEU888cQTTzzxSgO6Y0Y88cQTTzzxxCsPgQ0unnjiiSeeeOJBedDBxRNPPPHEE0+88Dzo4OKJJ5544oknXngedHDxxBNPPPHEEw/Dgw4unnjiiSeeeOJheNDBxRNPPPHEE088DA86uHjiiSeeeOKJh+FBBxdPPPHEE0888TC8TpkRTzzxxBNPPPHK8TplRjzxxBNPPPHEq8jrlBnxxBNPPPHEE0+Hv3jiiSeeeOKJxzS4eOKJJ5544omnw1888cQTTzzxxGuZ1ykz4oknnnjiiSdeOV6nzIgnnnjiiSeeeOV4nTIjnnjiiSeeeOKV43XKjHjiiSeeeOKJV47XKTPiiSeeeOKJJ15pQHfMiCeeeOKJJ5545SGwwcUTTzzxxBNPPCgPOrh44oknnnjiiReeBx1cPPHEE0888cQLz4MOLp544oknnnjiYXjQwcUTTzzxxBNPPAwPOrh44oknnnjiiYfhQQcXTzzxxBNPPPEwPOjg4oknnnjiiScehtcpM+KJJ5544oknXjlep8yIJ5544oknnnjleP8POpmz8KpHD40AAAAASUVORK5CYII="))
		assert.False(t, validator.IsBase64DataType(`PD94bWwgdmVyc2lvbj0iMS4wIiA/Pjxzdmcgd2lkdGg9IjQ4cHgiIGhlaWdodD0iNDhweCIgdmlld0JveD0iMCAwIDQ4IDQ4IiBkYXRhLW5hbWU9IkxheWVyIDEiIGlkPSJMYXllcl8xIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjx0aXRsZS8+PHBhdGggZD0iTTgsMUEyLDIsMCwwLDAsNiwzVjQ1YTIsMiwwLDAsMCw0LDBWM0EyLDIsMCwwLDAsOCwxWiIvPjxwYXRoIGQ9Ik00My41NSwxMy43NEMzOC4yMiw3LjE4LDMyLjcxLDcuNjIsMjcuODQsOGMtNC42My4zNy04LjI5LjY2LTEyLjI5LTQuMjdBMiwyLDAsMCwwLDEyLDVWMjJhMiwyLDAsMCwwLC45NCwxLjcsOS4wOSw5LjA5LDAsMCwwLDQuOTEsMS40NmM0LDAsNy44LTIuNjIsMTEuMjgtNSw1LjE0LTMuNTMsOC40OS01LjUyLDExLjgxLTMuNDVhMiwyLDAsMCwwLDIuNjEtM1pNMjYuODcsMTYuODVDMjIuMjIsMjAsMTksMjIsMTYsMjAuNzhWOS42NmM0LjE4LDMsOC4zNywyLjYzLDEyLjE2LDIuMzMsMi41NC0uMiw0Ljc5LS4zOCw3LC4zMUMzMi4yMywxMy4xNywyOS40NiwxNS4wNywyNi44NywxNi44NVoiLz48L3N2Zz4=`))
	})

	t.Run("IsURL is valid", func(t *testing.T) {
		assert.True(t, validator.IsURL("https://dictionary.cambridge.org/dictionary/english/mock"))
	})

	t.Run("IsURL is not valid", func(t *testing.T) {
		assert.False(t, validator.IsURL("https:dictionary.cambridge.org/dictionary/english/mock"))
		assert.False(t, validator.IsURL(`httpx://dictionary.cambridge.org/dictionary/english/mock`))
	})
}
