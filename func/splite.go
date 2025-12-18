package asciiart

import (
	"os"
	"strings"
)

func Splite(fileName string) [][]string {
	var all []string
	var res [][]string
	data, err := os.ReadFile(fileName)
	if err != nil {
		return [][]string{}
	}

	if fileName == "thinkertoy.txt" {

		all = strings.Split(string(data), "\r\n")
	} else {
		all = strings.Split(string(data), "\n")

	}

	for i := 1; i+8 < len(all); i += 9 {
		res = append(res, all[i:i+8])
	}
	return res
}
