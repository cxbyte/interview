package benchmarking

func loops2(loop int) int {
	var counter = 0

	for i := 0; i < loop; i++ {
		for j := 0; j < loop; j++ {
			counter++
		}
	}

	return counter
}

func loops3(loop int) int {
	var counter = 0

	for i := 0; i < loop; i++ {
		for j := 0; j < loop; j++ {
			for k := 0; k < loop; k++ {
				counter++
			}
		}
	}

	return counter
}
