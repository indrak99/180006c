package main

import "fmt"

/*
	3a7dc3a8c39c48c3923a7d3a7d3a71 30000005 B05
	3a7de8dc48d23a7d3a7d3a71
	3a7de8dc48d23a7d3a7d3a71
	f010ee01063a7de8dc48d23a7d3a7d3a715d
	3a7d3a7d3ac3ac48c3b9c2b57d3a72 10004444 A03
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

//MitagReqInit Init Request
func (req *MiTagReq) MitagReqInit() {
	req.Key = 23
	req.Status = 0
	req.Valid = 1
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
	req.Decode = req.Encode
	fmt.Println(req.Encode, req.Decode)
	/*
		var num int
		for _, c := range str {
			num = Letter2Number(c)
			fmt.Println(str)
		}
	*/
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
	Ukey   int    `json:"ukey"`
	Result string `json:"result"`
	Status string `json:"status"`
	Date   string `json:"dttm"`
}
