package utils

//去重
func RemoveDuplicates(s []string, s1 []string) []string {
	var ret []string
	for _, v := range s {
		found := false
		for _, v2 := range s1 {
			if v == v2 {
				found = true
				break
			}
		}
		if !found {
			ret = append(ret, v)
		}
	}
	return ret
}
