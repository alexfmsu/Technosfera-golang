package main

import (
	"bytes"
	"fmt"
	"strings"
)

type memoizeFunction func(int, ...int) interface{}

var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

func init() {
	fibonacci = Memoize(func(x int, items ...int) interface{} {
		if x <= 0 {
			panic("fibonacci(x) failed: x < 0")
		} else if x <= 2 {
			return 1
		}

		return fibonacci(x-1).(int) + fibonacci(x-2).(int)
	})

	rom := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	dec := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	romanForDecimal = Memoize(func(x int, items ...int) interface{} {
		if x < 0 || x > 3999 {
			panic("romanForDecimal(x) failed: x < 0 or x > 3999")
		}

		var buffer bytes.Buffer

		for i, d := range dec {
			mod := x / d

			x %= d

			if mod != 0 {
				buffer.WriteString(strings.Repeat(rom[i], mod))
			}
		}

		return buffer.String()
	})
}

func main() {
	fmt.Println("fibonacci(45) =", fibonacci(45).(int))

	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}

func Memoize(function memoizeFunction) memoizeFunction {
	mem := make(map[string]interface{})

	return func(x int, items ...int) interface{} {
		key := string(x)

		for _, i := range items {
			key += "," + string(i)
		}

		if value, exists := mem[key]; exists {
			return value
		}

		mem[key] = function(x, items...)

		return mem[key]
	}
}
