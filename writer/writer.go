package writer

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/saltchang/rf-data-calculator/models"
)

// Height3to15ExcelWriter func
func Height3to15ExcelWriter(rf *models.RFDATA) {

	data := rf.Height.H3to15

	f, err := excelize.OpenFile("./templates/height3to15km.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	title := rf.Basic.Name + "廣播電臺(" + fmt.Sprintf("%.1f", rf.Basic.Fequency) + rf.Basic.FQUnit + ")天線設置點平均地形高度估算表"

	f.SetCellValue("高度估算表", "A1", title)

	rowChar := []string{"B", "C", "D", "E", "F", "G", "H", "I"}
	rowNum := rf.Height.N3to15

	for index, row := range rowChar {
		for i := 0; i < rowNum; i++ {
			height := data[index][i]
			cellName := row + fmt.Sprintf("%d", 4+i)
			if height > 0 {
				f.SetCellValue("高度估算表", cellName, height)
			} else {
				f.SetCellValue("高度估算表", cellName, "—")
			}
		}
		finalCellName := row + fmt.Sprintf("%d", rowNum+4)
		f.SetCellValue("高度估算表", finalCellName, rf.Height.Hav3to15[index])
	}
	totalAvgCellName := fmt.Sprintf("B%d", rowNum+1+4)
	f.SetCellValue("高度估算表", totalAvgCellName, rf.Height.Htoav3to15)

	// Save xlsx file by the given path.
	err = f.SaveAs("./test3to15.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

// Height10to50ExcelWriter func
func Height10to50ExcelWriter(rf *models.RFDATA) {

	data := rf.Height.H10to50

	f, err := excelize.OpenFile("./templates/height10to50km.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	title := rf.Basic.Name + "廣播電臺(" + fmt.Sprintf("%.1f", rf.Basic.Fequency) + rf.Basic.FQUnit + ")天線設置點平均地形高度估算表"

	f.SetCellValue("高度估算表", "A1", title)

	rowChar := []string{"B", "C", "D", "E", "F", "G", "H", "I"}
	rowNum := rf.Height.N10to50

	for index, row := range rowChar {
		for i := 0; i < rowNum; i++ {
			height := data[index][i]
			cellName := row + fmt.Sprintf("%d", 4+i)
			if height > 0 {
				f.SetCellValue("高度估算表", cellName, height)
			} else {
				f.SetCellValue("高度估算表", cellName, "—")
			}
		}
		finalCellName := row + fmt.Sprintf("%d", rowNum+4)
		f.SetCellValue("高度估算表", finalCellName, rf.Height.Hav10to50[index])

		// 寫入地形起伏度
		haatCellName := row + "47"
		f.SetCellValue("高度估算表", haatCellName, rf.Height.DeltaH[index])
	}
	totalAvgCellName := fmt.Sprintf("B%d", rowNum+1+4)
	f.SetCellValue("高度估算表", totalAvgCellName, rf.Height.Htoav10to50)

	// Save xlsx file by the given path.
	err = f.SaveAs("./test10to50.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
