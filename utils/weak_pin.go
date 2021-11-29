package utils

//GetWeakPin6Digit return array weak pin
func GetWeakPin6Digit() []string {
	// ref
	// https://www.datagenetics.com/blog/september32012/

	weakPinList := []string{
		"111111", "222222", "333333", "444444", "555555", "666666", "777777", "888888", "999999", "000000", // 1
		"121212", "131313", "696969", "112233", "101010", // 2
		"123123", "789456", "123321", "007007", // 3
		"123456", "654321", "123654", // 6
		"159753", "292513",
	}

	return weakPinList
}
