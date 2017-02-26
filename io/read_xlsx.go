//Package io performs basic input output to file/console for the project
package io

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

//ReadFile takes filenames and load them into slice of maps keyed by the
// workbook name and valued by the resultant map of data
func ReadFile(filename string) map[string][]map[string]float64 {
	excelFileName := filename
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}

	workBook := make(map[string][]map[string]float64, 0)

	for _, sheet := range xlFile.Sheets {
		firstLine := true
		titles := make([]string, 0, 5000)
		var data = make([]map[string]float64, 0, 10000)
		for _, row := range sheet.Rows {
			title := 0
			var dataPoint = make(map[string]float64, 0)
			for _, cell := range row.Cells {
				text, _ := cell.String()
				if firstLine {
					titles = append(titles, text)
				} else {
					cellValue, err := strconv.ParseFloat(text, 64)
					//fmt.Println("CELLVALUE", cellValue)
					if err != nil {
						fmt.Println("ERROR")
						dataPoint[titles[title]] = 0
					} else {
						dataPoint[titles[title]] = cellValue
					}
				}
				title++
			}
			if firstLine {
				firstLine = false
			} else {
				data = append(data, dataPoint)
			}
		}
		workBook[sheet.Name] = data
	}
	return workBook
}
