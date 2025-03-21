package main
// https://projecteuler.net/problem=92 link of the question
import "fmt"

func main() {
	result := 0
	for i := 1; i < 10000001; i++ {
		result += sdf(i)
	}
	fmt.Println(result)

}

func sdf(x int) int {
	for true {
		if x == 1 {
			return 0
		}
		if x == 89 {
			return 1
		}
		temp := 0
		for x > 0 {
			a := x % 10
			temp += a * a
			x /= 10
		}
		x += temp
	}
	return 0

}