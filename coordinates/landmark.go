package geo

import "errors"

//keep the field in lower case so it can't be exported directlym, only via getter and setter methods/functions
type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) string {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}
