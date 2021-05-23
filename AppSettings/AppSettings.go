package appsettings

import (
	"encoding/json"
	"fmt"
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
	fmt.Printf("%+v\n", settings)

	err = envconfig.Init(settings)

	if err != nil {
		err = errors.Wrap(err, "Failed to update with env vars")
	}
	return err
}
