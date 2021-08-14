package main

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Prima(n int) int {
	var temp []int = []int{}
	index := 0

	for len(temp) < n {
		if isPrime(index) {
			temp = append(temp, index)
		}
		index++
	}
	total := 0
	for _, v := range temp {
		total += v
	}
	return total
}
