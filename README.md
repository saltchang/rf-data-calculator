# RF Data Calculator 射頻資料計算程式

This is an Go app for calculating the RF data of broadcast radio stations.
這是一個計算廣播電台發射站之射頻資料的 Go 應用。

2019-05-14
v1.0 Released

## Feature 功能

1. 取得八方位地形資料並計算平均高度以及地形起伏度，並且將資料匯出為 .xlsx 檔案。
2. 計算預估電場強度傳送距離。

## Usage 使用方法

- For normal using, download the archive for your OS, extract it and run the `RF_dataApp` executable file.

- For development, clone it into `$GOPATH/src/github.com/yourname/` and do whatever you want.

To run the app in command line:

```shell

$ go run main.go

```

## Dependency 依賴套件

- [excelize](https://github.com/360EntSecGroup-Skylar/excelize) - Used to write the data as .xlsx files.

- [JawgMaps(API)](https://www.jawg.io/docs/apidocs/elevation/#examples) - Used to Request elevation data.

- [Go Dep](https://github.com/golang/dep) - Used to manage the dependencies.

For development, to install the dependencies by execute the following dep command:

```shell

$ dep ensure -v

```

## Remark 備註

It calculates the rf field distance based on 1 ERP(kW), F(50, 50) Service Contour, FM Radio Channels 2 - 6.
I got the field distance calculating source code from the official website of [FCC](https://www.fcc.gov/media/radio/fm-and-tv-propagation-curves), and migrated it from Javascript to Go in this project.

程式中計算電場距離之公式建立在 1kW 之有效輻射功率、F(50, 50,)官方公開圖表、以及調頻廣播 2 - 6 頻道。其原始碼取自於 [FCC](https://www.fcc.gov/media/radio/fm-and-tv-propagation-curves) 的官方網站。
在本專案當中，我將其由 Javascript 移植至 Go。
