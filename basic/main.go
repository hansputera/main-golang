package main

import "log"

var (
	TESTCASES_JUMLAHTERUS    = []string{"123456789", "A123J82912"}
	TESTCASES_PALINDROME     = []string{"k!@!a#s()u$$rr#%u^s&%a*k", "#Al#p@!rO@*)A$sik*(k$is#A$O@r@$pl$!_A", "%#@$a$l%p#@rO%@S$u#@s!#a^h", "_*ma$!k@!an%#%n%$@a%k!@an"}
	TESTCASES_PRIMA          = []int{2, 5, 10}
	TESTCASES_HURUFJAGAJARAK = []string{"oooooooooooooooooooooooookkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"}
)

func main() {
	// Jumlah terus test
	for _, jumlahValueTerus := range TESTCASES_JUMLAHTERUS {
		log.Default().Printf("Test case jumlah terus '%v' hasilnya '%v'", jumlahValueTerus, JumlahTerus(jumlahValueTerus))
	}

	println("\n") // space

	// Palindrome test
	for _, palindromeValue := range TESTCASES_PALINDROME {
		log.Default().Printf("Test case palindrome '%v' hasilnya '%v'", palindromeValue, Palindrome(palindromeValue))
	}

	println("\n") // space

	// Prima test
	for _, primaValue := range TESTCASES_PRIMA {
		log.Default().Printf("Test case prima '%v' hasilnya '%v'", primaValue, Prima(primaValue))
	}

	println("\n") // space

	// Huruf jagaRak test
	for _, hurufJagaRakValue := range TESTCASES_HURUFJAGAJARAK {
		log.Default().Printf("Test case huruf jagaRak '%v' hasilnya '%v'", hurufJagaRakValue, HurufJagaJarak(hurufJagaRakValue))
	}

}
