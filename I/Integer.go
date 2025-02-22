package I

// Integer number support package

import (
	"strconv"
	"strings"
)

// simplified ternary operator (bool ? val : 0), returns second argument, if the condition (first arg) is true, returns 0 if not
//  I.If(true,3) // 3
//  I.If(false,3) // 0
func If(b bool, yes int64) int64 {
	if b {
		return yes
	}
	return 0
}

// ternary operator (bool ? val1 : val2), returns second argument if the condition (first arg) is true, third argument if not
//  I.IfElse(true,3,4) // 3
//  I.IfElse(false,3,4) // 4
func IfElse(b bool, yes, no int64) int64 {
	if b {
		return yes
	}
	return no
}

// simplified ternary operator (bool ? val1==0 : val2), returns second argument, if val1 (first arg) is zero, returns val2 if not
//  I.IfZero(0,3) // 3
//  I.IfZero(4,3) // 4
func IfZero(val1, val2 int64) int64 {
	if val1 == 0 {
		return val2
	}
	return val1
}

// simplified ternary operator (bool ? val1==0 : val2), returns second argument, if val1 (first arg) is zero, returns val2 if not
//  I.IsZero(0,3) // 3
//  I.IsZero(4,3) // 4
func IsZero(val1, val2 int) int {
	if val1 == 0 {
		return val2
	}
	return val1
}

// simplified ternary operator (bool ? val : 0), returns second argument, if the condition (first arg) is true, returns 0 if not
//  UI.UIf(true,3) // 3
//  UI.UIf(false,3) // 0
func UIf(b bool, yes uint64) uint64 {
	if b {
		return yes
	}
	return 0
}

// ternary operator (bool ? val1 : val2), returns second argument if the condition (first arg) is true, third argument if not
//  UI.UIfElse(true,3,4) // 3
//  UI.UIfElse(false,3,4) // 4
func UIfElse(b bool, yes, no uint64) uint64 {
	if b {
		return yes
	}
	return no
}

// simplified ternary operator (bool ? val1==0 : val2), returns second argument, if val1 (first arg) is zero, returns val2 if not
//  UI.UIfZero(0,3) // 3
//  UI.UIfZero(4,3) // 4
func UIfZero(val1, val2 uint64) uint64 {
	if val1 == 0 {
		return val2
	}
	return val1
}

// simplified ternary operator (bool ? val1==0 : val2), returns second argument, if val1 (first arg) is zero, returns val2 if not
//  UI.UIsZero(0,3) // 3
//  UI.UIsZero(4,3) // 4
func UIsZero(val1, val2 uint) uint {
	if val1 == 0 {
		return val2
	}
	return val1
}

// convert int64 to string
//  I.ToS(int64(1234)) // `1234`
func ToS(num int64) string {
	return strconv.FormatInt(num, 10)
}

// convert int to string
//  I.ToStr(1234) // `1234`
func ToStr(num int) string {
	return strconv.Itoa(num)
}

// convert uint64 to string
//  I.UToS(uint64(1234)) // `1234`
func UToS(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// convert int to string
//  I.UToStr(uint(1234)) // `1234`
func UToStr(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

// int64 min of two values
//  I.Min(int64(3),int64(4)) // 3
func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// int64 max of two values
//  I.Max(int64(3),int64(4)) // 4
func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// uint64 min of two values
//  I.UMin(uint64(3),uint64(4)) // 3
func UMin(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// uint64 max of two values
//  I.UMax(uint64(3),uint64(4)) // 4
func UMax(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// int min of two values
//  I.MinOf(3,4) // 3
func MinOf(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// int max of two values
//  I.MaxOf(3,4) // 4
func MaxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// uint min of two values
//  I.MinOf(3,4) // 3
func UMinOf(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// uint max of two values
//  I.MaxOf(3,4) // 4
func UMaxOf(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

// format ordinal number suffix such as st, nd, rd, and th.
//  I.ToEnglishNum(241)) // `241st`
//  I.ToEnglishNum(242)) // `242nd`
//  I.ToEnglishNum(244)) // `244th`
func ToEnglishNum(num int64) string {
	if num < 0 {
		return ``
	}
	prefix := ToS(num)
	n2 := num % 100
	num %= 10
	if n2 == 11 || n2 == 12 || n2 == 13 {
		prefix += `th`
	} else if num == 1 {
		prefix += `st`
	} else if num == 2 {
		prefix += `nd`
	} else if num == 3 {
		prefix += `rd`
	} else {
		prefix += `th`
	}
	return prefix
}

// converts int64 (first arg) to string with zero padded with maximum length
// I.PadZero(123,5) // `00123`
func PadZero(num int64, length int) string {
	str := ToS(num)
	slen := len(str)
	if slen >= length {
		return str
	}
	return strings.Repeat(`0`, length-slen) + str
}

var romanFig = []int64{100000, 10000, 1000, 100, 10, 1}

var romanI, romanV map[int64]rune

func init() {
	// M == ↀ
	romanI = map[int64]rune{1: 'I', 10: 'X', 100: 'C', 1000: 'M', 10000: 'ↂ', 100000: 'ↈ'}
	romanV = map[int64]rune{1: 'V', 10: 'L', 100: 'D', 1000: 'ↁ', 10000: 'ↇ'}
}

// convert int64 to roman number
//  I.ToRoman(16)) // output "XVI"
func Roman(num int64) string {
	res := []rune{}
	x := ' '
	for _, z := range romanFig {
		digit := num / z
		i, v := romanI[z], romanV[z]
		switch digit {
		case 1:
			res = append(res, i)
		case 2:
			res = append(res, i, i)
		case 3:
			res = append(res, i, i, i)
		case 4:
			res = append(res, i, v)
		case 5:
			res = append(res, v)
		case 6:
			res = append(res, v, i)
		case 7:
			res = append(res, v, i, i)
		case 8:
			res = append(res, v, i, i, i)
		case 9:
			res = append(res, i, x)
		}
		num -= digit * z
		x = i
	}
	return string(res)
}
