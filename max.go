package main

import "fmt"

func main() {
	var n uint64
	fmt.Scan(&n)
	var a []uint64
	for i := uint64(0); i < n; i++ {
		a = append(a, 0)
		fmt.Scan(&a[i])
	}
	mins := []uint64{a[0], a[1], a[2]}
	var max uint64 = findMax(mins)
	var m uint64 = 1
	for i := uint64(3); i < n; i++ {
		if mins[max] > a[i] {
			mins[max] = a[i]
			max = findMax(mins)
			m = 1
		} else if mins[max] == a[i] {
			m++
		}
	}
	fmt.Print(counter(mins, m, max))
}

func findMax(arr []uint64) uint64 {
	var max uint64 = 0
	for i := uint64(1); i < 3; i++ {
		if arr[max] < arr[i] {
			max = i
		}
	}
	return max
}

func counter(arr []uint64, n uint64, max uint64) uint64 {
	if n == 1 {
		return 1
	}
	n--
	var k uint64
	for i := range arr {
		if arr[i] == arr[max] {
			k++
			n++
		}
	}
	var f uint64 = 1
	var sum uint64 = 1
	for i := uint64(0); i < k; i++ {
		sum *= (n - uint64(i))
		f *= i+1
	}
	return sum/f
}
