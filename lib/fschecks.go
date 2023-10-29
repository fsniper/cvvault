package lib

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func Create_config() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Could not get home directory ", err)
	}

	path := filepath.Join(home, ".config", "cvvault")
	log.Println("Creating Config file in ", path)
	err = os.MkdirAll(path, 0700)
	if err != nil {
		log.Fatal("Could not create config directory ", err)
	}

	err = viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("Could not write config ", err)
	}
}
