package utility

import "fmt"

func FormatCaller(file, function string, line int) string {
	return fmt.Sprintf("%s(%s:%d)", file, function, line)
}

func Bytes(input string) []byte {
	return []byte(input)
}
