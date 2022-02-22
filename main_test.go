package main

import (
	"fmt"
	"testing"
)

type myStringSlice []string

//CALCULATOR TESTING
func TestCalculate(t *testing.T) {
	var test = []struct {
		input  []string
		output []float64
	}{
		{myStringSlice{"450.52-50-65.67-84.44", "23.44+775.32+89.10+100.59", "560/5/2.5/250.7", "3*76*4*52"}, []float64{250.41, 988.4500000000002, 0.17869964100518548, 47424}},
		{myStringSlice{"800-250-95-65", "2200+55+95+1000", "1500/5/5/2.5", "4.88*67.7*78.99*5.76"}, []float64{390, 3350, 24, 150315.26538240002}},
	}

	for _, val := range test {
		floatValue := Calculate(val.input...)
		for i := 0; i < len(val.output); i++ {
			if floatValue[i] != val.output[i] {
				t.Errorf("Expected output %v received %v", val.output, floatValue)
				fmt.Println(floatValue[i])
			}
		}
	}

}

//MULTIPLICATION TESTING
func TestMultiplication(t *testing.T) {
	var test = []struct {
		input  []float64
		output float64
	}{
		{[]float64{100, 60, 40, 5}, 1200000},
		{[]float64{500, 190, 210, 33.30}, 664335000},
	}
	for _, val := range test {
		floatVal := Multiply(val.input)
		if floatVal != val.output {
			t.Errorf("Multiply : Expected %v : output %v", val.output, floatVal)
			fmt.Println(floatVal)
		}
	}
}

//ADDITION TESTING
func TestAddition(t *testing.T) {
	var test = []struct {
		input  []float64
		output float64
	}{
		{[]float64{6000, 2000, 37.50, 78.50}, 8116},
		{[]float64{21.60, 66.66, 11.90, 8.20}, 108.36},
	}
	for _, val := range test {
		floatVal := Summation(val.input)
		if floatVal != val.output {
			t.Errorf("Substract : Expected %v : output %v", val.output, floatVal)
			fmt.Println(floatVal)
		}
	}
}

//SUBSTRCTION TESTING
func TestSubstraction(t *testing.T) {
	var test = []struct {
		input  []float64
		output float64
	}{
		{[]float64{5 - 100.20 - 130.67 - 500}, -725.87},
		{[]float64{6000 - 999.99 - 65 - 1000}, 3935.01},
	}
	for _, val := range test {
		floatVal := Substract(val.input)
		if floatVal != val.output {
			t.Errorf("Substract : Expected %v : output %v", val.output, floatVal)
			fmt.Println(floatVal)
		}
	}
}

//DIVISION TESTING

func TestDivision(t *testing.T) {
	var test = []struct {
		input  []float64
		output float64
	}{
		{[]float64{7000, 50, 1, 5}, 28},
		{[]float64{8000, 5, 2, 40.50}, 19.753086419753085},
	}
	for _, val := range test {
		floatVal := Divide(val.input)
		if floatVal != val.output {
			t.Errorf("Division : Expected %v : output %v", val.output, floatVal)
			fmt.Println(floatVal)
		}
	}
}
