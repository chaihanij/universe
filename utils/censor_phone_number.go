package utils

//CensorPhoneNumber replace Last Digit Phone Number
func CensorPhoneNumber(phoneNumber string, replaceDigit int) string {
	if replaceDigit < 0 {
		return phoneNumber
	}
	if len(phoneNumber) < replaceDigit {
		replaceDigit = len(phoneNumber)
	}
	phoneNumberRemoveLast := phoneNumber[:len(phoneNumber)-replaceDigit]
	for i := 0; i < replaceDigit; i++ {
		phoneNumberRemoveLast = phoneNumberRemoveLast + "X"
	}

	return phoneNumberRemoveLast
}
