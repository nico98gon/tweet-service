package tweets

import (
	"errors"
	"tweet-service/utils"
)

func CreateTweetValidations(t Tweet) error {
	var errorMessages []string

	if err := validateTweet(t.Content, true); err != nil {
		errorMessages = append(errorMessages, err.Error())
	}

	if len(errorMessages) > 0 {
		return errors.New("Errores de validación:\n" + utils.JoinErrors(errorMessages))
	}

	return nil
}