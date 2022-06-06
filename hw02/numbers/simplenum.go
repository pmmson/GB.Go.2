// Package numbers contains a number of methods for working with any numbers
//
// The SimpleN function returns a slice of prime positive numbers from 0 to n
//
// SimpleN(n int) []int
//
package numbers

// SimpleN returns a slice of prime positive numbers from 0 to n
func SimpleN(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var cntdiv int
	var numbers []int
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i%j != 0 {
				continue
			}
			cntdiv++
		}
		if cntdiv > 2 {
			cntdiv = 0
			continue
		}
		numbers = append(numbers, i)
		cntdiv = 0
	}
	return numbers
}
