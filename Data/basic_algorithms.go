package data

import (
)

func FilterIn(early int, late int, marketData map[int]map[string]float64) map[int]map[string]float64 {
	finalData := make(map[int]map[string]float64, 0)
	keys := make([]int, 0)
	for year, data := range marketData {
		if year >= early && year <= late {
			keys = append(keys, year)
			finalData[year] = data
		}
	}

	return finalData
}


func JoinOn(on string, data []map[string][]map[string]float64) map[int]map[string]float64 {
	flattenedData := make(map[int]map[string]float64, 0)
	for _, excelFile := range data {
		for _, workbook := range excelFile {
			for _, row := range workbook {
				_, keyPresent := flattenedData[int(row[on])]
				if keyPresent == false {
						flattenedData[int(row[on])] = make(map[string]float64)
					}
				for key, column := range row {
					if key != on {
						flattenedData[int(row[on])][key] = column
					}
				}
			}
		}
	}

	return flattenedData
}
