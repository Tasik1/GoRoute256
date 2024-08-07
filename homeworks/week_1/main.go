package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(stringOpener(`qwe\4\n`))
}

func stringOpener(toOpen string) (string, error) {
	if len(toOpen) == 0 {
		return toOpen, nil
	}
	res := strings.Builder{}
	for i := 0; i < len(toOpen); i++ {
		curr := toOpen[i]

		if isDigit(curr) {
			return "", errors.New("некорректная строка")
		}
		if i+1 < len(toOpen) && isDigit(toOpen[i+1]) {
			num, _ := strconv.Atoi(string(toOpen[i+1]))
			res.WriteString(strings.Repeat(string(curr), num))
			i++
		} else {
			res.WriteByte(curr)
		}

	}
	return res.String(), nil
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
