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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gitlab.com/metakeule/scaffold/lib/scaffold"
	"gopkg.in/yaml.v2"
)

type ProjectMeta struct {
	Read bool
	Path string
}

type Project struct {
	Meta      ProjectMeta `json:"-"`
	Name      string
	Basics    Basics
	Works     []Work
	Volunteer []Volunteer
	Education []Education
}

func (p *Project) GetFullPath() string {
	projects_path := viper.GetString("projects_directory")
	return fmt.Sprintf("%s/%s", projects_path, p.Name)
}

func (p *Project) Create() {

	log.Println("Creating Project: ", p.Name, p.GetFullPath())

	template := `
>>>{{.Name}}/
>>>data/
>>>basics.yaml
name: "{{.Basics.name}}"
label: "{{.Basics.label}}"
image: "{{.Basics.image}}"
email: "{{.Basics.email}}"
phone: "{{.Basics.phone}}"
url: "{{.Basics.url}}"
summary: "{{.Basics.summary}}"
location:
  address: "{{.Basics.location.address}}"
  postalCode: "{{.Basics.location.postalCode}}"
  city: "{{.Basics.location.city}}"
  countryCode: "{{.Basics.location.countryCode}}"
  region: "{{.Basics.location.region}}"
<<<basics.yaml
>>>works/
>>>my-first-company-STARTYEAR.yaml
name: "My First Company"
location: ""
description: ""
position: ""
#url: ""
startDate: ""
endDate: ""
summary: ""
highlights:
- description: "" 
	tags: ["",""]
<<<my-first-company-STARTYEAR.yaml
>>>exports/
<<<exports/
<<<works/
<<<data/
<<<{{.Name}}/
`
	ioreader := new(bytes.Buffer)
	json.NewEncoder(ioreader).Encode(p)

	projects_path := viper.GetString("projectsDirectory")
	err := scaffold.Run(projects_path, template, ioreader, os.Stdout, false)
	if err != nil {
		log.Fatal(err)
	}

}

func (p *Project) SetDefault() {
	viper.Set("project.default", p.Name)
	viper.WriteConfig()
}

func (p *Project) IsDefault() bool {
	default_project := viper.GetString("project.default")
	return (p.Name == default_project)
}

func (p *Project) Read() error {
	if p.Meta.Read == true {
		return nil
	}
	err := p.Basics.Read(p.Name)
	if err != nil {
		return err
	}

	works, err := Work{}.GetAll(p.Name)
	if err != nil {
		return err
	}
	p.Works = works
	p.Meta.Read = true
	return nil
}

func (p *Project) Print() {

	err := p.Read()
	if err != nil {
		log.Fatal("Error reading project ", err)
	}

	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
}

func (p Project) GetAll() ([]Project, error) {
	projects := []Project{}

	projectsDirectory := viper.GetString("projects_directory")
	files, err := ioutil.ReadDir(projectsDirectory)
	if err != nil {
		fmt.Println("Error reading directory", err)
		return projects, err
	}

	for _, file := range files {
		if file.IsDir() {
			project := Project{
				Name: file.Name(),
				Meta: ProjectMeta{
					Read: false,
					Path: filepath.Join(projectsDirectory, file.Name()),
				},
			}
			project.Read()
			projects = append(projects, project)
		}
	}
	return projects, nil
}
