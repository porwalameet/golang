package calc

// returns sums of two integers

func Add(number ...int ) int {
	sum := 0
	for _, num := range number {
		sum += num
	}

	return sum
}