package helpers

// ----------> If a field is empty return true otherwise return false
func IsItEmpty(value string) bool {
	if len(value) == 0 {
		return true
	}
	return false
}
