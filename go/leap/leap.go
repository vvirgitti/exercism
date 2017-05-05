package leap

const testVersion = 2

func IsLeapYear(year int) bool {
	if (year % 400) == 0 {
		return true
	}
	if (year % 4) == 0 {
		if (year % 100) == 0 {
			return false
		} else {
			return true
		}
	}

	return false

}
