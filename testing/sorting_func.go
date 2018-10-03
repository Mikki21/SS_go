/* signature
func ([] int ) []int
*/

package main

import "fmt"

// BubbleSort sorting function for slice of int.
func BubbleSort(a []int) []int {

	wasChanges := true
	for wasChanges {
		wasChanges = false
		for i := range a {

			j := i + 1
			if j == len(a) {
				break
			}
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				wasChanges = true
			}
		}
	}
	return a
}

/*--------------------------------------Quick sort------------------------------*/
func Quicksort(elements []int, low int, high int) {
	if high > low {
		pivot := Partition(elements, low, high)
		Quicksort(elements, low, pivot-1)
		Quicksort(elements, pivot+1, high)
	}
}

func Partition(elements []int, low int, high int) int {
	fmt.Printf("partition: %d, %d\n", low, high)
	pivot := high
	i := low - 1
	for j := low; j < pivot; j++ {
		if elements[j] < elements[pivot] {
			i++
			elements[i], elements[j] = elements[j], elements[i]
		}
	}
	elements[i+1], elements[pivot] = elements[pivot], elements[i+1]
	fmt.Println(elements)
	return i + 1
}

//________________________________Select_sort_________________________________________________

func Min(a []int) (k int) {
	mum := a[0]
	for i := range a {
		if a[i] < mum {
			mum = a[i]
			k = i
		}
	}
	return
}

func SelectSort(a []int) []int {
	k := len(a)
	s := make([]int, 0)
	loop := true
	for loop {
		for i := 0; i < k; i++ {
			s = append(s, a[Min(a)])
			a = append(a[:Min(a)], a[Min(a)+1:]...)
		}
		loop = false
	}
	a = append(a, s[:]...)
	return a
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SelectSort1(a []int) []int {
	k := len(a)
	for i := 0; i < k; i++ {
		indMin := i
		for j := i; j < k; j++ {
			if a[j] < a[indMin] {
				indMin = j
			}
		}
		a[i], a[indMin] = a[indMin], a[i]
	}
	return a
}

func main() {
	s := []int{-6, 4, 5, 3, 2, 1, -5}
	fmt.Println(s, "\n\n")
	fmt.Println(SelectSort(s))
	fmt.Println(SelectSort1(s))
}
