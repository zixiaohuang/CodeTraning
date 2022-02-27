package main

func plusOne(digits []int) []int {
	n := len(digits) - 1
	for i := n; i >= 0; i--{
		if(digits[i] != 9) {
			digits[i]++;
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	res := make([]int, n + 1)
	res[0] = 1
	return res
}
