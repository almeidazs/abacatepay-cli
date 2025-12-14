package prompts

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func AskAPIKey() (string, error) {
	var useENV bool
	var key = os.Getenv("ABACATE_PAY_API_KEY")

	isUsingFromEnv := len(key) > 0

	if isUsingFromEnv {
		err := survey.AskOne(&survey.Confirm{
			Default: false,
			Message: "We detected an API key in ABACATE_PAY_API_KEY enviroment key. Do you want to use it?",
		}, &useENV)

		if err != nil {
			return "", err
		}
	}

	if !useENV {
		err := survey.AskOne(&survey.Password{
			Message: "What's the API key?",
		}, &key, survey.WithValidator(survey.Required))

		if err != nil {
			return "", err
		}
	}

	return key, nil
}
