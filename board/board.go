package board

import (
	"github.com/jarek7410/bridge_calculating_box/contract"
	"github.com/jarek7410/bridge_calculating_box/nsew"
	"github.com/jarek7410/bridge_calculating_box/result"
	"strconv"
)

type Board struct {
	Contract contract.Contract
	Result   result.Result
	NSEW     nsew.Wind
	Board    int16 // (mod 16)+1
}

func (b Board) String() string {
	return strconv.Itoa(int(b.Board)) + ", " +
		b.Contract.String() + ", " +
		b.NSEW.String() + ", " +
		b.Result.String()
}

func (b *Board) IsVulnerable() bool {
	if b.NSEW == nsew.North || b.NSEW == nsew.South {
		return b.Board == 2 || b.Board == 5 || b.Board == 12 || b.Board == 15 ||
			b.Board == 4 || b.Board == 7 || b.Board == 10 || b.Board == 13
	}
	if b.NSEW == nsew.East || b.NSEW == nsew.West {
		return b.Board == 3 || b.Board == 6 || b.Board == 9 || b.Board == 16 ||
			b.Board == 4 || b.Board == 7 || b.Board == 10 || b.Board == 13
	}
	return false
}

func (b *Board) Points() (points int) {
	val := b.IsVulnerable()
	if b.Result.Outcome == result.Minus {
		if b.Contract.Doubled == contract.None {
			if val {
				points = -100
			} else {
				points = -50
			}
			return points * b.Result.Level
		} else if b.Contract.Doubled == contract.X {
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
		} else if b.Contract.Doubled == contract.XX {
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
	if b.Contract.Color == contract.Spade ||
		b.Contract.Color == contract.Hearts ||
		b.Contract.Color == contract.NT {
		points = 30 * (b.Contract.Level)
		if b.Contract.Color == contract.NT {
			points += 10
		}
		if b.Contract.Doubled == contract.X {
			points *= 2
		}
		if b.Contract.Doubled == contract.XX {
			points *= 4
		}
	}
	if b.Contract.Color == contract.Clubs ||
		b.Contract.Color == contract.Diamonds {
		points = 20 * (b.Contract.Level)
		if b.Contract.Doubled == contract.X {
			points *= 2
		}
		if b.Contract.Doubled == contract.XX {
			points *= 4
		}
	}
	if b.Contract.Color == contract.NT {
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
	if b.Contract.Color == contract.Spade ||
		b.Contract.Color == contract.Hearts {
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
	if b.Contract.Color == contract.Clubs ||
		b.Contract.Color == contract.Diamonds {
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
	if b.Contract.Doubled == contract.X {
		points += 50
	}
	if b.Contract.Doubled == contract.XX {
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
	if b.Result.Outcome == result.Plus {

		if b.Contract.Doubled == contract.None {
			if b.Contract.Color == contract.Spade ||
				b.Contract.Color == contract.Hearts ||
				b.Contract.Color == contract.NT {
				points += 30 * b.Result.Level

			}
			if b.Contract.Color == contract.Clubs ||
				b.Contract.Color == contract.Diamonds {
				points += 20 * b.Result.Level
			}
		} else if b.Contract.Doubled == contract.X {
			if val {
				points += 200 * b.Result.Level
			} else {
				points += 100 * b.Result.Level
			}
		} else if b.Contract.Doubled == contract.XX {
			if val {
				points += 400 * b.Result.Level
			} else {
				points += 200 * b.Result.Level
			}
		}
	}

	return points
}
