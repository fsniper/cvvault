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
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestProjectRead(t *testing.T) {

	viper.Set("projects_directory", "testdata")
	project := Project{Name: "test"}
	err := project.Read()
	if err != nil {
		t.Fatalf("Error reading project %v", err)
	}

	p := project.GetFullPath()
	if p != "testdata/test" {
		t.Errorf("Expected project directory to be 'testdata/test', but got %s", p)
	}
	if project.Basics.Name != "John Doe" {
		t.Errorf("Expected Name to be 'John Doe', but got %s", project.Basics.Name)
	}
	if !project.Meta.Read {
		t.Errorf("Expected project meta data to be set to read")
	}
	if project.Basics.Label != "Software Engineer" {
		t.Errorf("Expected Label to be 'Software Engineer', but got %s", project.Basics.Label)
	}
	if project.Basics.Profiles[0].Network != "linkedin" {
		t.Errorf("Expected first profile's network to be 'linkedin', but got %s", project.Basics.Label)
	}

	/* Let's test Print */
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	project.Print()
	w.Close()
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "name: John Doe") {
		t.Errorf("Expected Name in the output to be 'John Doe', but can't")
	}

	os.Stdout = rescueStdout

}
