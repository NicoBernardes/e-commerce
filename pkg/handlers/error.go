package handlers

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInternal error = errors.New("internal server error")

func ReturnError(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternal
	}
	return err
}
