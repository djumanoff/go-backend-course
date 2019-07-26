package main

import (
	"errors"
	"encoding/json"
)

var (
	ErrSizeUnit = errors.New("Wrong size unit (mm,cm,m)")
)

type SizeUnit int

const (
	UNKNOWN SizeUnit = iota
	MM
	CM
	M
)

func (a SizeUnit) String() string {
	switch a {
	case MM:
		return "mm"
	case CM:
		return "cm"
	case M:
		return "m"
	}
	return "unknown"
}

func (a *SizeUnit) FromString(b string) (SizeUnit, error) {
	switch b {
	case "mm":
		return MM, nil
	case "cm":
		return CM, nil
	case "m":
		return M, nil
	}
	return UNKNOWN, ErrSizeUnit
}

func (a *SizeUnit) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	nt, err := a.FromString(s)
	if err != nil {
		return err
	}
	*a = nt

	return nil
}

func (a SizeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}
