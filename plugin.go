package bot

import (
	"errors"
	"fmt"
	"os"
	"path"
	"plugin"
)

func (b *Bot) Register(dir, name string) error {
	_f := "(*Bot).Register"
	Log.Infof(_f, "reg plug %s", name)

	filename := path.Join(dir, name)
	Log.Infof(_f, "path %s", filename)

	plug, err := plugin.Open(filename)
	if err != nil {
		err = fmt.Errorf("plugin open %#v: %w", filename, err)
		Log.Error(_f, err)
		return err
	}

	iregf, err := plug.Lookup("Register")
	if err != nil {
		err = fmt.Errorf("lookup register: %w", err)
		Log.Error(_f, err)
		return err
	}

	regf, ok := iregf.(func(string, *Bot) error)
	if !ok {
		err := errors.New("failed to assert regf")
		Log.Error(_f, err)
		return err
	}

	return regf(name, b)
}

func (b *Bot) RegisterAll(dir string) error {
	_f := "(*Bot).RegisterAll"
	f, err := os.Open(dir)
	if err != nil {
		err = fmt.Errorf("open %#v: %w", dir, err)
		Log.Error(_f, err)
		return err
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		err = fmt.Errorf("readdir -1: %w", err)
		Log.Error(_f, err)
		return err
	}

	for _, file := range files {
		name := file.Name()
		_ = b.Register(dir, name)
	}

	return nil
}
