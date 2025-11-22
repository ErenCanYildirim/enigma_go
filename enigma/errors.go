package enigma

import "fmt"

type ErrInvalidRotorType string //returned when an invalid rotor type is specified

func (e ErrInvalidRotorType) Error() string {
	return fmt.Sprintf("invalid rotor type: %s", string(e))
}

// ErrInvalidReflectorType is returned when an invalid reflector type is specified
type ErrInvalidReflectorType string

func (e ErrInvalidReflectorType) Error() string {
	return fmt.Sprintf("invalid reflector type: %s", string(e))
}
