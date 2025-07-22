package v1

import "strings"

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic >= 9:
			result.WriteString("IX")
			arabic -= 9
		case arabic >= 5:
			result.WriteString("V")
			arabic -= 5
		case arabic >= 4:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic -= 1
		}
	}

	return result.String()
}
