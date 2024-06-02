package helper

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func CreateEstate(width int, length int) [][]int {
	// initiate a empty 2d array
	estate := [][]int{}

	for i := 0; i < length; i++ {
		// we fill each row an array of 0s with the width size
		l := make([]int, width)
		// insert it into estate
		estate = append(estate, l)
	}

	return estate
}
