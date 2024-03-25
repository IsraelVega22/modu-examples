package slices

func Filter[T int](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums { //{2, 12, 11, 99, 10}
		if callback(num) {
			result = append(result, num)
		}
	}
	//devuelve el slice con los numeros mayores
	return result
}
