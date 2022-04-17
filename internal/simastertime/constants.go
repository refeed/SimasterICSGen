package simastertime

import "time"

var (
	IDN_MONTH_NUM_MAP = map[string]time.Month{
		"Januari":   1,
		"Februari":  2,
		"Maret":     3,
		"April":     4,
		"Mei":       5,
		"Juni":      6,
		"Juli":      7,
		"Agustus":   8,
		"September": 9,
		"Oktober":   10,
		"November":  11,
		"Desember":  12,
	}

	// Simaster seemingly always uses UTC+7
	TZ = time.FixedZone("UTC+7", 7*60*60)

	IDN_DAY_NUM_MAP = map[string]time.Weekday{
		"Minggu": 0,
		"Senin":  1,
		"Selasa": 2,
		"Rabu":   3,
		"Kamis":  4,
		"Jumat":  5,
		"Sabtu":  6,
	}
)
