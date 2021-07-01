package _package

import "fmt"

// print One-dimensional array
func print1_tree(result []int) {
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

// print Two-dimensional array
func print2_tree(result [][]int) {
	for _, list := range result {
		for _, val := range list {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
}



