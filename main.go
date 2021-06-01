package main

import "fmt"

func main() {

	a := 7
	b := 5
	c := 6
	d := 4

	slices := []float64{2.00, 3.00, 5.00, 9.00, 12.00, 9.00, 7.00, 3.00, 2.00, 19.00, 17.00, 15.00, 10.00, 5.00, 2.00}

	sliceA := calcAverages(slices, a)
	sliceB := calcAverages(slices, b)
	sliceC := calcAverages(slices, c)
	MACDline := calcMACDline(sliceB, sliceC, c)
	signalLine := calcSignalLine(slices, MACDline, c, d)

	printSlice("Averages for a", sliceA)
	printSlice("Averages for b", sliceB)
	printSlice("Averages for c", sliceC)
	printSlice("calcMACDline", MACDline)
	printSlice("signalLine", signalLine)

}

func simpleMovingAverage(slices []float64, a int) float64 {
	total := 0.00
	for _, slice := range slices {
		total = total + slice
	}

	return total / float64(a+1)
}

func exponentialMovingAvarage(currentPrice, simpleAverage float64, a int) float64 {
	ka := float64(2) / float64(a+1)
	la := float64(1 - ka)
	return ka*currentPrice + (la * simpleAverage)
}

func calcAverages(prices []float64, a int) []float64 {
	var result []float64
	var simpleAverage, exponentialAvarage float64
	temp := 0.00

	for i, price := range prices {
		if i == a {
			// fmt.Printf("i = %v, value= %f \n", i, slice)
			simpleAverage = simpleMovingAverage(prices[:i+1], a)
			// fmt.Printf("simple Moving Average %v \n", simpleAverage)
		}

		if temp == 0.00 {
			temp = simpleAverage
		}

		if i > a {
			// fmt.Printf("i = %v, value= %f \n", i, slice)
			exponentialAvarage = exponentialMovingAvarage(price, temp, a)
			// fmt.Printf("exponential Moving Average %f \n", exponentialAvarage)
			temp = exponentialAvarage
		}
		result = append(result, temp)
	}

	return result
}

func calcMACDline(slicesB, slicesC []float64, c int) []float64 {
	var result []float64

	for i := range slicesB {
		temp := 0.00
		if i > c {
			temp = slicesB[i] - slicesC[i]
		}
		result = append(result, temp)
	}

	return result
}

func calcSignalLine(slices, MACDline []float64, c, d int) []float64 {
	var result []float64

	// fmt.Printf("c %d \n", c)
	// fmt.Printf("d %d \n", d)

	for i := range slices {

		if i-c == d {
			r := 0.00
			for j, mlSlice := range MACDline {
				if j > i {
					break
				}
				r = r + mlSlice
			}
			x := float64(r) / float64(d)
			result = append(result, x)
		} else if i-c > d {
			kd := float64(2) / float64(d+1)
			ld := 1 - kd

			currentMACDline := MACDline[i]
			previousSignalLineValue := result[i-1]

			z := kd*currentMACDline + ld*previousSignalLineValue
			result = append(result, z)

		} else {
			result = append(result, 0.00)
		}

	}

	return result
}

func printSlice(title string, slices []float64) {
	fmt.Printf("%s:\n", title)
	for _, v := range slices {
		fmt.Printf("%.4f, ", v)
	}
	fmt.Println("\n=============================")
}
