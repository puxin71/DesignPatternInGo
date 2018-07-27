package sort

const (
	InsertionSortSize = 10000
)

// Some claims that O(N*N) works for up to 10,000 items
// Some says that the size should be <= 25
func InsertionSort(values []int) {
	if len(values) > InsertionSortSize {
		panic("exceed the recommended data size for insertion sort")
	}

	for i := 1; i < len(values); i++ {
		j := i
		// sort data before index i
		for j > 0 {
			if values[j-1] > values[j] {
				values[j-1], values[j] = values[j], values[j-1]
			}
			j--
		}
	}
}
