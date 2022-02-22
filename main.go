package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(Calculate("2*3*4*5", "2+3+4+5.9+6+7.8", "20.54-7.65-2.897", "25/5/2"))
	fmt.Println(Multiply([]float64{2, 3, 4, 5}))
	fmt.Println(Summation([]float64{2, 3, 4, 5.9, 6, 7.8}))
	fmt.Println(Substract([]float64{20.54, 7.65, 2.897}))
	fmt.Println(Divide([]float64{25, 5, 2}))

}

func Calculate(parameters ...string) []float64 {

	var result []float64 = make([]float64, len(parameters))

	for i := 0; i < len(parameters); i++ {
		if strings.Contains(parameters[i], "*") {
			strSplit := strings.Split(parameters[i], "*")
			sliceNum := []float64{}

			for i := 0; i < len(strSplit); i++ {
				convToInt, _ := strconv.ParseFloat(strSplit[i], 64)
				sliceNum = append(sliceNum, convToInt)
			}
			result[i] = Multiply(sliceNum)
		}

		//ADDITION OPERATION
		if strings.Contains(parameters[i], "+") {
			strSplit := strings.Split(parameters[i], "+")
			sliceNum := []float64{}

			for i := 0; i < len(strSplit); i++ {
				convToInt, _ := strconv.ParseFloat(strSplit[i], 64)
				sliceNum = append(sliceNum, convToInt)
			}

			result[i] = Summation(sliceNum)

		}

		//SUBSTRACTION OPERATION

		if strings.Contains(parameters[i], "-") {
			strSplit := strings.Split(parameters[i], "-")
			sliceNum := []float64{}

			for i := 0; i < len(strSplit); i++ {
				convToInt, _ := strconv.ParseFloat(strSplit[i], 64)
				sliceNum = append(sliceNum, convToInt)
			}

			result[i] = Substract(sliceNum)
		}

		//DIVISION OPERATION

		if strings.Contains(parameters[i], "/") {
			strSplit := strings.Split(parameters[i], "/")
			sliceNum := []float64{}

			for i := 0; i < len(strSplit); i++ {
				convToInt, _ := strconv.ParseFloat(strSplit[i], 64)
				sliceNum = append(sliceNum, convToInt)
			}

			result[i] = Divide(sliceNum)

		}
	}

	return result

}

//MUTIPLICATION FLOAT OF SLICE OPERATION

func Multiply(multi []float64) float64 {

	var mult float64 = 1
	for i := 0; i < len(multi); i++ {
		mult = mult * multi[i]
	}

	return mult
}

//ADDITION FLOAT OF SLICE OPERATION
func Summation(sum []float64) float64 {

	var sumNumb float64
	for _, num := range sum {
		sumNumb += num
	}
	return sumNumb
}

//SUBSTRACTION FLOAT OF SLICE OPERATION

func Substract(minus []float64) float64 {
	var result float64 = minus[0]
	for i := 1; i < len(minus); i++ {
		result -= minus[i]
	}
	s, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", result), 64)
	return s
}

//DIVISION FLOAT OF SLICE OPERATION

func Divide(div []float64) float64 {
	var divNum float64 = div[0]

	for i := 1; i < len(div); i++ {
		divNum /= div[i]
	}

	return divNum
}
