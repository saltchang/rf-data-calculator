package models

import (
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/saltchang/rf-data-calculator/helper"
)

// RFDATA struct
type RFDATA struct {

	// 基本資料 OK
	Basic struct {
		// 電台名稱 [B3]
		// 使用者輸入
		Name string

		// 射頻類型 (FM or AM)
		// 使用者輸入
		RFTYPE string

		// 電台頻率 (Hz) [B4]
		// 使用者輸入
		Fequency float64

		// 頻率單位 (FM: MHz, AM: KHz)
		FQUnit string

		// 發射站座標 (緯度, 經度)
		// 使用者輸入
		RFLocation [2]float64
	}

	// 高度資料 OK
	Height struct {
		// 地形取樣點數量
		// 常數：61, 41
		N3to15  int
		N10to50 int

		// 所有取樣點高度資料 (m)
		H3to15  [8][]float64
		H10to50 [8][]float64

		// 平均地形海拔高度 (m) [B12:I12]
		Hav3to15  [8]float64
		Hav10to50 [8]float64

		// 總平均地形海拔高度 (m)
		Htoav3to15  float64
		Htoav10to50 float64

		// 地形起伏度 (deltam) [B18:I18]
		DeltaH [8]float64

		// 地形起伏校正因數 (dB) [B19:I19]
		DeltaF [8]float64
	}

	Antenna struct {
		// 天線增益 (KW, dBd) [B6, D6]
		GaKW  float64
		GadBd float64 // 使用者輸入, 預設: 3.22dBd

		// 天線輻射中心海拔高度 (m) [B11]
		// 使用者輸入
		Hr int

		// 有效天線高度 (m) [B13:I13]
		HAAT [8]float64

		// 天線相對場型比 [B14:I14]
		// 使用者輸入
		AFR [8]float64

		// 天線相對場型比 (dB) [B15:I15]
		AFRdB [8]float64

		// 各方位天線增益 (dB) [B16:I16]
		GadB [8]float64
	}

	Lost struct {
		// 傳輸線損失 (dB) [B7]
		// 使用者輸入, 預設: -0.07dB
		Lt1c float64

		// 接頭損失 (dB) [B8]
		// 使用者輸入, 預設: -0.2dB
		Lt2t float64
	}

	Power struct {
		// 預估發射功率 (KW, dBK) [B5, D5]
		// 使用者輸入
		P0 [2]float64

		// 發射機輸出端有效輻射功率 (dBK) [B9]
		Pt float64

		// 各方位有效輻射功率 (dBK) [B17:I17]
		ERP [8]float64
	}

	FieldStrength struct {
		// 54場強
		FS54 struct {

			// 預估54場強 (dB) [B20:I20]
			F [8]float64

			// 預估54場強傳送距離 (km) [B21:I21]
			D [8]float64
		}

		// 60場強
		FS60 struct {

			// 預估60場強 (dB) [B25:I25]
			F [8]float64

			// 預估60場強傳送距離 (km) [B26:I26]
			D [8]float64
		}

		// 80場強
		FS80 struct {

			// 預估80場強 (dB) [B30:I30]
			F [8]float64

			// 預估80場強傳送距離 (km) [B31:I31]
			D [8]float64
		}
	}
}

// GetAllData function
func (rf *RFDATA) GetAllData() {
	fmt.Println("\nInit the RF data...")

	rf.Height.N3to15 = 61
	rf.Height.N10to50 = 41

	h3to15, h10to50 := helper.GetHeightData(rf.Basic.RFLocation)

	rf.Height.H3to15, rf.Height.H10to50 = h3to15, h10to50

	// 各項高度數據計算
	tosum3to15 := 0.0
	tosum10to50 := 0.0

	for i := 0; i < 8; i++ {
		av := 0.0

		sum := 0.0

		// 計算高度非零(海上)的地點
		noneZeroData := 0

		r := len(rf.Height.H3to15[i])
		for j := 0; j < r; j++ {
			sum += rf.Height.H3to15[i][j]
			if rf.Height.H3to15[i][j] != 0 {
				noneZeroData++
			}
		}
		if noneZeroData > 0 {
			av = math.Round(sum/float64(noneZeroData)*100.0) / 100.0
		} else {
			av = 0
		}
		rf.Height.Hav3to15[i] = av
		tosum3to15 += av

		sum = 0.0
		noneZeroData = 0

		r = len(rf.Height.H10to50[i])
		for j := 0; j < r; j++ {
			sum += rf.Height.H10to50[i][j]
			if rf.Height.H10to50[i][j] != 0 {
				noneZeroData++
			}
		}
		if noneZeroData > 0 {
			av = math.Round(sum/float64(noneZeroData)*100.0) / 100.0
		} else {
			av = 0
		}
		rf.Height.Hav10to50[i] = av
		tosum10to50 += av

		sort.Float64s(h10to50[i])
		indexRangeMin := [5]int{0, 1, 2, 3, 4}
		indexRangeMax := [5]int{36, 37, 38, 39, 40}
		maxsum := 0.0
		minsum := 0.0
		for j := 0; j < 5; j++ {
			maxsum += h10to50[i][indexRangeMax[j]]
			minsum += h10to50[i][indexRangeMin[j]]
		}

		deltaH := maxsum/5.0 - minsum/5.0
		if deltaH < 400 {
			rf.Height.DeltaH[i] = deltaH
		} else {
			rf.Height.DeltaH[i] = 400.0
		}
		rf.Height.DeltaF[i] = math.Round((1.9-0.03*rf.Height.DeltaH[i]*(1+rf.Basic.Fequency/300.0))*100.0) / 100.0
	}
	rf.Height.Htoav3to15 = math.Round(tosum3to15/8.0*100.0) / 100.0
	rf.Height.Htoav10to50 = math.Round(tosum10to50/8.0*100.0) / 100.0

	if rf.Basic.RFTYPE == "FM" {
		rf.Basic.FQUnit = "MHz"
	} else if rf.Basic.RFTYPE == "AM" {
		rf.Basic.FQUnit = "KHz"
	} else {
		fmt.Println("Wrong RF type! (must be 'FM' or 'AM')..")
		os.Exit(1)
	}
}
