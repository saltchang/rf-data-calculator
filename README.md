# RF Field Distance Calculator

Building more larger structure and more more functions now...
May not use the feture in readme for now, sorry.

目前正在建立更大的結構以及更多更完整的功能，
在 readme 當中所說明的功能目前無法使用，十分抱歉..

## Introduction 基本介紹

This is an engineering calculator application for calculating the RF data of broadcast radio stations.
Including rf power data, terrain height data, antenna data, and field strength data, etc.
It stores data into xlsx format as a file.

這是一個廣播電台發射站之射頻資料的工程計算程式，計算包括發射功率、天線、地形、場強在內等等射頻資料。並將資料儲存為 xlsx 格式的檔案。

Remark:
It calculates the rf field distance based on 1 ERP(kW), F(50, 50) Service Contour, FM Radio Channels 2 - 6.

I got the field distance calculating source code from the official website of [FCC](https://www.fcc.gov/media/radio/fm-and-tv-propagation-curves), and migrated it from Javascript to Go in this project.

備註：
程式中計算電場距離之公式建立在 1kW 之有效輻射功率、F(50, 50,)官方公開圖表、以及調頻廣播 2 - 6 頻道。其原始碼取自於 [FCC](https://www.fcc.gov/media/radio/fm-and-tv-propagation-curves) 的官方網站。
在本專案當中，我將其由 Javascript 移植至 Go。

## Requirements 執行需求

- For development, [Go](https://golang.org/) needs to install in your computer.

- For normal using, just run the excutable file in `exec/{yourOS}`.

## Usage 使用方法

Clone the source code and the exec files from here:

HTTPS:

```shell

$ git clone https://github.com/saltchang/rf-field-distance-calculator.git

```

or SSH:

```shell

$ git clone git@github.com:saltchang/rf-field-distance-calculator.git

```

Enter the folder, build the excutable file:

```shell

$ cd rf-field-distance-calculator/

$ go build

```

After build, you will now have a excutable file in the current folder , just run it this way in the terminal:

```shell

$ ./rf-field-distance-calculator [HAAT] [Field strength]

> [Result: distance]

```

example:

```shell

$ ./rf-field-distance-calculator 30 60

> Distance: 10.161 km
> Done

```

If you don't have Go in your computer,
just run the excutable file in the `exec/` which matches your OS.

Take macOS as an example:

```shell

$ cd exec/macOS/

$ ./rf-field-distance-calculator 30 60

> Distance: 10.161 km
> Done

```

Thank you.
