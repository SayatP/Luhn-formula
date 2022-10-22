package main

import "fmt"
import "os"

func main() {
	arg := os.Args[1]
	fmt.Println(Valid(arg))
}

func Valid(id string) bool {
	var converted_id []int32
	var invalid_id bool

	converted_id, invalid_id = convert(id)

	if invalid_id {
		return false
	}

	var sum int32

	var idx int
	if len(converted_id)%2 == 1 {
		idx = 1
	}

	for i := 0; i < len(converted_id); i++ {

		if i%2 == idx {
			converted_id[i] = doubleAndSub(converted_id[i])
		}

		sum += converted_id[i]
	}

	return sum%10 == 0

}

func doubleAndSub(i int32) (r int32) {
	r = i * 2
	if r > 9 {
		r -= 9
	}
	return
}

func convert(s string) (res []rune, invalid_id bool) {

	for _, char := range s {
		if isNumeric(char) {
			res = append(res, char-48)
		} else if char != 32 {
			invalid_id = true
		}
	}

	if len(res) < 2 {
		invalid_id = true
		return
	}

	return

}

func isNumeric(s rune) bool {
	return s > 47 && s < 58
}
