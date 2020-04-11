package bot

import (
	"errors"
	"os"
	"path"
	"plugin"
	"strings"

	"github.com/go-snart/snart/lib/errs"
)

func (b *Bot) Register(dir, name string) error {
	_f := "(*Bot).Register"
	Log.Infof(_f, "reg plug %s", name)

	filename := path.Join(dir, name+".so")
	Log.Infof(_f, "path %s", filename)

	plug, err := plugin.Open(filename)
	if err != nil {
		errs.Wrap(&err, `plugin.Open("%s")`, filename)
		Log.Error(_f, err)
		return err
	}

	iregf, err := plug.Lookup("Register")
	if err != nil {
		errs.Wrap(&err, `plug.Lookup("Register")`)
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
		errs.Wrap(&err, `os.Open(%#v)`, dir)
		Log.Error(_f, err)
		return err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			errs.Wrap(&err, `f.Close()`)
			Log.Error(_f, err)
		}
	}()

	files, err := f.Readdir(-1)
	if err != nil {
		errs.Wrap(&err, `f.Readdir(-1)`)
		Log.Error(_f, err)
		return err
	}

	for _, file := range files {
		fname := file.Name()
		if !strings.HasSuffix(fname, ".so") {
			continue
		}
		name := strings.TrimSuffix(fname, ".so")
		err = b.Register(dir, name)
		if err != nil {
			errs.Wrap(&err, `b.Register(%#v, %#v)`, dir, name)
			Log.Error(_f, err)
			return err
		}
	}

	return nil
}
