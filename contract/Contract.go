package contract

type Color string
type Doubled int

const (
	Spade    Color = "S"
	Hearts         = "H"
	Clubs          = "C"
	Diamonds       = "D"
	NT             = "NT"
)
const (
	x Doubled = iota
	xx
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
	default:
		return ""
	}
}
func ()  {
	
}