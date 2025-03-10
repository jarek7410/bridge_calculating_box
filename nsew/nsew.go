package nsew

import "errors"

type Wind int

const (
	North = iota
	South
	East
	West
)

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
