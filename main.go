package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/saltchang/rf-data-calculator/helper"
)

var (
	xhaat  = 30.0
	xfield = 60.0
)

func main() {
	helper.RFDistanceCalculator(xhaat, xfield)
	data3to15, data10to50 := helper.GetHeightData(23.958585, 120.915759)
	helper.Height3to15ExcelWriter(data3to15)
	helper.Height10to50ExcelWriter(data10to50)
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
