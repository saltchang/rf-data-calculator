package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/saltchang/rf-field-distance-calculator/dclr"
	"github.com/saltchang/rf-field-distance-calculator/dcthr"
	"github.com/saltchang/rf-field-distance-calculator/exwr"
)

var (
	xhaat  = 30.0
	xfield = 60.0
)

func main() {
	dclr.GetFMDistance(xhaat, xfield)
	data3to15, data10to50 := dcthr.GetHeightData(23.958585, 120.915759)
	exwr.Height3to15ExcelWriter(data3to15)
	exwr.Height10to50ExcelWriter(data10to50)
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
