package drunkenbishop

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Position struct {
	x, y int
}

type DrunkenBishop struct {
	fingerprint string
	topLabel string
	bottomLabel string
	w, h int
	startPos, currPos Position
	board [][]int
}

func NewDrunkenBishop() DrunkenBishop {
	db := DrunkenBishop{}
	return db
}

func (db *DrunkenBishop) Reset() {
	db.w = 17 // OpenSSH standard dimensions
	db.h = 9
	db.startPos = Position{
		x: int(db.w / 2),
		y: int(db.h / 2),
	}
	db.currPos = db.startPos
	db.board = make([][]int, db.w)

	for x := 0; x < db.w; x++ {
		db.board[x] = make([]int, db.h)
	}
}

func (db *DrunkenBishop) SetTopLabel(label string) {
	db.topLabel = label // db.generateLabel(label)
}

func (db *DrunkenBishop) SetBottomLabel(label string) {
	db.bottomLabel = label // db.generateLabel(label)
}

func (db *DrunkenBishop) ToAscii(fingerprint string) (string, error) {
	db.Reset()
	db.fingerprint = fingerprint
	err := db.walkFingerprint()
	if err != nil {
		return "", err
	}

	cm := db.getCharacterMap()
	var sb strings.Builder
	sb.WriteString(db.generateLabel(db.topLabel) + "\n")
	for y := 0; y < db.h; y++ {
		sb.WriteRune('|')
		for x := 0; x < db.w; x++ {
			var c rune
			if x == db.startPos.x && y == db.startPos.y {
				c = 'S'
			} else if x == db.currPos.x && y == db.currPos.y {
				c = 'E'
			} else {
				c = cm[db.board[x][y]%15]
			}
			sb.WriteRune(c)
		}
		sb.WriteRune('|')
		sb.WriteString("\n")
	}
	sb.WriteString(db.generateLabel(db.bottomLabel) + "\n")
	return sb.String(), nil
}

func (db *DrunkenBishop) generateLabel(label string) string {
	if len(label) > db.w - 3 {
		label = fmt.Sprintf("%s", label[:db.w-3])
	}
	if len(label) > 0 {
		label = fmt.Sprintf("[%s]", label)
	}
	np := db.w - len(label)
	lp := int(float64(np) / 2.0)
	rp := int(float64(np) / 2.0 + 0.5)
	padded := fmt.Sprintf("+%s%s%s+", strings.Repeat("-", lp), label, strings.Repeat("-", rp))
	return padded
}

func (db *DrunkenBishop) getCharacterMap() []rune {
	return []rune {
		0: ' ', 1: '.', 2: 'o', 3: '+',
		4: '=', 5: '*', 6: 'B', 7: 'O',
		8: 'X', 9: '@', 10: '%', 11: '&',
		12: '#', 13: '/', 14: '^',
	}
}

func (db *DrunkenBishop) walkFingerprint() error {
	fp := db.fingerprint
	fp = strings.Replace(fp, ":", "", -1)
	fp = strings.Replace(fp, "-", "", -1)

	fpBytes, err := hex.DecodeString(fp)
	if err != nil {
		return err
	}

	for _, b := range fpBytes {
		for i := 0; i < 4; i++ {
			dx := -1
			dy := -1
			if b>>(i*2) & 0b1 == 1 { // left - right bit
				dx = 1
			}
			if b>>(i*2+1) & 0b1 == 1 { // up - down bit
				dy = 1
			}
			x := db.currPos.x + dx
			y := db.currPos.y + dy
			if x < 0 {
				x = 0
			}
			if y < 0 {
				y = 0
			}
			if x == db.w {
				x = db.w - 1
			}
			if y == db.h {
				y = db.h - 1
			}
			db.currPos = Position {x, y}
			db.board[x][y] += 1
		}
	}
	return nil
}
