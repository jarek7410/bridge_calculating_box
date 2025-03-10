package result

import (
	"errors"
	"strconv"
)

type Outcome int

const (
	Minus Outcome = iota - 1
	Equal
	Plus
)

type Result struct {
	Outcome Outcome
	Level   int
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
