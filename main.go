package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Nomor 1
	fmt.Println("Nomor 1:")
	array1 := []int{5, 7, 7, 9, 10, 4, 5, 10, 6, 5}
	array2 := []int{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5}
	array3 := []int{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	array4 := []int{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(socks(array1))
	fmt.Println(socks(array2))
	fmt.Println(socks(array3))
	fmt.Println(socks(array4))

	// Nomor 2
	fmt.Println("Nomor 2:")
	fmt.Println(countWords("Kemarin Shopia per[gi ke mall"))

}

func socks(ar []int) int {
    m := map[int]int{}
    
    for i := 0; i < len(ar); i++ {
		m[ar[i]] += 1
    }
    
    dup := 0
    
    for _, v := range m {
        dup += v / 2
    }
    
    return dup      
}

func countWords(s string) int {
	re := regexp.MustCompile(`[^a-zA-Z0-9\S]+`)
    
    results := re.FindAllString(s, -1)
    return len(results)
}