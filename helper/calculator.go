package helper

import (
	"fmt"
	"math"
)

var (
	// Constant

	// D50 = Distance values and Height for F(50,50) curves
	D50 = []float64{1.609344, 3.218688, 4.828032, 6.437376, 8.046720, 16.09344, 32.18688, 48.28032, 64.37376, 80.46720, 96.56064, 112.65408, 128.74752, 144.84096, 160.93440, 177.02784, 193.12128, 209.21472, 225.30816, 241.40160, 257.49504, 273.58848, 289.68192, 305.77536, 321.86880, 0.0}

	// H50 = Distance values and Height for F(50,50) curves
	H50 = []float64{30.48, 60.96, 121.92, 182.88, 243.84, 304.80, 381.00, 457.20, 533.40, 609.60, 914.40, 1219.20, 1524.00, 0.0}

	// F55LV = Distance values and Height for F(50,50) curves
	F55LV = []float64{92., 79.7, 72.7, 67.8, 64., 52., 39.4, 31., 25.3, 20.3, 16.2, 12.8, 9.8, 6.9, 4., 1.5, -1.1, -3.6, -5.8, -8.1, -10.6, -13., -15.1, -17.2, -19.2, 98., 85.9, 79., 73.8, 70., 58., 45.5, 37., 29.5, 23.5, 18.1, 14.5, 11., 8.2, 5.5, 2.9, .3, -2.2, -4.8, -7., -9.4, -11.7, -14., -16.1, -18.3, 100.6, 91., 84.8, 80., 76., 64., 51.5, 43., 35.5, 28.8, 22., 17.1, 13.4, 10.2, 7.4, 4.8, 2.2, -.3, -3., -5.2, -7.6, -10., -12.2, -14.6, -16.9, 101.5, 93.4, 87.8, 83.3, 79.6, 67.6, 55., 46.7, 39., 32., 25.3, 19.8, 15.2, 11.8, 8.9, 6., 3.7, 1., -1.4, -3.9, -6.1, -8.7, -11., -13.2, -15.6, 101.9, 94.6, 89.4, 85.4, 82., 70., 57.6, 49., 41.5, 34.4, 27.7, 22., 17., 13.1, 10.1, 7.2, 4.8, 2., -.3, -2.7, -5.1, -7.6, -10., -12.1, -14.6, 102., 95., 90.4, 86.8, 83.7, 72., 59.6, 51., 43.6, 36.7, 29.9, 23.9, 18.8, 14.7, 11.5, 8.4, 5.7, 3., .6, -1.8, -4.2, -6.6, -9., -11.2, -13.6, 102.1, 95.6, 91.2, 87.7, 85., 73.9, 61.7, 53.2, 45.9, 39.1, 32., 26., 21., 16.8, 13.1, 9.9, 7., 4.1, 1.7, -.7, -3.2, -5.6, -8., -10.2, -12.5, 102.2, 95.9, 91.8, 88.3, 85.8, 75.4, 63.3, 55.1, 47.9, 41.5, 34.4, 28.3, 23.2, 18.8, 14.9, 11.1, 8., 5.2, 2.7, .2, -2.2, -4.6, -7., -9.2, -11.6, 102.3, 96., 92., 88.9, 86.3, 76.7, 64.9, 57., 50., 43.5, 36.7, 30.7, 25.2, 20.4, 16., 12.5, 9.1, 6.2, 3.8, 1.1, -1.3, -3.6, -6.1, -8.4, -10.6, 102.4, 96.1, 92.2, 89.2, 86.7, 77.9, 66.2, 58.5, 51.5, 45., 38.2, 32.4, 27., 22., 17.3, 13.7, 10.1, 7.1, 4.6, 2., -.4, -2.7, -5.1, -7.6, -10., 102.5, 96.3, 92.5, 89.9, 87.6, 80.2, 70., 62.6, 55.4, 48.9, 42.5, 36.9, 31., 25.7, 21., 17.1, 13.6, 10.3, 7.8, 5.1, 2.8, .5, -2.1, -4.5, -6.8, 102.5, 96.5, 92.5, 90.1, 88., 81.3, 72.4, 65., 57.8, 51.2, 44.9, 39.1, 33.2, 28.1, 23.5, 19.8, 16.1, 13., 10.4, 8., 5.5, 3.1, .6, -2., -4.1, 102.5, 96.5, 92.5, 90.2, 88.1, 81.9, 74.2, 66.5, 59.6, 53., 46.4, 40.8, 35., 30., 25.5, 21.8, 18.3, 15., 12.4, 10., 7.7, 5.1, 2.8, .2, -2.}

	xerp      = 1.0
	xdistance = 0.0

	flag = false

	nPoints = 1001
	d       = make([]float64, 1002, 1002)
	f       = make([]float64, 1002, 1002)
)

