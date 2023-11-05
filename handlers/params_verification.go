package handlers

import (
	"fmt"
)

func ValidateParams(params map[string]interface{}) error {
	seenKeys := make(map[string]bool)

	for key, value := range params {
		if key == "" {
			return fmt.Errorf("empty key found in parameters")
		}
		if seenKeys[key] {
			return fmt.Errorf("duplicate key found in parameters: %s", key)
		}
		if value == nil {
			return fmt.Errorf("parameter for key '%s' is nil", key)
		}
		seenKeys[key] = true
	}
	return nil
}
