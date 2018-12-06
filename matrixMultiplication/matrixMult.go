package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transpose(x [10][10]float32) [][]float32 {
	out := make([][]float32, len(x[0]))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func dot(x [10][10]float32, y [][]float32) ([][]float32, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("Can't do matrix multiplication")
	}

	out := make([][]float32, len(x), len(y[0]))
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(x[0]); k++ {
				if len(out[i]) < 1 {
					out[i] = make([]float32, len(y))
				}
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return out, nil
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	/* 	X := [][]float32{
		[]float32{1.0, 2.0, 3.0},
		[]float32{2.0, 2.0, 2.0},
		[]float32{-1.0, 0.0, 1.0},
	} */

	file, err := os.Open("file.txt")
	check(err)
	defer file.Close()

	var X [10][10]float32

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines := scanner.Text()
		s := strings.Split(lines, " ")
		if len(s) > 0 {
			i, err := strconv.ParseInt(s[0], 10, 32)
			j, err := strconv.ParseInt(s[1], 10, 32)
			k, err := strconv.ParseFloat(s[2], 32)

			X[i][j] = float32(k)

			check(err)
		}

	}
	fmt.Println(X)

	out, err := dot(X, transpose(X))

	if err != nil {
		panic(err)
	}

	fmt.Println(out)

}
