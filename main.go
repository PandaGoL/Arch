package main

import (
	"fmt"
)


func bSearch(a []int, b int) int{
	max := len(a) - 1
	min := 0
	step := 0

	for min <= max {
		half := min + (max - min) / 2
		if b == a[half]{
			step++
			fmt.Printf("Индекс: %d Число поиска: %d Кол-во шагов: %d \n", half, a[half], step)
			return 0
		} else if b < a[half]{
			step ++
			max = half - 1
		} else {
			step ++
			min = half + 1
		}
		
		
	}
	return 1
}




func main(){
	var n int
	var velue = make([]int, n, n)
	var number int

	fmt.Println("Введите размер числовой последовательности:")
	fmt.Scan(&n)
	fmt.Printf("Введите число для поиска 0<number<=%d:\n", n)
	fmt.Scan(&number)
	
for i := 0; i <= n; i++ {
	velue = append(velue, i)

	}

bSearch(velue, number)

}