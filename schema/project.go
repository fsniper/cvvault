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

	"github.com/spf13/viper"
	"gitlab.com/metakeule/scaffold/lib/scaffold"
)

type Project struct {
	Name   string
	Basics Basics
	Works  []Work
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
url: ""
startDate: ""
endDate: ""
summary: ""
highlights":
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

	projects_path := viper.GetString("projects_directory")
	err := scaffold.Run(projects_path, template, ioreader, os.Stdout, false)
	if err != nil {
		log.Fatal(err)
	}

}

func (p *Project) Read() {
	fmt.Println("parsing project basics:", p.Name)
	/* if p.Basics == nil {
		p.Basics = Basics{}
	}*/
	err := p.Basics.Read(p.Name)
	if err != nil {
		fmt.Println("Error parsing project basics:", err)
		return
	}
}

func (p Project) GetAll() ([]Project, error) {
	projects := []Project{}

	projects_directory := viper.GetString("projects_directory")
	files, err := ioutil.ReadDir(projects_directory)
	if err != nil {
		fmt.Println("Error reading directory", err)
		return projects, err
	}

	for _, file := range files {
		if file.IsDir() {
			project := Project{
				Name: file.Name(),
			}
			project.Read()
			projects = append(projects, project)
		}
	}
	return projects, nil
}
