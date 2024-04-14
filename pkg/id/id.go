package id

import (
	"fmt"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (u *ID) FromString(s string) error {
	id, err := uuid.Parse(s)
	if err != nil {
		return err
	}

	*u = ID(id)
	return nil
}

func (u ID) String() string {
	return uuid.UUID(u).String()
}

func (u *ID) Scan(data any) error {
	return scanUUID((*uuid.UUID)(u), data)
}

func (u ID) MarshalText() ([]byte, error) {
	return []byte(uuid.UUID(u).String()), nil
}

func (u *ID) UnmarshalText(data []byte) error {
	return unmarshalUUID((*uuid.UUID)(u), data)
}

func scanUUID(u *uuid.UUID, data any) error {
	if err := u.Scan(data); err != nil {
		return fmt.Errorf("scanning id value: %w", err)
	}
	return nil
}

func unmarshalUUID(u *uuid.UUID, data []byte) error {
	if err := u.UnmarshalText(data); err != nil {
		return fmt.Errorf("parsing value: %w", err)
	}
	return nil
}
