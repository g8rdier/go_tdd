package v1

import "strings"

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		case arabic > 0:
			result.WriteString("I")
			arabic -= 1
		}
	}

	return result.String()
}
