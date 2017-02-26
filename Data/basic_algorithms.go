package data

/*
	FilterIn
		Filters map within two keys

		Inputs:
			early:	(int) - first key
			late	(int) - last keys
			marketData	(map[int]map[string]float64) - data to filter on

		Outputs:
			finalData	(map[int]map[string]float64) - filtered map
*/
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

/*
	JoinOn
		Joins data structures in a slice into one structure on specified
		key. Equivalent to SQL LEFT JoinOn

		Inputs:
			on	(string) - key to join the maps on
			data	([]map[string][]map[string]float64) - slice of maps to be
					joined

		Outputs:
			flattenedData	(map[int]map[string]float64) - result of the join
*/
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
