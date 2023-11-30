package rtp

import "strconv"

type Amount float64

func (a Amount) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(a), 'f', -1, 64)), nil
}

func (a Amount) Validate() error {
	_, err := a.MarshalText()
	return err
}
