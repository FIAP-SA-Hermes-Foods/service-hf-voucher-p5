package valueobject

import (
	"errors"
	"time"
)

type ExpiresAt struct {
	Value *time.Time `json:"expiresAt,omitempty"`
}

func (e *ExpiresAt) Validate() error {
	if e.Value == nil {
		return errors.New("is requred a valid expiration time value")
	}

	if len(e.Format()) == 0 || e.Format() == "null" {
		return errors.New("is requred a valid expiration time value")
	}

	return nil
}

var expiresAtFormatLayout string = `02-01-2006 15:04:05`

func (e *ExpiresAt) Format() string {
	if e.Value == nil {
		return "null"
	}

	return e.Value.Format(expiresAtFormatLayout)
}

var (
	expiresAtSaveFromLayout string = `02-01-2006 15:04:05`
	expiresAtSaveToLayout   string = `2006-01-02 15:04:05.999999`
)

func (e *ExpiresAt) SetTimeFromString(du string) error {
	if len(du) == 0 {
		return nil
	}

	if e.Value == nil {
		return errors.New("is not possible set time at expiresAt because value is null")
	}

	t, err := time.Parse(expiresAtSaveFromLayout, du)

	if err != nil {
		return err
	}

	fmtT := t.Format(expiresAtSaveToLayout)

	tt, err := time.Parse(expiresAtSaveToLayout, fmtT)

	if err != nil {
		return err
	}

	e.Value = &tt

	return nil
}
