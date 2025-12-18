package asciiart

import (
	"strings"
)

func PrintSymbole(arr [][]string, woord string) string {
	text := ""
	if woord == "" {
		return ""
	}

	str := woord

	words := strings.Split(str, "\r\n")

	for _, val := range words {
		if val == "" {
			text += "\n"
			continue
		} else {
			for i := 0; i < 8; i++ {
				for _, sVal := range val {
					if sVal < ' ' || sVal > '~' {
						continue
					} else {
						if sVal >= 0 && int(rune(sVal)-32) < len(arr) {
							text += arr[int(rune(sVal)-32)][i]
						}
					}
				}
				text += "\n"
			}

		}
	}
	return text
}
