package helper

func StartWith(src,short string) bool {
	if len(src)< len(short) {
		return false
	}

	if src[:len(short)]==short{
		return true
	}
	return false
}
