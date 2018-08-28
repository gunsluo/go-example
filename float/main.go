package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	var temp float64 = 0
	var evaluationScore *float64
	evaluationScore = &temp

	v, err := getFloat(*evaluationScore)
	if err != nil {
		panic(err)
	}

	fmt.Println("->", v)
}

func getFloat(unk interface{}) (float64, error) {
	switch i := unk.(type) {
	case float64:
		return float64(i), nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int16:
		return float64(i), nil
	case int8:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint16:
		return float64(i), nil
	case uint8:
		return float64(i), nil
	case int:
		return float64(i), nil
	case uint:
		return float64(i), nil
	case string:
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return math.NaN(), err
		}
		return f, err
	default:
		return math.NaN(), errors.New("evaluation score is of incompatible type")
	}
}
