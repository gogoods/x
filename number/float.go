package number

import (
	"strconv"
	"strings"
)

func FloatToStr(f float64) string {
	val := strconv.FormatFloat(f, 'f', 5, 64)
	if strings.Contains(val, ".") {
		val = strings.TrimRight(val, "0")
		val = strings.TrimRight(val, ".")
	}

	return val
}

func StrToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	} else {
		return f
	}
}
