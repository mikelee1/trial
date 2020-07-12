package string

func ListContain(list []string, str string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}
	return false
}
