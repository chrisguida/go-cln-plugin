package util

import "strconv"

func FormatMsat(balanceMsat int64) string {
	inputStr := strconv.FormatInt(balanceMsat, 10)
	length := len(inputStr)
	if length <= 3 {
		return inputStr
	}

	if length <= 6 {
		// 123.456
		return string(inputStr[:length-3]) + "." + string(inputStr[length-3:])
	}

	if length <= 9 {
		// 123,456.789
		return string(inputStr[:length-6]) + "," + string(inputStr[length-6:length-3]) + "." + string(inputStr[length-3:])
	}

	if length <= 11 {
		// 89|123,456.789
		return string(inputStr[:length-9]) + "|" + string(inputStr[length-9:length-6]) + "," + string(inputStr[length-6:length-3]) + "." + string(inputStr[length-3:])
	}

	if length <= 14 {
		// 567.89|123,456.789
		return string(inputStr[:length-11]) + "." + string(inputStr[length-11:length-9]) + "|" + string(inputStr[length-9:length-6]) + "," + string(inputStr[length-6:length-3]) + "." + string(inputStr[length-3:])
	}

	if length <= 17 {
		// 234,567.89|123,456.789
		return string(inputStr[:length-14]) + "," + string(inputStr[length-14:length-11]) + "." + string(inputStr[length-11:length-9]) + "|" + string(inputStr[length-9:length-6]) + "," + string(inputStr[length-6:length-3]) + "." + string(inputStr[length-3:])
	}

	// 20,234,567.89|123,456.789
	return string(inputStr[:length-17]) + "," + string(inputStr[length-17:length-14]) + "," + string(inputStr[length-14:length-11]) + "." + string(inputStr[length-11:length-9]) + "|" + string(inputStr[length-9:length-6]) + "," + string(inputStr[length-6:length-3]) + "." + string(inputStr[length-3:])
}
