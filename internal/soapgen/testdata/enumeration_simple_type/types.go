package enumeration_simple_type

import (
	"encoding/xml"
)

// Enumeration types

// ColorType represents an enumeration type
type ColorType string

// ColorType enumeration values
const (
	ColorTypeRED   ColorType = "RED"
	ColorTypeGREEN ColorType = "GREEN"
	ColorTypeBLUE  ColorType = "BLUE"
)

// String returns the string representation of ColorType
func (e ColorType) String() string {
	return string(e)
}

// IsValid returns true if the ColorType value is valid
func (e ColorType) IsValid() bool {
	switch e {
	case ColorTypeRED, ColorTypeGREEN, ColorTypeBLUE:
		return true
	default:
		return false
	}
}
