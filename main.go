package main

import "fmt"

func main() {
	somethings := []int{23, 45, 67, 54, 66, 19, 56, 78, 89, 44, 11, 22, 33, 44, 55, 66, 77, 88, 99, 23, 34, 32, 23, 12}

	length := len(somethings)

	var chunks [3][]int

	amount := length / 3
	maxTotal := 0
	minTotal := 0
	maxSliceTotal := []int{}
	minSliceTotal := []int{}

	for i := 0; i < 3; i++ {
		iteration := i + 1
		finish := iteration * amount
		start := i * amount
		slice := somethings[start:finish]
		chunks[i] = slice
	}

	for key, slice := range chunks {
		iteration := key + 1
		fmt.Println("Kumpulan ke-", iteration)
		fmt.Println(slice)
		max := 0
		min := 0
		total := 0
		for _, value := range slice {
			total += value
			if min == 0 || value < min {
				min = value
			}
			if value > max {
				max = value
			}
		}
		average := total / len(slice)
		fmt.Println("Rata-rata kumpulan adalah", average)
		fmt.Println("Nilai paling tinggi adalah", max)
		fmt.Println("Nilai paling rendah adalah", min)
		fmt.Println("Total Kumpulan adalah", total)
		fmt.Println("")
		if maxTotal < total {
			maxTotal = total
			maxSliceTotal = slice
		}
		if minTotal == 0 || minTotal > total {
			minTotal = total
			minSliceTotal = slice
		}
	}

	fmt.Println("Total Paling Tinggi adalah", maxTotal, "oleh Kumpulan ", maxSliceTotal)
	fmt.Println("Total Paling Rendah adalah", minTotal, "oleh Kumpulan ", minSliceTotal)
}
