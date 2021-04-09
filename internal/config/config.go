package config

import (
	"errors"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	GuildID string
	Debug   bool

	IllegalWords []string
	AutoReplyWithBuild []string
}

var emptyConfig = Config{
	GuildID:      "",
	Debug:        false,
	IllegalWords: []string{},
	AutoReplyWithBuild: []string{},
}

var C Config

func Load() error {
	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	if _, err = os.Stat(path.Join(wd, "config.toml")); os.IsNotExist(err) {
		b, err := toml.Marshal(emptyConfig)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(path.Join(wd, "config.toml"), b, 0644)
		if err != nil {
			return err
		}

		return errors.New("no config was found one has been created")
	}

	b, err := ioutil.ReadFile(path.Join(wd, "config.toml"))
	if err != nil {
		return err
	}

	C = Config{}
	err = toml.Unmarshal(b, &C)

	if err != nil {
		return err
	}

	return nil
}
