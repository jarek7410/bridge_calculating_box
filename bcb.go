package bridge_calculating_box

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Version string = "ack"

func VersionName() {
	fmt.Printf("%s\n", Version)

}

type Color int64

const (
	Spade Color = iota
	Hearts
	Clubs
	Diamonds
	NT
)

type Doubled int64

const (
	None Doubled = iota
	X
	XX
)

type Wind int

const (
	North Wind = iota
	South
	East
	West
)

type Outcome int

const (
	Minus Outcome = iota - 1
	Equal
	Plus
)

type Board struct {
	Contract Contract
	Result   Result
	NSEW     Wind
	Board    int16 // (mod 16)+1
}
type Contract struct {
	Level   int
	Color   Color
	Doubled Doubled
}
type Result struct {
	Outcome Outcome
	Level   int
}

func (w Wind) String() string {
	switch w {
	case North:
		return "N"
	case South:
		return "S"
	case East:
		return "E"
	case West:
		return "W"
	}
	return "??"
}

func (w *Wind) Parse(windd string) (Wind, error) {
	switch windd {
	case "N":
		*w = North
		return North, nil
	case "S":
		*w = South
		return South, nil
	case "E":
		*w = East
		return East, nil
	case "W":
		*w = West
		return West, nil
	default:
		return -1, errors.New("invalid wind")
	}

}

func (o Outcome) String() string {
	switch o {
	case Minus:
		return "-"
	case Equal:
		return "="
	case Plus:
		return "+"
	}
	return "??"
}

func (r Result) String() string {
	if r.Outcome == Equal {
		return "="
	}
	return r.Outcome.String() + strconv.Itoa(r.Level)
}

func (r *Result) Parse(Result string) error {
	if len(Result) == 1 && Result[0] == '=' {
		r.Outcome = Equal
		return nil
	}
	if len(Result) == 1 {
		return errors.New("Invalid result")
	}
	switch Result[0] {
	case '-':
		r.Outcome = Minus
	case '+':
		r.Outcome = Plus
	default:
		return errors.New("Invalid result")
	}
	Level, err := strconv.Atoi(Result[1:])
	if err != nil {
		return errors.New("Invalid result")
	}
	r.Level = Level
	return nil
}

func (c Color) String() string {
	switch c {
	case Spade:
		return "S"
	case Hearts:
		return "H"
	case Clubs:
		return "C"
	case Diamonds:
		return "D"
	case NT:
		return "NT"
	}
	return "??"
}
func (d Doubled) String() string {
	switch d {
	case None:
		return ""
	case X:
		return "x"
	case XX:
		return "xx"
	}
	return "??"

}

func (c Contract) String() string {
	return strconv.Itoa(c.Level) + " " + c.Color.String() + " " + c.Doubled.String()
}

func (c *Contract) Parse(constract string) (err error) {
	parts := strings.Split(constract, " ")
	c.Level, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}

	switch parts[1] {
	case "S":
		c.Color = Spade
	case "H":
		c.Color = Hearts
	case "C":
		c.Color = Clubs
	case "D":
		c.Color = Diamonds
	case "NT":
		c.Color = NT
	default:
		return errors.New("Invalid color")
	}

	if len(parts) == 3 {
		switch parts[2] {
		case "":
			c.Doubled = None
		case "x":
			c.Doubled = X
		case "xx":
			c.Doubled = XX
		default:
			return errors.New("Invalid doubled")
		}
	} else {
		c.Doubled = None
	}

	return nil
}
func Imps(point int) (imps int) {
	return GiveMeImps(point)
}
func GiveMeImps(points int) (imps int) {
	if points < 0 {
		return giveMeImps(points)*-1
	}
	return giveMeImps(points)
}

func giveMeImps(points int) (imps int) {

	if points < 0 {
		points = -points
	}
	if points < 20 {
		return 0
	}
	if points < 50 {
		return 1
	}
	if points < 90 {
		return 2
	}
	if points < 130 {
		return 3
	}
	if points < 170 {
		return 4
	}
	if points < 220 {
		return 5
	}
	if points < 270 {
		return 6
	}
	if points < 320 {
		return 7
	}
	if points < 370 {
		return 8
	}
	if points < 430 {
		return 9
	}
	if points < 500 {
		return 10
	}
	if points < 600 {
		return 11
	}
	if points < 750 {
		return 12
	}
	if points < 900 {
		return 13
	}
	if points < 1100 {
		return 14
	}
	if points < 1300 {
		return 15
	}
	if points < 1500 {
		return 16
	}
	if points < 1750 {
		return 17
	}
	if points < 2000 {
		return 18
	}
	if points < 2250 {
		return 19
	}
	if points < 2500 {
		return 20
	}
	if points < 3000 {
		return 21
	}
	if points < 3500 {
		return 22
	}
	if points < 4000 {
		return 23
	}
	return 24 //max as per The Laws of Duplicate Bridge 2017
}

