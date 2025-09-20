package designspreadsheet

import "strings"

type Spreadsheet struct {
	rows int
	val  map[string]int // sparse: "A1" -> value
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{rows: rows, val: make(map[string]int)}
}

func (s *Spreadsheet) SetCell(cell string, value int) {
	s.val[cell] = value
}

func (s *Spreadsheet) ResetCell(cell string) {
	s.val[cell] = 0
}

func (s *Spreadsheet) GetValue(formula string) int {
	// "=X+Y"
	body := formula[1:]
	i := strings.IndexByte(body, '+')
	left := body[:i]
	right := body[i+1:]
	return s.getToken(left) + s.getToken(right)
}

func (s *Spreadsheet) getToken(tok string) int {
	// число или ссылка
	if tok != "" && tok[0] >= '0' && tok[0] <= '9' {
		// parse int (в токенах нет знака и пробелов)
		n := 0
		for i := 0; i < len(tok); i++ {
			n = n*10 + int(tok[i]-'0')
		}
		return n
	}
	if v, ok := s.val[tok]; ok {
		return v
	}
	return 0
}
