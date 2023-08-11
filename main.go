package main

import "fmt"

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
func main() {
	// countdown(6)
	// fmt.Println("sum :",sum(10))
	fmt.Println("Merge Sort..")
	nums1 := []int{40, 30, 20, 10, 0, 0, 0}
	m := 4
	nums2 := []int{50, 20, 60}
	n := 3
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	unsorted := append(nums1, nums2...)
	sorted := mergeSort(unsorted)
	fmt.Println("sorted ::", sorted)
}
func mergeSort(items []int) []int {
	fmt.Println("view unsorted item ::", items)
	if len(items) < 2 {
		return items
	}
	if len(items)%2 != 0 {
		// odd number
		// len : 7/2+1 = 4
		first := mergeSort(items[:len(items)/2+1])
		second := mergeSort(items[len(items)/2+1:])
		return merge(first, second)
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}
func merge(a, b []int) []int {
	final := []int{}
	fmt.Println("##### view slice from merge function ->>", a, b)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for i < len(a) {
		final = append(final, a[i])
		i++
	}
	for j < len(b) {
		final = append(final, b[j])
		j++
	}
	return final
}


// func mergeSort(first,second []int,m,n int) []int{
// 	first = first[:m]
// 	fmt.Println(first)
// 	fmt.Println(len(first))
// 	var deviderL1 int
// 	if len(first) %2 == 1 {
// 		deviderL1 = len(first)/2+1
// 		l1:=first[:deviderL1]
// 		l2:=first[deviderL1:]

// 		fmt.Println("l1 ::",l1)
// 		fmt.Println("l2 ::",l2)
// 		if l1[0] > l1[1]{
// 			l1[0],l1[1] = l1[1],l1[0]
// 		}
// 		fmt.Println("new l1 ::",l1)
// 		if l2[0] > l2[1]{
// 			l2[0],l2[1] = l2[1],l2[0]
// 		}
// 		fmt.Println("new l2 ::",l2)
// 	}else{
// 		deviderL1 = len(first)/2
// 		l1:=first[:deviderL1]
// 		l2:=first[deviderL1:]
// 		fmt.Println("l1 ::",l1)
// 		fmt.Println("l2 ::",l2)
// 		if l1[0] > l1[1]{
// 			l1[0],l1[1] = l1[1],l1[0]
// 		}
// 		fmt.Println("new l1 ::",l1)
// 		if l2[0] > l2[1]{
// 			l2[0],l2[1] = l2[1],l2[0]
// 		}
// 		fmt.Println("new l2 ::",l2)
// 	}

// 	r1:=second[:len(second)/2+1]
// 	r2:=second[len(second)/2+1:]

// 	fmt.Println("r1 ::",r1)
// 	fmt.Println("r2 ::",r2)
// 	// second = second[:n]
// 	// fmt.Println(second)
// 	// fmt.Println(len(second))
// 	return merge(first,second)
// }

// func countdown(value int) {
// 	if value == 0 {
// 		fmt.Println("stop counting")
// 		return
// 	}
// 	fmt.Println("counting down :: ", value)
// 	countdown(value - 1)
// }
// func sum(number int) int {
// 	if number == 0 {
// 		return number
// 	}
// 	return number+sum(number-1)
// }

