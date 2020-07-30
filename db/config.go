package db

import "fmt"

// Configs returns a list of useable db config strings.
func Configs() []string {
	const _f = "Configs"

	allConfs := []string(nil)

	confs, err := EnvConfigs()
	if err != nil {
		err = fmt.Errorf("env confs: %w", err)

		Log.Warn(_f, err)
	} else {
		allConfs = append(allConfs, confs...)
	}

	if len(allConfs) == 0 {
		confs, err = StdinConfigs()
		if err != nil {
			err = fmt.Errorf("stdin confs: %w", err)

			Log.Warn(_f, err)
		} else {
			allConfs = append(allConfs, confs...)
		}
	}

	return allConfs
}
