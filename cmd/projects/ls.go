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
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"

	"github.com/fsniper/cvvault/schema"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		directory := "projects"
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			fmt.Println("Error reading directory \"projects\":", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Directory", "Project", "Name", "Label"})
		table.SetBorder(false)

		for _, file := range files {
			if file.IsDir() {
				project_name := file.Name()
				fmt.Println("parsing project basics:", project_name)
				file_path := filepath.Join(directory, project_name, "data", "basics.json")
				b := schema.Basics{}
				err := b.Read(file_path)
				if err != nil {
					fmt.Println("Error parsing project basics:", err)
					return
				}

				table.Append([]string{directory, project_name, b.Name, b.Label})
			}
		}
		table.Render() // Send output
	},
}

func init() {
	projectsCmd.AddCommand(lsCmd)
}
