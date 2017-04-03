package goutil

func SubStr(str string, start int, end int) string {
	tempStr := []rune(str)
	length := len(tempStr)

	if start >= length && end >= length {
		return ""
	}

	if start > end {
		temp := start
		start = end
		end = temp
	}

	if end > length {
		end = length
	}

	retRune := tempStr[start:end]
	return string(retRune)
}

func SubStrByOffset(str string, start int, offset int) string {
	tempStr := []rune(str)
	length := len(tempStr)

	if start > length {
		return ""
	}

	if length-start <= offset {
		offset = length - start
	}

	retRune := tempStr[start : start+offset]

	return string(retRune)
}
