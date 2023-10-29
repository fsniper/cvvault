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
	"bytes"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/metakeule/scaffold/lib/scaffold"
)

var (
	fullName    string
	label       string
	image       string
	email       string
	phone       string
	url         string
	summary     string
	address     string
	postcode    string
	city        string
	countrycode string
	region      string
)

// createCmd represents the add command
var createCmd = &cobra.Command{
	Use:   "create [project name]",
	Short: "Create a new Cv Vault Project",
	Long:  `This will create a new CV Vault project under the projects directory.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		template := `
>>>projects/
>>>{{.Directory}}/
>>>data/
>>>basics.json
{
  "name": "{{.FullName}}",
  "label": "{{.Label}}",
  "image": "{{.Image}}",
  "email": "{{.Email}}",
  "phone": "{{.Phone}}",
  "url": "{{.Url}}",
  "summary": "{{.Summary}}",
  "location": {
    "address": "{{.Address}}",
    "postalCode": "{{.Postcode}}",
    "city": "{{.City}}",
    "countryCode": "{{.CountryCode}}",
    "region": "{{.Region}}"
  }
}
<<<basics.json
>>>works/
>>>my-first-company-STARTYEAR.json
    {
      "name": "My First Company",
      "location": "",
      "description": "",
      "position": "",
      "url": "",
      "startDate": "",
      "endDate": "",
      "summary": "",
      "highlights": [
				{ "description": "", "labels": ["",""] },
      ]
    }
<<<my-first-company-STARTYEAR.json
>>>exports/
<<<exports/
<<<works/
<<<data/
<<<{{.Directory}}/
<<<projects/
`
		fields := map[string]string{
			"Directory":   args[0],
			"FullName":    fullName,
			"Label":       label,
			"Image":       image,
			"Email":       email,
			"Phone":       phone,
			"Url":         url,
			"Summary":     summary,
			"Address":     address,
			"Postcode":    postcode,
			"City":        city,
			"CountryCode": countrycode,
			"Region":      region,
		}

		ioreader := new(bytes.Buffer)
		json.NewEncoder(ioreader).Encode(fields)

		scaffold.Run(".", template, ioreader, os.Stdout, false)
	},
}

func init() {
	projectsCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&fullName, "fullname", "n", "", "Full Name of the CV owner (required)")
	createCmd.MarkFlagRequired("fullname")
	createCmd.Flags().StringVarP(&label, "label", "l", "", "Label of the CV owner - e.g. \"Support Engineer\" (required)")
	createCmd.MarkFlagRequired("label")

	createCmd.Flags().StringVarP(&image, "image", "i", "", "image path of the CV owner")
	createCmd.Flags().StringVarP(&email, "email", "e", "", "email of the CV owner")
	createCmd.Flags().StringVarP(&phone, "phone", "p", "", "phone of the CV owner")
	createCmd.Flags().StringVarP(&url, "url", "u", "", "url of the CV owner")
	createCmd.Flags().StringVarP(&summary, "summary", "s", "", "summary of the CV owner")
	createCmd.Flags().StringVarP(&address, "address", "a", "", "address of the CV owner")
	createCmd.Flags().StringVarP(&postcode, "postcode", "", "", "postcode of the CV owner")
	createCmd.Flags().StringVarP(&city, "city", "", "", "city of the CV owner")
	createCmd.Flags().StringVarP(&countrycode, "countrycode", "", "", "countrycode of the CV owner")
	createCmd.Flags().StringVarP(&region, "region", "r", "", "region of the CV owner")
}
