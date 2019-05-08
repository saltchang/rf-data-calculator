package exwr

import (
	"fmt"
	"math"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Height3to15ExcelWriter func
func Height3to15ExcelWriter(data [8][]float64) {
	f, err := excelize.OpenFile("./model/height3to15km.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowChar := []string{"B", "C", "D", "E", "F", "G", "H", "I"}
	rowNum := 61

	totalsum := 0.0
	for index, row := range rowChar {
		sum := 0.0
		for i := 0; i < rowNum; i++ {
			height := data[index][i]
			sum += height
			cellName := row + fmt.Sprintf("%d", 4+i)
			fmt.Println(cellName, ": ", height)
			f.SetCellValue("高度估算表", cellName, height)
		}
		avg := sum / float64(rowNum)
		avg = math.Round(avg*100) / 100.0
		totalsum += avg
		finalCellName := row + fmt.Sprintf("%d", rowNum+4)
		fmt.Println(finalCellName, "平均高度: ", avg)
		f.SetCellValue("高度估算表", finalCellName, avg)
	}
	totalavg := math.Round(totalsum/8.0*100) / 100.0
	totalAvgCellName := fmt.Sprintf("B%d", rowNum+1+4)
	f.SetCellValue("高度估算表", totalAvgCellName, totalavg)

	// Save xlsx file by the given path.
	err = f.SaveAs("./test3to15.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

// Height10to50ExcelWriter func
func Height10to50ExcelWriter(data [8][]float64) {
	f, err := excelize.OpenFile("./model/height10to50km.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rowChar := []string{"B", "C", "D", "E", "F", "G", "H", "I"}
	rowNum := 41

	totalsum := 0.0
	for index, row := range rowChar {
		sum := 0.0
		for i := 0; i < rowNum; i++ {
			height := data[index][i]
			sum += height
			cellName := row + fmt.Sprintf("%d", 4+i)
			fmt.Println(cellName, ": ", height)
			f.SetCellValue("高度估算表", cellName, height)
		}
		avg := sum / float64(rowNum)
		avg = math.Round(avg*100) / 100.0
		totalsum += avg
		finalCellName := row + fmt.Sprintf("%d", rowNum+4)
		fmt.Println(row, "平均高度: ", avg)
		f.SetCellValue("高度估算表", finalCellName, avg)
	}
	totalavg := math.Round(totalsum/8.0*100) / 100.0
	totalAvgCellName := fmt.Sprintf("B%d", rowNum+1+4)
	f.SetCellValue("高度估算表", totalAvgCellName, totalavg)

	// Save xlsx file by the given path.
	err = f.SaveAs("./test10to50.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
