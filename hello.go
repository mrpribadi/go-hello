package go_hello

import "strconv"

func Hello(Name string, Age int64) string {
	return "Hello, "+Name+", I'm "+strconv.FormatInt(Age, 10)+" years old"
}
