package models

import (
	"errors"
)

type REsize struct {
	OutSize       []uint
	PathSavePhoto string
}
type ConfigResize struct {
	Resize REsize
}

type ValidConfig struct {
	Size uint
}

func (a *ValidConfig) ValidConfigSize() error {
	if a.Size >= 1 && a.Size <= 5000 {
		return nil
	} else {
		return errors.New("!")
	}
}
