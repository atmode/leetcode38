package main

import (
	"fmt"
)

func mapFunc(input string) [][]interface{} {
	var result [][]interface{}
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			result = append(result, []interface{}{count, string(input[i-1])})
			count = 1
		}
	}
	result = append(result, []interface{}{count, string(input[len(input)-1])})

	return result
}
func transform(n [][]interface{}) [][]interface{} {
	transformed := make([][]interface{}, len(n))
	for i, pair := range n {
		transformed[i] = []interface{}{pair[1], pair[0]}
	}
	return transformed
}
func combine(transform [][]interface{}) string {
	var result string
	for _, pair := range transform {
		for _, elem := range pair {
			result += fmt.Sprintf("%v", elem)
		}
	}
	return result
}
func main() {

	input := "223314444411"
	fmt.Println(mapFunc(input))
	fmt.Println(transform(mapFunc(input)))
	fmt.Println(combine(transform(mapFunc(input))))
}