// RFDistanceCalculator func
func RFDistanceCalculator(xhaat, xfield float64) (float64, string) {

	result := metric(xerp, xhaat, xfield, xdistance)
	result = math.Round(result*1000) / 1000
	comment := comment()

	fmt.Println("Distance:", result, "km")
	fmt.Println(comment)

	return result, comment
}

func printUsage() {
	fmt.Println("Useage: ./rf-field-distance-calculator [HAAT] [Field strength]")
}

// func()
func metric(erp, haat, field, distance float64) float64 {
	id50 := 25.0
	ih50 := 13.0
	delta := 0.5
	erpDB := 0.0
	dFirst := 1.5
	// dLast := 300.0
	eVoltsMeter := 0.0

	nPoints = 1001

	d = make([]float64, 1002, 1002)
	f = make([]float64, 1002, 1002)
	h := make([]float64, 1002, 1002)
	f5050 := make([]float64, 1002, 1002)

	erpDB = 10.0 * (math.Log(erp) / math.Log(10))

	if haat < 30.0 {
		// All HAAT below 30 meters are set to 30
		haat = 30.0
	} else if haat > 1600.0 {
		// All HAAT above 1600 are set to 1600
		haat = 1600.0
	}

	for i := 0; i <= nPoints; i++ {
		h[i] = haat
		f[i] = 0.0
		d[i] = 0.0
		f5050[i] = 0.0
	}

	// dFirst/delta must be an integer
	k := int(math.Floor(dFirst / delta))

	for i := k; i <= nPoints; i++ {
		d[i] = (float64(i) * delta)
	}

	itplbv(id50, ih50, D50, H50, F55LV, h)

	// i = 2,1,0 (3 points) 1.5, 1, 0.5
	for i := 3; i > 0; i-- {
		if field > f[i] {
			// Use the free space equation to find the field strength and distance
			flag = true
			eVoltsMeter = 1.0e-6 * math.Pow(10, (field/20.))
			distance = (7.014271e-3 * math.Sqrt(erp*1000.)) / eVoltsMeter
			return distance
		}
	}

	// points i=0,1,2 covered by free space equation immediately above
	for i := 0; i < nPoints; i++ {
		f[i] = f[i] + erpDB
	}
	// Most common, for service and interfering contours
	// i start at 1
	for i := 1; i < nPoints; i++ {
		if field > f[i] && field < f[i-1] {
			distance = (((f[i-1] - field) / (f[i-1] - f[i])) * (d[i] - d[i-1])) + d[i-1]
			return distance
		}
	} // should not get here!
	return distance
} // end metric ----------------------------------------------------------------------------------
// itplbv
func itplbv(lx, ly float64, x, y, z []float64, v []float64) {

	// declarations and initializations
	var (
		lxm1 = 0.0
		// lxp1 = 0.0
		lym1 = 0.0
		// lyp1 = 0.0
		ixpv = 0.0
		iypv = 0.0
		// k    = 0.0
		ix  = 0.0
		iy  = 0.0
		imn = 0.0
		imx = 0.0
		jx  = 0.0
		jy  = 0.0
		jx1 = 0.0
		jy1 = 0.0

		zaRow0 = make([]float64, 2, 2)
		zaRow1 = make([]float64, 2, 2)
		zaRow2 = make([]float64, 2, 2)
		zaRow3 = make([]float64, 2, 2)
		zaRow4 = make([]float64, 2, 2)
		za     = [][]float64{zaRow0, zaRow1, zaRow2, zaRow3, zaRow4} // za[5,2]

		zbRow0 = make([]float64, 5, 5)
		zbRow1 = make([]float64, 5, 5)
		zb     = [][]float64{zbRow0, zbRow1} // zb[2,5]

		zabRow0 = make([]float64, 3, 3)
		zabRow1 = make([]float64, 3, 3)
		zabRow2 = make([]float64, 3, 3)

		zab = [][]float64{zabRow0, zabRow1, zabRow2} // zab[3,3]

		zxRow0 = make([]float64, 4, 4)
		zxRow1 = make([]float64, 4, 4)
		zxRow2 = make([]float64, 4, 4)
		zxRow3 = make([]float64, 4, 4)
		zx     = [][]float64{zxRow0, zxRow1, zxRow2, zxRow3} // zx[4,4]

		zyRow0 = make([]float64, 4, 4)
		zyRow1 = make([]float64, 4, 4)
		zyRow2 = make([]float64, 4, 4)
		zyRow3 = make([]float64, 4, 4)
		zy     = [][]float64{zyRow0, zyRow1, zyRow2, zyRow3} // zy[4,4]

		zxyRow0 = make([]float64, 4, 4)
		zxyRow1 = make([]float64, 4, 4)
		zxyRow2 = make([]float64, 4, 4)
		zxyRow3 = make([]float64, 4, 4)
		zxy     = [][]float64{zxyRow0, zxyRow1, zxyRow2, zxyRow3} // zy[4,4]

		x3  = 0.0
		x4  = 0.0
		a3  = 0.0
		y3  = 0.0
		y4  = 0.0
		b3  = 0.0
		z33 = 0.0
		z43 = 0.0
		z34 = 0.0
		z44 = 0.0
		x2  = 0.0
		a2  = 0.0
		z23 = 0.0
		z24 = 0.0
		x5  = 0.0
		a4  = 0.0
		z53 = 0.0
		z54 = 0.0

		a1, a5, y2, b2, z32, z42, y5, b4, z35, z45, b1, b5, w2, w3, sw, wx2, wx3, wy2, wy3, w1, w4, w5 float64

		zx3b3, zx4b3, zy3a3, zy4a3, a, b, c, ddd, e, a3sq, b3sq, p02, p03, p12, p13, p20, p21, p22 float64

		p23, p30, p31, p32, p33, dy, q0, q1, q2, q3, dx float64
	)

	// Calculations begin
	lx = math.Floor(lx)
	ly = math.Floor(ly)

	lxm1 = math.Floor(lx - 1)
	// lxp1 = math.Floor(lx + 1)
	lym1 = math.Floor(ly - 1)
	// lyp1 = math.Floor(ly + 1)
	ixpv = -1
	iypv = -1

	for k := 0; k < nPoints; k++ {
		if d[k] >= x[int(lxm1)] {
			ix = lx
		} else {
			if d[k] < x[0] {
				ix = 0
			} else {
				imn = 1
				imx = lxm1
				for {
					ix = math.Floor((imn + imx) / 2)
					if d[k] >= x[int(ix)] {
						imn = ix + 1
					} else {
						imx = ix
					}
					if imx <= imn {
						break
					}
				}
				ix = imx
			}
		}
		ix = math.Floor(ix)

		if v[k] >= y[int(lym1)] {
			iy = ly
		} else {
			if v[k] < y[0] {
				iy = 0
			} else {
				imn = 1
				imx = lym1
				for {
					iy = math.Floor((imn + imx) / 2)
					if v[k] >= y[int(iy)] {
						imn = iy + 1
					} else {
						imx = iy
					}
					if imx <= imn {
						break
					}
				}
				iy = imx
			}
		}
		iy = math.Floor(iy)

		if ix != ixpv || iy != iypv {
			ixpv = ix
			iypv = iy
			if ix == 0 {
				jx = 1
			} else {
				if ix == lx {
					jx = lxm1
				} else {
					jx = ix
				}
			}
			if iy == 0 {
				jy = 1
			} else {
				if iy == ly {
					jy = lym1
				} else {
					jy = iy
				}
			}
			jx = math.Floor(jx)
			jy = math.Floor(jy)

			x3 = x[int(jx)-1]
			x4 = x[int(jx)]
			a3 = 1.0 / (x4 - x3)
			y3 = y[int(jy)-1]
			y4 = y[int(jy)]
			b3 = 1.0 / (y4 - y3)
			z33 = z[(int(jx)-1)+((int(jy)-1)*int(lx))]
			z43 = z[int(jx)+((int(jy)-1)*int(lx))]
			z34 = z[(int(jx)-1)+int(jy*lx)]
			z44 = z[int(jx)+int(jy*lx)]
			za[2][0] = (z43 - z33) * a3
			za[2][1] = (z44 - z34) * a3
			zb[0][2] = (z34 - z33) * b3
			zb[1][2] = (z44 - z43) * b3
			zab[1][1] = (zb[1][2] - zb[0][2]) * a3

			if jx > 1 {
				x2 = x[int(jx)-2]
				a2 = 1.0 / (x3 - x2)
				z23 = z[int(jx-2)+(int(jy-1)*int(lx))]
				z24 = z[int(jx-2)+int(jy*lx)]
				za[1][0] = (z33 - z23) * a2
				za[1][1] = (z34 - z24) * a2
				if jx == lxm1 {
					za[3][0] = (2.0 * za[2][0]) - za[1][0]
					za[3][1] = (2.0 * za[2][1]) - za[1][1]
				}
			}

			if jx < lxm1 {
				x5 = x[int(jx)+1]
				a4 = 1.0 / (x5 - x4)
				z53 = z[int(jx+1)+int((jy-1)*lx)]
				z54 = z[int(jx+1)+int(jy*lx)]
				za[3][0] = (z53 - z43) * a4
				za[3][1] = (z54 - z44) * a4
				if jx == 1 {
					za[1][0] = (2.0 * za[2][0]) - za[3][0]
					za[1][1] = (2.0 * za[2][1]) - za[3][1]
				}
			}

			zab[0][1] = (za[1][1] - za[1][0]) * b3
			zab[2][1] = (za[3][1] - za[3][0]) * b3
			if jx > 2 {
				a1 = 1.0 / (x2 - x[int(jx)-3])
				za[0][0] = (z23 - z[int(jx-3)+int((jy-1)*lx)]) * a1
				za[0][1] = (z24 - z[int(jx-3)+int(jy*lx)]) * a1
			} else {
				za[0][0] = (2.0 * za[1][0]) - za[2][0]
				za[0][1] = (2.0 * za[1][1]) - za[2][1]
			}
			if jx < (lx - 2) {
				a5 = 1.0 / (x[int(jx)+2] - x5)
				za[4][0] = (z[int(jx+2)+int((jy-1)*lx)] - z53) * a5
				za[4][1] = (z[int(jx+2)+int(jy*lx)] - z54) * a5
			} else {
				za[4][0] = (2.0 * za[3][0]) - za[2][0]
				za[4][1] = (2.0 * za[3][1]) - za[2][1]
			}
			if jy > 1 {
				y2 = y[int(jy)-2]
				b2 = 1.0 / (y3 - y2)
				z32 = z[int(jx-1)+int((jy-2)*lx)]
				z42 = z[int(jx)+int((jy-2)*lx)]
				zb[0][1] = (z33 - z32) * b2
				zb[1][1] = (z43 - z42) * b2
				if jy == lym1 {
					zb[0][3] = (2.0 * zb[0][2]) - zb[0][1]
					zb[1][3] = (2.0 * zb[1][2]) - zb[1][1]
				}
			}
			if jy < lym1 {
				y5 = y[int(jy)+1]
				b4 = 1.0 / (y5 - y4)
				z35 = z[int(jx-1)+int((jy+1)*lx)]
				z45 = z[int(jx)+int((jy+1)*lx)]
				zb[0][3] = (z35 - z34) * b4
				zb[1][3] = (z45 - z44) * b4
				if jy == 1 {
					zb[0][1] = (2.0 * zb[0][2]) - zb[0][3]
					zb[1][1] = (2.0 * zb[1][2]) - zb[1][3]
				}
			}
			zab[1][0] = (zb[1][1] - zb[0][1]) * a3
			zab[1][2] = (zb[1][3] - zb[0][3]) * a3
			if jy > 2 {
				b1 = 1.0 / (y2 - y[int(jy)-3])
				zb[0][0] = (z32 - z[int(jx-1)+int((jy-3)*lx)]) * b1
				zb[1][0] = (z42 - z[int(jx)+int((jy-3)*lx)]) * b1
			} else {
				zb[0][0] = (2.0 * zb[0][1]) - zb[0][2]
				zb[1][0] = (2.0 * zb[1][1]) - zb[1][2]
			}
			if jy < (ly - 2) {
				b5 = 1.0 / (y[int(jy)+2] - y5)
				zb[0][4] = (z[int(jx-1)+int((jy+2)*lx)] - z35) * b5
				zb[1][4] = (z[int(jx)+int((jy+2)*lx)] - z45) * b5
			} else {
				zb[0][4] = (2.0 * zb[0][3]) - zb[0][2]
				zb[1][4] = (2.0 * zb[1][3]) - zb[1][2]
			}
			if jx < lxm1 {
				if jy > 1 {
					zab[2][0] = ((z53-z[int(jx+1)+int((jy-2)*lx)])*b2 - zb[1][1]) * a4
					if jy < lym1 {
						zab[2][2] = ((z[int(jx+1)+int((jy+1)*lx)]-z54)*b4 - zb[1][3]) * a4
					} else {
						zab[2][2] = (2.0 * zab[2][1]) - zab[2][0]
					}
				} else {
					zab[2][2] = ((z[int(jx+1)+int((jy+1)*lx)]-z54)*b4 - zb[1][3]) * a4
					zab[2][0] = (2.0 * zab[2][1]) - zab[2][2]
				}
				if jx == 1 {
					zab[0][0] = (2.0 * zab[1][0]) - zab[2][0]
					zab[0][2] = (2.0 * zab[1][2]) - zab[2][2]
				}
			}
			if jx > 1 {
				if jy > 1 {
					zab[0][0] = (zb[0][1] - (z23-z[int(jx-2)+int((jy-2)*lx)])*b2) * a2
					if jy < lym1 {
						zab[0][2] = (zb[0][3] - (z[int(jx-2)+int((jy+1)*lx)]-z24)*b4) * a2
					} else {
						zab[0][2] = (2.0 * zab[0][1]) - zab[0][0]
					}
				} else {
					zab[0][2] = (zb[0][3] - (z[int(jx-2)+int((jy+1)*lx)]-z24)*b4) * a2
					zab[0][0] = (2.0 * zab[0][1]) - zab[0][2]
				}
				if jx == lxm1 {
					zab[2][0] = (2.0 * zab[1][0]) - zab[0][0]
					zab[2][2] = (2.0 * zab[1][2]) - zab[0][2]
				}
			}
			for jy = 1; jy < 3; jy++ {
				for jx = 1; jx < 3; jx++ {
					w2 = math.Abs(za[int(jx)+2][int(jy)-1] - za[int(jx)+1][int(jy)-1])
					w3 = math.Abs(za[int(jx)][int(jy)-1] - za[int(jx)-1][int(jy)-1])
					sw = w2 + w3
					if sw >= 1.e-7 {
						wx2 = w2 / sw
						wx3 = w3 / sw
					} else {
						wx2 = 0.5
						wx3 = 0.5
					}
					zx[int(jx)][int(jy)] = wx2*za[int(jx)][int(jy)-1] + wx3*za[int(jx)+1][int(jy)-1]
					w2 = math.Abs(zb[int(jx)-1][int(jy)+2] - zb[int(jx)-1][int(jy)+1])
					w3 = math.Abs(zb[int(jx)-1][int(jy)] - zb[int(jx)-1][int(jy)-1])
					sw = w2 + w3
					if sw >= 1.e-7 {
						wy2 = w2 / sw
						wy3 = w3 / sw
					} else {
						wy2 = 0.5
						wy3 = 0.5
					}
					zy[int(jx)][int(jy)] = wy2*zb[int(jx)-1][int(jy)] + wy3*zb[int(jx)-1][int(jy)+1]
					zxy[int(jx)][int(jy)] = wy2*(wx2*zab[int(jx)-1][int(jy)-1]+wx3*zab[int(jx)][int(jy)-1]) + wy3*(wx2*zab[int(jx)-1][int(jy)]+wx3*zab[int(jx)][int(jy)])
				}
			}
			if ix == 0 {
				w2 = a4 * (3.0*a3 + a4)
				w1 = 2.0*a3*(a3-a4) + w2
				for jy = 1; jy < 3; jy++ {
					zx[0][int(jy)] = (w1*za[0][int(jy)-1] + w2*za[1][int(jy)-1]) / (w1 + w2)
					zy[0][int(jy)] = (2.0 * zy[1][int(jy)]) - zy[2][int(jy)]
					zxy[0][int(jy)] = (2.0 * zxy[1][int(jy)]) - zxy[2][int(jy)]
					for jx1 = 1; jx1 < 3; jx1++ {
						jx = 3 - jx1
						zx[int(jx)][int(jy)] = zx[int(jx)-1][int(jy)]
						zy[int(jx)][int(jy)] = zy[int(jx)-1][int(jy)]
						zxy[int(jx)][int(jy)] = zxy[int(jx)-1][int(jy)]
					}
				}
				x3 -= 1.0 / a4
				z33 -= za[1][0] / a4
				for jy = 0; jy < 5; jy++ {
					zb[1][int(jy)] = zb[0][int(jy)]
				}
				for jy = 1; jy < 4; jy++ {
					zb[0][int(jy)] -= zab[0][int(jy)-1] / a4
				}
				a3 = a4
				za[2][0] = za[1][0]
				for jy = 0; jy < 3; jy++ {
					zab[1][int(jy)] = zab[0][int(jy)]
				}
			}
			if ix == lx {
				w4 = a2 * (3.0*a3 + a2)
				w5 = 2.0*a3*(a3-a2) + w4
				for jy = 1; jy < 3; jy++ {
					zx[3][int(jy)] = (w4*za[3][int(jy)-1] + w5*za[4][int(jy)-1]) / (w4 + w5)
					zy[3][int(jy)] = (2.0 * zy[2][int(jy)]) - zy[1][int(jy)]
					zxy[3][int(jy)] = (2.0 * zxy[2][int(jy)]) - zxy[1][int(jy)]
					for jx = 1; jx < 3; jx++ {
						zx[int(jx)][int(jy)] = zx[int(jx)+1][int(jy)]
						zy[int(jx)][int(jy)] = zy[int(jx)+1][int(jy)]
						zxy[int(jx)][int(jy)] = zxy[int(jx)+1][int(jy)]
					}
				}
				x3 = x4
				z33 = z43
				for jy = 0; jy < 5; jy++ {
					zb[0][int(jy)] = zb[1][int(jy)]
				}
				a3 = a2
				za[2][0] = za[3][0]
				for jy = 0; jy < 3; jy++ {
					zab[1][int(jy)] = zab[2][int(jy)]
				}
			}
			if iy == 0 {
				w2 = b4 * (3.0*b3 + b4)
				w1 = 2.0*b3*(b3-b4) + w2
				for jx = 1; jx < 3; jx++ {
					if (ix > 0 || jx == 2) && (ix < lx || jx == 1) {
						zy[int(jx)][0] = (w1*zb[int(jx)-1][0] + w2*zb[int(jx)-1][1]) / (w1 + w2)
						zx[int(jx)][0] = (2.0 * zx[int(jx)][1]) - zx[int(jx)][2]
						zxy[int(jx)][0] = (2.0 * zxy[int(jx)][1]) - zxy[int(jx)][2]
					}
					for jy1 = 1; jy1 < 3; jy1++ {
						jy = 3 - jy1
						zy[int(jx)][int(jy)] = zy[int(jx)][int(jy)-1]
						zx[int(jx)][int(jy)] = zx[int(jx)][int(jy)-1]
						zxy[int(jx)][int(jy)] = zxy[int(jx)][int(jy)-1]
					}
				}
				y3 -= 1.0 / b4
				z33 -= zb[0][1] / b4
				za[2][0] -= zab[1][0] / b4
				zb[0][2] = zb[0][1]
				zab[1][1] = zab[1][0]
				b3 = b4
				if ix == 0 || ix == lx {
					if ix == 0 {
						jx = 1
						jx1 = 2
					} else {
						jx = 2
						jx1 = 1
					}
					jx1 = math.Floor(jx1)
					zx[int(jx)][1] = zx[int(jx1)][1] + zx[int(jx)][2] - zx[int(jx1)][2]
					zy[int(jx)][1] = zy[int(jx1)][1] + zy[int(jx)][2] - zy[int(jx1)][2]
					zxy[int(jx)][1] = zxy[int(jx1)][1] + zxy[int(jx)][2] - zxy[int(jx1)][2]
				}
			}
			if iy == ly {
				w4 = b2 * (3.0*b3 + b2)
				w5 = 2.0*b3*(b3-b2) + w4
				for jx = 1; jx < 3; jx++ {
					if (ix > 0 || jx == 2) && (ix < lx || jx == 1) {
						zy[int(jx)][3] = (w4*zb[int(jx)-1][3] + w5*zb[int(jx)-1][4]) / (w4 + w5)
						zx[int(jx)][3] = (2.0 * zx[int(jx)][2]) - zx[int(jx)][1]
						zxy[int(jx)][3] = (2.0 * zxy[int(jx)][2]) - zxy[int(jx)][1]
					}
					for jy = 1; jy < 3; jy++ {
						zy[int(jx)][int(jy)] = zy[int(jx)][int(jy)+1]
						zx[int(jx)][int(jy)] = zx[int(jx)][int(jy)+1]
						zxy[int(jx)][int(jy)] = zxy[int(jx)][int(jy)+1]
					}
				}
				y3 = y4
				z33 += zb[0][2] / b3
				za[2][0] += zab[1][1] / b3
				zb[0][2] = zb[0][3]
				zab[1][1] = zab[1][2]
				b3 = b2
				if ix == 0 || ix == lx {
					if ix == 0 {
						jx = 1
						jx1 = 2
					} else {
						jx = 2
						jx1 = 1
					}
					zx[int(jx)][2] = zx[int(jx1)][2] + zx[int(jx)][1] - zx[int(jx1)][1]
					zy[int(jx)][2] = zy[int(jx1)][2] + zy[int(jx)][1] - zy[int(jx1)][1]
					zxy[int(jx)][2] = zxy[int(jx1)][2] + zxy[int(jx)][1] - zxy[int(jx1)][1]
				}
			}
			zx3b3 = (zx[1][2] - zx[1][1]) * b3
			zx4b3 = (zx[2][2] - zx[2][1]) * b3
			zy3a3 = (zy[2][1] - zy[1][1]) * a3
			zy4a3 = (zy[2][2] - zy[1][2]) * a3
			a = zab[1][1] - zx3b3 - zy3a3 + zxy[1][1]
			b = zx4b3 - zx3b3 - zxy[2][1] + zxy[1][1]
			c = zy4a3 - zy3a3 - zxy[1][2] + zxy[1][1]
			ddd = zxy[2][2] - zxy[2][1] - zxy[1][2] + zxy[1][1]
			e = a + a - b - c
			a3sq = a3 * a3
			b3sq = b3 * b3
			p02 = (2.0*(zb[0][2]-zy[1][1]) + zb[0][2] - zy[1][2]) * b3
			p03 = (-2.0*zb[0][2] + zy[1][2] + zy[1][1]) * b3sq
			p12 = (2.0*(zx3b3-zxy[1][1]) + zx3b3 - zxy[1][2]) * b3
			p13 = (-2.0*zx3b3 + zxy[1][2] + zxy[1][1]) * b3sq
			p20 = (2.0*(za[2][0]-zx[1][1]) + za[2][0] - zx[2][1]) * a3
			p21 = (2.0*(zy3a3-zxy[1][1]) + zy3a3 - zxy[2][1]) * a3
			p22 = (3.0*(a+e) + ddd) * a3 * b3
			p23 = (-3.0*e - b - ddd) * a3 * b3sq
			p30 = (-2.0*za[2][0] + zx[2][1] + zx[1][1]) * a3sq
			p31 = (-2.0*zy3a3 + zxy[2][1] + zxy[1][1]) * a3sq
			p32 = (-3.0*e - c - ddd) * b3 * a3sq
			p33 = (ddd + e + e) * a3sq * b3sq
		}
		dy = v[k] - y3
		q0 = z33 + dy*(zy[1][1]+dy*(p02+dy*p03))
		q1 = zx[1][1] + dy*(zxy[1][1]+dy*(p12+dy*p13))
		q2 = p20 + dy*(p21+dy*(p22+dy*p23))
		q3 = p30 + dy*(p31+dy*(p32+dy*p33))
		dx = d[k] - x3
		f[k] = q0 + dx*(q1+dx*(q2+dx*q3))
	}
}

// end itplbv func

// comment
func comment() string {
	if flag {
		return "Done\n(Free Space equation used to compute distance)"
	}
	return "Done"
}
