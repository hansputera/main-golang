package main

func main() {
	data := []int{}
	for i := 0; i <= 10; i++ {
		data = append(data, i)
	}
	index := len(data) - 1
	data = append(data[:index], data[index+1:]...)
	println(data[10])
}
