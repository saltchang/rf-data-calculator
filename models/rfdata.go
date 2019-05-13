package models

import "fmt"

// RFDATA struct
type RFDATA struct {
	Basic struct {
		// 電台名稱 [B3]
		Name string

		// 射頻類型 (FM or AM)
		RFTYPE string

		// 電台頻率 (FM: MHz, AM: KHz) [B4]
		Fequency float64

		// 發射站座標 (緯度, 經度)
		RFLocation [2]float64
	}

	Height struct {
		// 地形取樣點數量
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

		// 地形起伏度 (dm) [B18:I18]
		DeltaH [8]float64

		// 地形起伏校正因數 (dB) [B19:I19]
		DeltaF [8]float64
	}

	Antenna struct {
		// 天線增益 (KW, dBd) [B6, D6]
		Ga [2]float64

		// 天線輻射中心海拔高度 (m) [B11]
		Hr int

		// 有效天線高度 (m) [B13:I13]
		HAAT [8]float64

		// 天線相對場型比 [B14:I14]
		AFR [8]float64

		// 天線相對場型比 (dB) [B15:I15]
		AFRdB [8]float64

		// 各方位天線增益 (dB) [B16:I16]
		GadB [8]float64
	}

	Lost struct {
		// 傳輸線損失 (dB) [B7]
		Lt1c float64

		// 接頭損失 (dB) [B8]
		Lt2t float64
	}

	Power struct {
		// 預估發射功率 (KW, dBK) [B5, D5]
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

func (rf *RFDATA) initData() {
	fmt.Println("Init the RF data...")
}
