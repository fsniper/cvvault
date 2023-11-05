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
package schema

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/spf13/viper"
)

type Work struct {
	Path        string `json:"-"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Position    string `json:"position"`
	Url         string `json:"url"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Summary     string `json:"summary"`
	Highlights  []struct {
		Description string   `json:"description"`
		Tags        []string `json:"-"`
	} `json:"highlights"`
}

func (w *Work) Read() error {
	yaml_content, err := ioutil.ReadFile(w.Path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yaml_content, &w)
	if err != nil {
		return err
	}
	return nil
}

func (w Work) GetAll(project_name string) ([]Work, error) {

	works := []Work{}
	projects_directory := viper.GetString("projects_directory")
	works_directory := filepath.Join(projects_directory, project_name, "data", "works")

	files, err := ioutil.ReadDir(works_directory)
	if err != nil {
		fmt.Println("Error reading directory", err)
		return works, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			work := Work{Path: filepath.Join(works_directory, file.Name())}
			err := work.Read()
			if err != nil {
				fmt.Println("Error reading work: ", work.Path)
				return works, err
			}
			works = append(works, work)
		}
	}

	less := func(i, j int) bool {
		const shortForm = "2006-01-02"
		ti, _ := time.Parse(shortForm, works[i].StartDate)
		tj, _ := time.Parse(shortForm, works[j].StartDate)
		return ti.Before(tj)
	}

	sort.Slice(works, less)

	return works, nil
}
