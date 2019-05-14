package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/saltchang/rf-data-calculator/helper"

	"github.com/saltchang/rf-data-calculator/models"
	"github.com/saltchang/rf-data-calculator/writer"
)

var (
	xhaat  = 30.0
	xfield = 60.0

	rf models.RFDATA
)

func main() {

	fmt.Print("請輸入代號以選擇功能(1 : 產生高度估算表, 2 : 計算預估場強傳送距離): ")
	var function string
	fmt.Scanln(&function)
	if function == "1" {
		generateHeightSheet()
	} else if function == "2" {
		computeFMDistance()
	} else {
		log.Fatal("輸入的功能代號錯誤")
	}

	return
}

func generateHeightSheet() {
	fmt.Print("電台名稱(例: 地理中心): ")
	fmt.Scanln(&rf.Basic.Name)

	fmt.Print("電台類型(請輸入 'FM' 或 'AM'): ")
	fmt.Scanln(&rf.Basic.RFTYPE)
	if rf.Basic.RFTYPE != "FM" && rf.Basic.RFTYPE != "AM" {
		fmt.Println("Wrong RF type! (must be 'FM' or 'AM')..")
		os.Exit(1)
	}

	fmt.Print("電台頻率(例: 94.1): ")
	var rfFQ string
	fmt.Scanln(&rfFQ)
	rfFQfloat, err := strconv.ParseFloat(rfFQ, 64)
	if err != nil {
		log.Fatal("Wrong RF Fequency")
	}
	rf.Basic.Fequency = rfFQfloat

	fmt.Print("電台座標(緯度, 例: 23.958585): ")
	var rflocatLATstr string
	fmt.Scanln(&rflocatLATstr)
	rfLAT, err := strconv.ParseFloat(rflocatLATstr, 64)
	if err != nil {
		log.Fatal("Wrong RF Location")
	}

	fmt.Print("電台座標(經度, 例: 120.915759): ")
	var rflocatLNGstr string
	fmt.Scanln(&rflocatLNGstr)
	rfLNG, err := strconv.ParseFloat(rflocatLNGstr, 64)
	if err != nil {
		log.Fatal("Wrong RF Location")
	}

	// for test: 23.958585, 120.915759
	rf.Basic.RFLocation = [2]float64{rfLAT, rfLNG}

	rf.GetAllData()

	// helper.RFDistanceCalculator(xhaat, xfield)
	writer.Height3to15ExcelWriter(&rf)
	writer.Height10to50ExcelWriter(&rf)

	fmt.Printf("\n高度估算表已經順利產生。\n\n")

	return
}

func computeFMDistance() {
	fmt.Print("有效天線高度: ")
	var haatstr string
	fmt.Scanln(&haatstr)
	haat, err := strconv.ParseFloat(haatstr, 64)
	if err != nil {
		log.Fatal("有效天線高度錯誤")
	}

	fmt.Print("預估54場強: ")
	var f54str string
	fmt.Scanln(&f54str)
	f54, err := strconv.ParseFloat(f54str, 64)
	if err != nil {
		log.Fatal("預估54場強錯誤")
	}

	fmt.Printf("\n----------------------------------------\n\n計算結果：\n\n----------------------------------------\n\n")

	f60 := f54 + 6.0
	f66 := f60 + 6.0
	f80 := f60 + 20.0
	f88 := f80 + 8.0

	d54, comment54 := helper.RFDistanceCalculator(haat, f54)
	fmt.Printf("預估54場強傳送距離(FM、AM皆適用): %.2f (km)\n", d54)
	fmt.Print(comment54)
	fmt.Println()

	d60, comment60 := helper.RFDistanceCalculator(haat, f60)
	fmt.Printf("預估60場強傳送距離(FM適用): %.2f (km)\n", d60)
	fmt.Print(comment60)
	fmt.Println()

	d80, comment80 := helper.RFDistanceCalculator(haat, f80)
	fmt.Printf("預估80場強傳送距離(FM適用): %.2f (km)\n", d80)
	fmt.Print(comment80)
	fmt.Println()

	fmt.Printf("----------------------------------------\n\n")

	d66, comment66 := helper.RFDistanceCalculator(haat, f66)
	fmt.Printf("預估66場強傳送距離(AM適用): %.2f (km)\n", d66)
	fmt.Print(comment66)
	fmt.Println()

	d88, comment88 := helper.RFDistanceCalculator(haat, f88)
	fmt.Printf("預估88場強傳送距離(AM適用): %.2f (km)\n", d88)
	fmt.Print(comment88)
	fmt.Println()

	fmt.Printf("----------------------------------------\n")

}

// Get elevations by api from JawgMaps
// https://www.jawg.io/docs/apidocs/elevation/#examples
func getEvalationByPath() {
	res, err := http.Get("https://blog.syhlion.tw/sitemap.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", sitemap)
}
