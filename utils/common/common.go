package common

func InterfaceSliceContain(slices []interface{}, target string) bool {
	for _, slice := range slices {
		if target == slice.(string) {
			return true
		}
	}
	return false
}
