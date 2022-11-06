package utils

import "time"

func getTime(input string) time.Time {
	const format = "15:04"
	t, _ := time.Parse(format, input)
	return t
}

func GetStartTime(class int) time.Time {
	var startTime time.Time
	switch class {
	case 1:
		startTime = getTime("9:00")
	case 2:
		startTime = getTime("10:45")
	case 3:
		startTime = getTime("13:15")
	case 4:
		startTime = getTime("15:00")
	}

	return startTime
}

func GetEndTime(class int) time.Time {
	var endTime time.Time
	switch class {
	case 1:
		endTime = getTime("10:30")
	case 2:
		endTime = getTime("12:15")
	case 3:
		endTime = getTime("14:45")
	case 4:
		endTime = getTime("16:30")
	}

	return endTime
}
