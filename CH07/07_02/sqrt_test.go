package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calc %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("vad value %f", val)
	}
}

type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatal("error")
			}
			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}

func TestManyFromCSV(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s\n", record)
		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			fmt.Print(err)
		}
		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%f %f\n", value, expected)
		t.Run(fmt.Sprintf("%f", value), func(t *testing.T) {
			out, err := Sqrt(value)
			if err != nil {
				t.Fatal("error")
			}
			if !almostEqual(out, expected) {
				t.Fatalf("%f != %f", out, expected)
			}
		})
	}
}
func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
