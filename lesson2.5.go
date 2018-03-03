package main

//if

var number int

func main() {
	if 100 < number {
		number++
	}

	if 100 < number {
		number++
	} else {
		number--
	}

	if test := 100 - number; 100 < number {
		number++
	} else {
		number--
	}

	if diff := 100 - number; 100 < diff {
		number++
	} else if 200 < diff {
		number--
	} else {
		number -= 2
	}
}
