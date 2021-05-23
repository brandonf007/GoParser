package appsettings

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/vrischmann/envconfig"
)

// ReadFromFileAndEnv reads the settings from a local file and applies any existing environment variables to it.
// This is taken from https://travix.io/making-your-go-app-configurable-bb5e5f4a9df9
func ReadFromFileAndEnv(settings interface{}) error {
	file, err := os.Open("./appsettings.json")

	if err != nil {
		return err
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return errors.Wrap(err, "Failed to read appsettings")
	}

	err = json.Unmarshal(data, settings)

	if err != nil {
		return errors.Wrap(err, "Failed to unmarshal appsettings")
	}
	// Begin Validation appsettings.json
	validate := settings.(*AppSettings)
	if _, err := validateAppSetting(validate); err != nil {
		return errors.Wrap(err, "Failed to validate appsettings.json file")
	}
	// End Validation

	err = envconfig.Init(settings)

	if err != nil {
		err = errors.Wrap(err, "Failed to update with env vars")
	}
	return err
}

func validateAppSetting(validate *AppSettings) (bool, error) {
	// Ensure that the log file is being passed
	if validate.InputFilePath == "" {
		return false, errors.Wrap(errors.New("App Settings - error"), "InputFilePath must have a valid filepath to a .log or .txt file to read in the logs to be parsed")
	}
	// Ensure regulare expression is set and not empty
	if validate.RegularExpression == "" {
		return false, errors.Wrap(errors.New("Regular Expression - error"), "Provided Regular Expression pattern can not be empty, it should contain two groups the first identifying the IP Address, and the second identifying the URL")
	}

	if validate.Output.NumberToDisplay <= 0 {
		return false, errors.Wrap(errors.New("App Settings - error"), "Output.NumberToDisplay must be set to greater than 0")
	}
	return true, nil
}