func (b Board) String() string {
	return strconv.Itoa(int(b.Board)) + ", " +
		b.Contract.String() + ", " +
		b.NSEW.String() + ", " +
		b.Result.String()
}

func (b *Board) IsVulnerable() bool {
	if b.NSEW == North || b.NSEW == South {
		return b.Board == 2 || b.Board == 5 || b.Board == 12 || b.Board == 15 ||
			b.Board == 4 || b.Board == 7 || b.Board == 10 || b.Board == 13
	}
	if b.NSEW == East || b.NSEW == West {
		return b.Board == 3 || b.Board == 6 || b.Board == 9 || b.Board == 16 ||
			b.Board == 4 || b.Board == 7 || b.Board == 10 || b.Board == 13
	}
	return false
}

//points for North-South pair (negative for East-West)
func (b *Board) Points() (points int) {
	if b.NSEW == North || b.NSEW == South {
		return b.points()
	}
	return b.points()*-1
}

func (b *Board) points() (points int) {
	val := b.IsVulnerable()
	if b.Result.Outcome == Minus {
		if b.Contract.Doubled == None {
			if val {
				points = -100
			} else {
				points = -50
			}
			return points * b.Result.Level
		} else if b.Contract.Doubled == X {
			if val {
				points = -200
			} else {
				points = -100
			}
			var lew int
			if val {
				lew = -300
			} else {
				lew = -200
			}
			if b.Result.Level >= 4 && !val {
				points += -100 * (b.Result.Level - 3)
			}
			return points + (b.Result.Level-1)*lew
		} else if b.Contract.Doubled == XX {
			if val {
				points = -400
			} else {
				points = -200
			}
			var lew int
			if val {
				lew = -600
			} else {
				lew = -400
			}
			if b.Result.Level >= 4 && !val {
				points += -200 * (b.Result.Level - 3)
			}
			return points + (b.Result.Level-1)*lew
		}
	}
	//punkty za lewy:
	if b.Contract.Color == Spade ||
		b.Contract.Color == Hearts ||
		b.Contract.Color == NT {
		points = 30 * (b.Contract.Level)
		if b.Contract.Color == NT {
			points += 10
		}
		if b.Contract.Doubled == X {
			points *= 2
		}
		if b.Contract.Doubled == XX {
			points *= 4
		}
	}
	if b.Contract.Color == Clubs ||
		b.Contract.Color == Diamonds {
		points = 20 * (b.Contract.Level)
		if b.Contract.Doubled == X {
			points *= 2
		}
		if b.Contract.Doubled == XX {
			points *= 4
		}
	}
	if b.Contract.Color == NT {
		if b.Contract.Level < 3 {
			points += 50
		} else {
			if val {
				points += 500
			} else {
				points += 300
			}
		}
	}
	if b.Contract.Color == Spade ||
		b.Contract.Color == Hearts {
		if b.Contract.Level < 4 {
			points += 50
		} else {
			if val {
				points += 500
			} else {
				points += 300
			}
		}
	}
	if b.Contract.Color == Clubs ||
		b.Contract.Color == Diamonds {
		if b.Contract.Level < 5 {
			points += 50
		} else {
			if val {
				points += 500
			} else {
				points += 300
			}
		}
	}
	if b.Contract.Doubled == X {
		points += 50
	}
	if b.Contract.Doubled == XX {
		points += 100
	}

	if b.Contract.Level == 6 {
		if val {
			points += 750
		} else {
			points += 500
		}
	}
	if b.Contract.Level == 7 {
		if val {
			points += 1500
		} else {
			points += 1000
		}
	}
	//nadrubki
	if b.Result.Outcome == Plus {

		if b.Contract.Doubled == None {
			if b.Contract.Color == Spade ||
				b.Contract.Color == Hearts ||
				b.Contract.Color == NT {
				points += 30 * b.Result.Level

			}
			if b.Contract.Color == Clubs ||
				b.Contract.Color == Diamonds {
				points += 20 * b.Result.Level
			}
		} else if b.Contract.Doubled == X {
			if val {
				points += 200 * b.Result.Level
			} else {
				points += 100 * b.Result.Level
			}
		} else if b.Contract.Doubled == XX {
			if val {
				points += 400 * b.Result.Level
			} else {
				points += 200 * b.Result.Level
			}
		}
	}

	return points
}

func Gettrue