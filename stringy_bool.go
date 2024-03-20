package siwago

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StringyBool is a small utility type for decoding bools that might
// be encoded as the strings "true" or "false". Note that if you
// _encode_ with this type, you always get a real bool.
type StringyBool bool

func (b *StringyBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(*b))
}

func (b *StringyBool) UnmarshalJSON(buf []byte) error {
	var x interface{}
	err := json.Unmarshal(buf, &x)
	if err != nil {
		return err
	}
	switch xx := x.(type) {
	case bool:
		*b = StringyBool(xx)
	case string:
		s := strings.ToLower(xx)
		switch s {
		case "true":
			*b = true
		case "false":
			*b = false
		default:
			return fmt.Errorf("can't decode bool from JSON value %q", s)
		}
	default:
		return fmt.Errorf("can't decode JSON value of type %T into a bool", x)
	}
	return nil
}
