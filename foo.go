package foo

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("%s, 你好！version v0.3.0", name)
}
