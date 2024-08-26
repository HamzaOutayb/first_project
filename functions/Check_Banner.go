package functions

func Check_Banner(Banner string) bool {
	if Banner == "standard" || Banner == "shadow" || Banner == "thinkertoy" {
		return true
	}
	return false
}
