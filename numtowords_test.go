package numtowords_test

import (
	"testing"

	"github.com/Samikshapaliwal/numtowords"
)

func TestInvalidNumber(t *testing.T) {
	_, err := numtowords.Convert(numtowords.MaxNum + 1)
	if err == nil {
		t.Log("Expected error for number greater than MaxNum, but got nil")
		t.Fail()
	}

	_, err = numtowords.Convert(-numtowords.MaxNum - 1)
	if err == nil {
		t.Logf("Expected error for number less than -%d, but got nil", numtowords.MaxNum)
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	result, err := numtowords.Convert(0)
	if err != nil {
		t.Log("Convert to zero failed")
		t.FailNow()
	}
	if result != "zero" {
		t.Logf("Expected 'zero' for number 0, but got '%s'", result)
		t.FailNow()
	}

}

func TestSingleDigits(t *testing.T) {
	for i := 1; i <= 9; i++ {
		result, err := numtowords.Convert(i)
		if err != nil {
			t.Logf("Convert to single digit %d failed", i)
			t.FailNow()
		}
		expected := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		if result != expected[i-1] {
			t.Logf("Expected '%s' for number %d, but got '%s'", expected[i-1], i, result)
			t.FailNow()
		}
	}
}

func TestTens(t *testing.T) {
	testcases := map[int]string{
		20: "twenty",
		33: "thirty three",
		19: "nineteen",
	}
	for number, expected := range testcases {
		t.Logf("Now Testing %v...", number)

		result, err := numtowords.Convert(number)
		if err != nil {
			t.Logf("Convert to number %d failed", number)
			t.FailNow()
		}
		if result != expected {
			t.Logf("Expected '%s' for number %d, but got '%s'", expected, number, result)
			t.FailNow()
		}
	}
}

func TestHundreds(t *testing.T) {
	testcases := map[int]string{
		109: "one hundred and nine",
		333: "three hundred and thirty three",
		700: "seven hundred",
		340: "three hundred and forty",
	}
	for number, expected := range testcases {
		t.Logf("Now Testing %v...", number)

		result, err := numtowords.Convert(number)
		if err != nil {
			t.Logf("Convert to number %d failed", number)
			t.FailNow()
		}
		if result != expected {
			t.Logf("Expected '%s' for number %d, but got '%s'", expected, number, result)
			t.FailNow()
		}
	}
}

func TestNegatives(t *testing.T) {
	testcases := map[int]string{
		-1:   "minus one",
		-19:  "minus nineteen",
		-20:  "minus twenty",
		-99:  "minus ninety nine",
		-100: "minus one hundred",
		-109: "minus one hundred and nine",
		-333: "minus three hundred and thirty three",
		-700: "minus seven hundred",
		-340: "minus three hundred and forty",
	}
	for number, expected := range testcases {
		t.Logf("Now Testing %v...", number)

		result, err := numtowords.Convert(number)
		if err != nil {
			t.Logf("Convert to number %d failed", number)
			t.FailNow()
		}
		if result != expected {
			t.Logf("Expected '%s' for number %d, but got '%s'", expected, number, result)
			t.FailNow()
		}
	}
}

func TestMinusZero(t *testing.T) {
	result, err := numtowords.Convert(-0)
	if err != nil {
		t.Logf("Convert to number -0 failed")
		t.FailNow()
	}
	if result != "zero" {
		t.Logf("Expected 'zero' for number -0, but got '%s'", result)
		t.FailNow()
	}
}
