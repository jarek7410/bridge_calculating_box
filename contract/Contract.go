package contract

import (
	"errors"
	"strconv"
	"strings"
)

type Color int64
type Doubled int64

const (
	Spade Color = iota
	Hearts
	Clubs
	Diamonds
	NT
)
const (
	None Doubled = iota
	X
	XX
)

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

type Contract struct {
	Level   int
	Color   Color
	Doubled Doubled
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
