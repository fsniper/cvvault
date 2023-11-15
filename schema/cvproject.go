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

	embedcontent "github.com/fsniper/cvvault/emb"
	"github.com/fsniper/cvvault/lib"
	"gopkg.in/yaml.v2"

	"github.com/spf13/viper"
	"gitlab.com/metakeule/scaffold/lib/scaffold"
)

type CVProjectMeta struct {
	Read bool
	Path string
}

type CVProject struct {
	Meta      CVProjectMeta `json:"-"`
	Name      string
	Basics    Basics
	Works     []Work
	Volunteer []Volunteer
	Education []Education
}

func (p *CVProject) GetFullPath() string {
	cvprojects_path := viper.GetString("projects_directory")
	return fmt.Sprintf("%s/%s", cvprojects_path, p.Name)
}

func (p *CVProject) Create() {

	log.Println("Creating CVProject: ", p.Name, p.GetFullPath())

	template, err := embedcontent.EmbeddedContent.ReadFile("project.tmpl")

	ioreader := new(bytes.Buffer)
	json.NewEncoder(ioreader).Encode(p)

	cvprojectsPath := viper.GetString("projects_directory")
	err = scaffold.Run(cvprojectsPath, string(template), ioreader, os.Stdout, false)
	if err != nil {
		log.Fatal(err)
	}

}

func (p *CVProject) SetDefault() {
	viper.Set("cvproject.default", p.Name)
	viper.WriteConfig()
}

func (p *CVProject) IsDefault() bool {
	default_cvproject := viper.GetString("cvproject.default")
	return (p.Name == default_cvproject)
}

func (p *CVProject) Read() error {
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

func (p *CVProject) Export(ignoreTags []string, templateUrl string) {

	err := p.Read()
	if err != nil {
		log.Fatal("Error reading cvproject ", err)
	}

	for w, _ := range p.Works {
		p.Works[w].Filter(ignoreTags)
	}

	y, err := json.Marshal(p)
	if err != nil {
		log.Fatal("error yaml marhal: ", err)
	}

	fmt.Println(string(y))

	path := lib.CloneGitRepo(templateUrl)
	log.Println(path)

	//path := filepath.Join(projectsDirectory, projectName, "data", "basics.yaml")

	//yamlContent, err := ioutil.ReadFile(path)
	//var css = fs.readFileSync(__dirname+"/style.css", "utf-8")
	//var tpl = fs.readFileSync(__dirname+"/resume.hbs", "utf-8")
}

func (p *CVProject) Print() {

	err := p.Read()
	if err != nil {
		log.Fatal("Error reading cvproject ", err)
	}

	y, err := yaml.Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
}

func (p CVProject) GetAll() ([]CVProject, error) {
	cvprojects := []CVProject{}

	cvprojectsDirectory := viper.GetString("projects_directory")
	files, err := ioutil.ReadDir(cvprojectsDirectory)
	if err != nil {
		fmt.Println("Error reading directory", err)
		return cvprojects, err
	}

	for _, file := range files {
		if file.IsDir() {
			cvproject := CVProject{
				Name: file.Name(),
				Meta: CVProjectMeta{
					Read: false,
					Path: filepath.Join(cvprojectsDirectory, file.Name()),
				},
			}
			cvproject.Read()
			cvprojects = append(cvprojects, cvproject)
		}
	}
	return cvprojects, nil
}
