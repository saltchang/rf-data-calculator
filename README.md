# RF Field Distance Calculator

This is a calculator application of engineering for calculating the RF field distance based on 1 ERP(kW), F(50, 50) Service Contour, FM Radio Channels 2 - 6, migrate from Javascript to Go.

It returns the distance, by given the field strength (in dBu) and the HAAT (meters) as input arguments.

I got the javascript source code from the official website of [FCC](https://www.fcc.gov/media/radio/fm-and-tv-propagation-curves).

這是一個工程用途的電場距離計算程式，建立在 1kW 之有效輻射功率、F(50, 50,)官方公開圖表、以及調頻廣播 2 - 6 頻道。

我特別將本應用程式由 Javascript 移植至 Go，JS 原始碼取自於 [FCC](https://www.fcc.gov/media/radio/fm-and-tv-propagation-curves) 的官方網站。

## Requirements

- [Go](https://golang.org/) installed in your computer.

OR

- If you don't have Go and you don't want to install it, just run the excutable file in `exec/{yourOS}`.

## Usage

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
