/*
Copyright © 2023 M.Onur YALAZI <onur.yalazi@gmail.com>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

 1. Redistributions of source code must retain the above copyright notice,
    this list of conditions and the following disclaimer.

 2. Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.

 3. Neither the name of the copyright holder nor the names of its contributors
    may be used to endorse or promote products derived from this software
    without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
*/
package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/fsniper/cvvault/cmd"
	"github.com/fsniper/cvvault/lib"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func initialize() {
	viper.SetConfigName("cvvault-config")        // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.config/cvvault") // call multiple times to add many search paths
	viper.AddConfigPath("/etc/cvvault/")         // path to look for the config file in
	viper.AddConfigPath(".")                     // optionally look for config in the working directory

	path, err := homedir.Expand("~/Documents/CVVault/projects")
	if err != nil {
		log.Fatal("Could not expand path for projects: ", err)
	}
	viper.SetDefault("projects_directory", path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config was not found")
			lib.CreateConfig()
		} else {
			log.Fatal("fatal error reading config file", err)
		}
	}
	log.Println("Config file in use:", viper.ConfigFileUsed())

	path = viper.GetString("projects_directory")
	lib.CheckDirectory(path)
}

func main() {
	initialize()
	cmd.Execute()
}
