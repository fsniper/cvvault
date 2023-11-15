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

package lib

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
)

func getGitName(url string) string {
	re := regexp.MustCompile(`/([^/]+?)(?:\.git)?$`)

	match := re.FindStringSubmatch(url)
	if match != nil {
		return match[1]
	} else {
		log.Fatal("Can't get the name from the url")
		return ""
	}
}

func CloneGitRepo(url string) string {
	log.Print("Cloning template")
	templatesDirectory := viper.GetString("templates_directory")

	name := getGitName(url)
	path := filepath.Join(templatesDirectory, name)

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		repo, err := git.PlainOpen(path)
		if err != nil {
			log.Fatal("Could not open template git repo ", err)
		}

		worktree, err := repo.Worktree()
		err = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			log.Fatal("Could not pull template git repo ", err)
		}
	} else {
		// Repo does not exist, so clone it
		_, err := git.PlainClone(path, false, &git.CloneOptions{URL: url, Progress: os.Stdout})
		if err != nil {
			log.Fatal("Could not clone emplate git repo", err)
		}
		return path
	}

	return path
}
