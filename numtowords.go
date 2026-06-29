package numtowords

import "fmt"

// MaxNum is the maximum number that can be converted to words
const MaxNum = 999

// Convert converts a number to words
func Convert(number int) (string, error) {
	if number < 0 || number > MaxNum {
		return "", fmt.Errorf("can only convert numbers between 0 and %d", MaxNum)
	}
	if number == 0 {
		return "zero", nil
	}

	units := [20]string{"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tens := [8]string{"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety"}

	result := ""

	if number > 99 {
		hundredsindex := number / 100
		result += units[hundredsindex] + " hundred"
		number = number % 100
		if number == 0 {
			return result, nil
		}

		result += " and "
	}
	if number > 19 {
		tensindex := number/10 - 2
		result += tens[tensindex]
		number = number % 10
		if number == 0 {
			return result, nil
		}
		result += " "
	}

	result += units[number]

	return result, nil
}
