package functions

import (
	"os"
	"strings"
	)
	
func Web_ascii_art(input string, Banner string) string {
	
	var builder strings.Builder
	var slice []string

	themplate, err := os.ReadFile("./src/" + Banner + ".txt")
	if err != nil {
		os.Exit(1)
	}
	if Banner == "thinkertoy" {
		slice = strings.Split(string(themplate), "\r\n")[1:]
	} else {
		slice = strings.Split(string(themplate), "\n")[1:]
	}
	Splitedinput := strings.Split(input, "\r\n")

	for _, words := range Splitedinput {
		if words == "" {
			builder.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range words {
					if char < 32 || char > 126 {
						continue
			}
				builder.WriteString(slice[int(char-' ')*9+i])
			}
			builder.WriteString("\n")
		}
	}
	return builder.String()
}