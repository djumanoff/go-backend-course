package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrWrongSizeFormat = errors.New("Wrong format for size: (100x100x100:mm)")
)

var sizeRegex, _ = regexp.Compile(`(\d+)x(\d+)x(\d+):([a-z]+)`)

type Size struct {
	Top  int
	Middle int
	Bottom int
	Unit   SizeUnit
}

func (a Size) String() string {
	return fmt.Sprintf("%dx%dx%d:%s", a.Top, a.Middle, a.Bottom, a.Unit)
}

func (a *Size) FromString(b string) (err error) {
	result := sizeRegex.FindStringSubmatch(b)
	if len(result) <= 0 {
		return ErrWrongSizeFormat
	}

	a.Top, err = strconv.Atoi(result[1])
	if err != nil {
		return err
	}

	a.Middle, err = strconv.Atoi(result[2])
	if err != nil {
		return err
	}

	a.Bottom, err = strconv.Atoi(result[3])
	if err != nil {
		return err
	}

	nth, err := a.Unit.FromString(result[4])
	if err != nil {
		return err
	}
	a.Unit = nth

	return nil
}

func (a *Size) UnmarshalJSON(b []byte) error {
	var s string

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if err := a.FromString(s); err != nil {
		return err
	}

	return nil
}

func (a Size) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

