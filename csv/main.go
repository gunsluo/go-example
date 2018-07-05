package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/pkg/errors"
)

func main() {
	//testXLSX()
	testCSV()
}

func testXLSX() {
	file, err := os.Open("test.xlsx")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows, err := readPositionsFromXLSX(file)
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		fmt.Println(i, "==>", len(row), row)
	}
}

func testCSV() {
	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows, err := readPositionsFromCSV(file)
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		fmt.Println(i, "==>", len(row), row)
	}
}

func readPositionsFromXLSX(r io.Reader, args ...interface{}) ([][]string, error) {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read Excel file")
	}
	if xlsx.SheetCount == 0 {
		return nil, errors.New("Excel file has no sheets")
	}

	return xlsx.GetRows(xlsx.GetSheetName(1)), nil
}

func readPositionsFromCSV(r io.Reader, args ...interface{}) ([][]string, error) {
	var delim string
	if len(args) > 0 {
		if str, ok := args[0].(string); ok {
			delim = str
		}
	}

	delim = strings.ToLower(strings.TrimSpace(delim))
	var delimRune rune
	switch delim {
	case ",", ";", "\t", "|", ".": // add more if needed
		delimRune = []rune(delim)[0]
	case "", "comma":
		delimRune = ','
	case "tab":
		delimRune = '\t'
	case "semicolon":
		delimRune = ';'
	default:
		return nil, fmt.Errorf("unrecognized delimiter parameter %s", delim)
	}

	csvReader := csv.NewReader(r)
	csvReader.Comma = delimRune
	csvReader.Comment = '#'
	csvReader.LazyQuotes = true

	return csvReader.ReadAll()
}
