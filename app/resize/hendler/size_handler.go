package hendler

import (
	_ "photobank-item-photo/app/pkg/resize"
	. "photobank-item-photo/app/pkg/resize/models"
)

func ControlSize(size uint) error {
	m := ValidConfig{Size: size}
	err := m.ValidConfigSize()
	if err != nil {
		return err
	}
	return nil
}
