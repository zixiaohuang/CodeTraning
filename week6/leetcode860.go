package main

func lemonadeChange(bills []int) bool {
	changes := make(map[int]int)
	val := []int{20, 10, 5}
	var exchange func(int) bool
	exchange = func(money int) bool {
		for _, v := range(val) {
			for money >= v && changes[v] > 0 {
				money -= v
				changes[v]--
			}
		}
		return money == 0
	}

	for i := 0; i < len(bills); i++ {
		changes[bills[i]]++
		if !exchange(bills[i] - 5) {
			return false
		}
	}
	return true
}

func main() {
	lemonadeChange([]int{5,5,10,20,5,5,5,5,5,5,5,5,5,10,5,5,20,5,20,5})
}