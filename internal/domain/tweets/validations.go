package tweets

import (
	"errors"
)

func validateTweet(content string, required bool) error {
	if required && len(content) == 0 {
		return errors.New("el tweet es requerido")
	}
	if len(content) > 0 && len(content) < 1 {
		return errors.New("el tweet debe tener al menos 1 caracter")
	} 
	if len(content) > 0 && len(content) > 280 {
		return errors.New("el tweet debe tener menos de 280 caracteres")
	}

	return nil
}