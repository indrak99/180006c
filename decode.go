package main

import (
	"fmt"
	"math"
)

/*
	3a7de8dc48d23a7d3a7d3a71 30000005 B05
	3a7d3a7d3aec48f9b57d3a72 10004444 A03
	00005555 00005555 A02
*/

// Letter2Number Convert Char to Number
func Letter2Number(chr rune) int {
	var num int
	switch chr {
	case '0':
		num = 16
	case '1':
		num = 15
	case '2':
		num = 1
	case '3':
		num = 2
	case '4':
		num = 14
	case '5':
		num = 13
	case '6':
		num = 4
	case '7':
		num = 3
	case '8':
		num = 5
	case '9':
		num = 6
	case 'A':
	case 'a':
		num = 8
	case 'B':
	case 'b':
		num = 7
	case 'C':
	case 'c':
		num = 10
	case 'D':
	case 'd':
		num = 9
	case 'E':
	case 'e':
		num = 12
	case 'F':
	case 'f':
		num = 11
	default:
		num = 0
	}
	return num
}

// PositionShift Get Position Shift
func PositionShift(str rune, ln int) int {
	i := Letter2Number(str)
	var (
		pos int
		shf int
	)

	switch str {
	case '5':
		pos = 0
	case '4':
		pos = 1
	case '6':
		pos = 2
	case '8':
		pos = 3
	case '7':
		pos = 4
	case 'A':
	case 'a':
		pos = 5
	case 'F':
	case 'f':
		pos = 6
	case 'D':
	case 'd':
		pos = 7
	case 'E':
	case 'e':
		pos = 8
	case '3':
		pos = 9
	case '1':
		pos = 10
	case '2':
		pos = 11
	case '0':
		pos = 12
	case '9':
		pos = 13
	case 'B':
	case 'b':
		pos = 14
	default:
		pos = 15
	}
	if pos == 0 {
		pos = 16
	}

	shf = i % pos

	if shf == 0 {
		shf = pos
	}

	if shf > ln {
		shf = shf % ln
	}

	return shf
}

// MapChar Mapping Character
func MapChar(chr rune) rune {
	var rchr rune

	switch chr {
	case '0':
		rchr = 'C'
	case '1':
		rchr = 'B'
	case '2':
		rchr = 'F'
	case '3':
		rchr = '1'
	case '4':
		rchr = '9'
	case '5':
		rchr = '2'
	case '6':
		rchr = 'A'
	case '7':
		rchr = 'E'
	case '8':
		rchr = '3'
	case '9':
		rchr = '4'
	default:
		rchr = 'D'
	}
	return rchr
}

// DeMapChar Mapping Character
func DeMapChar(chr rune) rune {
	var rchr rune
	switch chr {
	case 'C':
	case 'c':
		rchr = '0'
	case 'B':
	case 'b':
		rchr = '1'
	case 'F':
	case 'f':
		rchr = '2'
	case '1':
		rchr = '3'
	case '9':
		rchr = '4'
	case '2':
		rchr = '5'
	case 'A':
	case 'a':
		rchr = '6'
	case 'E':
	case 'e':
		rchr = '7'
	case '3':
		rchr = '8'
	case '4':
		rchr = '9'
	default:
		rchr = ' '
	}
	return rchr
}

// MapChar1 Mapping Character
func MapChar1(chr rune) rune {
	var rchr rune
	switch chr {
	case '0':
		rchr = '4'
	case '1':
		rchr = '6'
	case '2':
		rchr = 'A'
	case '3':
		rchr = 'E'
	case '4':
		rchr = 'B'
	case '5':
		rchr = 'D'
	case '6':
		rchr = '7'
	case '7':
		rchr = 'C'
	case '8':
		rchr = '0'
	case '9':
		rchr = '2'
	default:
		rchr = '3'
	}
	return rchr
}

// DeMapChar1 Mapping Character
func DeMapChar1(chr rune) rune {
	var rchr rune
	switch chr {
	case '4':
		rchr = '0'
	case '6':
		rchr = '1'
	case 'A':
	case 'a':
		rchr = '2'
	case 'E':
	case 'e':
		rchr = '3'
	case 'B':
	case 'b':
		rchr = '4'
	case 'D':
	case 'd':
		rchr = '5'
	case '7':
		rchr = '6'
	case 'C':
	case 'c':
		rchr = '7'
	case '0':
		rchr = '8'
	case '2':
		rchr = '9'
	default:
		rchr = ' '
	}
	return rchr
}

