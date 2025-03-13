package leetcode

func IntToRoman(num int) string {
	romanStr := ""
	for num > 0 {
		switch {
		case num >= 1000:
			romanStr += "M"
			num -= 1000
		case num >= 900:
			romanStr += "CM"
			num -= 900
		case num >= 500:
			romanStr += "D"
			num -= 500
		case num >= 400:
			romanStr += "CD"
			num -= 400
		case num >= 100:
			romanStr += "C"
			num -= 100
		case num >= 90:
			romanStr += "XC"
			num -= 90
		case num >= 50:
			romanStr += "L"
			num -= 50
		case num >= 40:
			romanStr += "XL"
			num -= 40
		case num >= 10:
			romanStr += "X"
			num -= 10
		case num >= 9:
			romanStr += "IX"
			num -= 9
		case num >= 5:
			romanStr += "V"
			num -= 5
		case num >= 4:
			romanStr += "IV"
			num -= 4
		case num >= 1:
			romanStr += "I"
			num -= 1

		}

	}
	return romanStr

}
