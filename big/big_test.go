package big

import (
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	for n := -100; n <= 100; n++ {
		expected := strconv.Itoa(n)
		actual := New(n).String()
		if actual != expected {
			t.Errorf("Expected '%s', found '%s'\n", expected, actual)
		}
	}
}

func TestPlus(t *testing.T) {
	t.Parallel()

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			expected := strconv.Itoa(a + b)
			actual, err := New(a).Plus(New(b))
			if err != nil {
				t.Error(err)
				continue
			}

			if actual.String() != expected {
				t.Errorf("Expected '%s', found '%s', a=%d, b=%d\n", expected, actual, a, b)
			}
		}
	}
}

func TestTimes(t *testing.T) {
	t.Parallel()

	for a := -100; a <= 100; a++ {
		for b := -100; b <= 100; b++ {
			expected := strconv.Itoa(a * b)
			actual := New(a).Times(New(b))

			if actual.String() != expected {
				t.Errorf("Expected '%s', found '%s', a=%d, b=%d\n", expected, actual, a, b)
			}
		}
	}
}