// MapChar2 Mapping Character
func MapChar2(chr rune) rune {
	var rchr rune
	switch chr {
	case '0':
		rchr = '8'
	case '1':
		rchr = '4'
	case '2':
		rchr = '3'
	case '3':
		rchr = '2'
	case '4':
		rchr = '5'
	case '5':
		rchr = '1'
	case '6':
		rchr = 'C'
	case '7':
		rchr = '9'
	case '8':
		rchr = '0'
	case '9':
		rchr = 'E'
	default:
		rchr = 'A'
	}
	return rchr
}

// DeMapChar2 Mapping Character
func DeMapChar2(chr rune) rune {
	var rchr rune
	switch chr {
	case '8':
		rchr = '0'
	case '4':
		rchr = '1'
	case '3':
		rchr = '2'
	case '2':
		rchr = '3'
	case '5':
		rchr = '4'
	case '1':
		rchr = '5'
	case 'C':
	case 'c':
		rchr = '6'
	case '9':
		rchr = '7'
	case '0':
		rchr = '8'
	case 'E':
	case 'e':
		rchr = '9'
	default:
		rchr = ' '
	}
	return rchr
}

// MapChar3 Mapping Character
func MapChar3(chr rune) rune {
	var rchr rune
	switch chr {
	case '0':
		rchr = 'D'
	case '1':
		rchr = 'E'
	case '2':
		rchr = 'C'
	case '3':
		rchr = 'B'
	case '4':
		rchr = 'F'
	case '5':
		rchr = 'A'
	case '6':
		rchr = '1'
	case '7':
		rchr = '2'
	case '8':
		rchr = '4'
	case '9':
		rchr = '6'
	default:
		rchr = '7'
	}
	return rchr
}

// DeMapChar3 Mapping Character
func DeMapChar3(chr rune) rune {
	var rchr rune
	switch chr {
	case 'D':
	case 'd':
		rchr = '0'
	case 'E':
	case 'e':
		rchr = '1'
	case 'C':
	case 'c':
		rchr = '2'
	case 'B':
	case 'b':
		rchr = '3'
	case 'F':
	case 'f':
		rchr = '4'
	case 'A':
	case 'a':
		rchr = '5'
	case '1':
		rchr = '6'
	case '2':
		rchr = '7'
	case '4':
		rchr = '8'
	case '6':
		rchr = '9'
	default:
		rchr = ' '
	}
	return rchr
}

//MitagReqInit Init Request
func (req *MiTagReq) MitagReqInit() {
	req.Key = 23
	req.Status = 0
	req.Valid = 1
	req.Decode = req.Encode
	if len(req.Encode) < 8 {
		req.Status = 1 // Min data length 8
	}
	if len(req.Encode) == 8 {
		req.Status = 2 // Eql data length 8
	}
	if len(req.Encode) > 8 {
		req.Status = 3 // More data length 8
	}
}

// MiTag MiTag information
func (req *MiTagReq) MiTag() {
	req.MitagReqInit()
	fmt.Println(req.Status, req.Encode, req.Decode)
	if req.Status == 3 {
		senc := req.Encode
		lenc := len(senc)
		shf := PositionShift(rune(senc[req.Key]), lenc)

		if shf == (lenc - 1) {
			shf = lenc
		} else {
			if shf == 1 {
				shf = int(math.Trunc(float64(lenc / 2)))
			}
		}
		dirc := shf % 2

		buf := make([]byte, lenc)
		for j := 1; j < lenc; j++ {
			k := (j) % 4
			switch k {
			case 0:
				buf[j-1] = byte(DeMapChar(rune(senc[j-1])))
			case 1:
				buf[j-1] = byte(DeMapChar1(rune(senc[j-1])))
			case 2:
				buf[j-1] = byte(DeMapChar2(rune(senc[j-1])))
			case 3:
				buf[j-1] = byte(DeMapChar3(rune(senc[j-1])))
			}
		}

		pos := 0
		s := ""
		for i := 1; i < lenc; i++ {
			switch dirc {
			case 0:
				pos = i + shf
				pos = (pos % lenc) + 1
			case 1:
				pos = i - shf + lenc
				pos = (pos % lenc) + 1
			}

			if buf[i-1] != 0x20 {
				s += string(buf[i-1])
			}
		}

		req.Decode = s
	}
}

// MiTagReq struct tag manage
type MiTagReq struct {
	Key    int
	Valid  int
	Encode string
	Decode string
	Status int
}

// MiTagRes struct manage tag response
type MiTagRes struct {
	Device string `json:"device"`
	Ukey   int    `json:"ukey"`
	Result string `json:"result"`
	Status string `json:"status"`
	Date   string `json:"dttm"`
}
