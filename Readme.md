# CV Vault

This is a cli application to manage cvs in (mostly) json-resume compatible yaml files.
Some sections are extended with tags so you can create different cvs by tags. (check export command)

```
name: "Parkyeri"
location: "Istanbul / Turkey"
description: "Parkyeri is a Istanbul based Information Technology & Services Company"
position: "Linux Systems Administrator"
url: "https://www.parkyeri.com"
startDate: "2005-02-01"
endDate: "2007-02-01"
summary: ""
highlights:
- description: "Linux Systems Administration" 
  tags: ["sre","all"]
- description: "Postfix Mail Server Administration" 
  tags: ["sre","all"]
- description: "Mysql Management" 
  tags: ["sre","all"]
- description: "Postgresql Management" 
  tags: ["sre","all"]
- description: "GSM Value Added Services Development"
  tags: ["swe", "gsm"]
- description: "Kannel Management"
  tags: ["sre", "gsm"]
```

## Commands

### cvvault project
Available Commands:
  create      Create a new Cv Vault Project
  delete      Delete project
  export      Export project as json
  ls          List projects
  print       Print project
  set-default Set a project as default

```
% go run cvvault.go projects ls
INFO[0000] Config file in use: /Users/yalazi/.config/cvvault/cvvault-config.yaml
INFO[0000] Reading basics for project: yalazi
INFO[0000] Validating json
INFO[0000] Reading basics for project: yzl
INFO[0000] Validating json
  DEFAULT |                      DIRECTORY                       | PROJECT |        NAME        |       LABEL
----------+------------------------------------------------------+---------+--------------------+--------------------
          | /Users/yalazi/Documents/CVVault/WORK/projects/yalazi | yalazi  | Mehmet Onur Yalazi | Platform Engineer
          | /Users/yalazi/Documents/CVVault/WORK/projects/yzl    | yzl     | Mehmet Onur Yalazi | PLatform Engineer
```

```
% go run cvvault.go projects create onur -n "Mehmet Onur Yalazi" -l "PLatform Engineer" -a "7 Oak Park View" --city Naas --countrycode IE -e onur.yalazi@gmail.com -r Co.Kildare
INFO[0000] Config file in use: /Users/yalazi/.config/cvvault/cvvault-config.yaml
INFO[0000] Creating CVProject:  onur /Users/yalazi/Documents/CVVault/WORK/projects/onur
/Users/yalazi/Documents/CVVault/WORK/projects/onur/data/basics.yaml
/Users/yalazi/Documents/CVVault/WORK/projects/onur/data/works/my-first-company-STARTYEAR.yaml
yalazi@Onurs-MacBook-Pro cvvault % go run cvvault.go projects create onur -n "Mehmet Onur Yalazi" -l "PLatform Engineer" -a "7 Oak Park View" --city Naas --countrycode IE -e onur.yalazi@gmail.com -r Co.Kildare
INFO[0000] Config file in use: /Users/yalazi/.config/cvvault/cvvault-config.yaml
INFO[0000] Creating CVProject:  onur /Users/yalazi/Documents/CVVault/WORK/projects/onur
/Users/yalazi/Documents/CVVault/WORK/projects/onur/data/basics.yaml
/Users/yalazi/Documents/CVVault/WORK/projects/onur/data/works/my-first-company-STARTYEAR.yaml
```

```
% go run cvvault.go projects delete onur
INFO[0000] Config file in use: /Users/yalazi/.config/cvvault/cvvault-config.yaml
Are you sure to delete this project: onur? [y/n]: y
Deleting project:  onur
```

```
% go run cvvault.go projects export yalazi --ignoreTags gsm
INFO[0000] Config file in use: /Users/yalazi/.config/cvvault/cvvault-config.yaml
INFO[0000] Reading basics for project: yalazi
INFO[0000] Validating json
{"Name":"yalazi","basics":{"name":"Mehmet Onur Yalazi","label":"Platform Engineer","image":"","email":"onur.yalazi@gmail.com",
....
{"name":"Parkyeri","location":"Istanbul / Turkey","description":"Parkyeri is a Istanbul based Information Technology Services Company",
"position":"Linux Systems Administrator","url":"https://www.parkyeri.com","startDate":"2005-02-01","endDate":"2007-02-01","summary":"",
"highlights":["Linux Systems Administration","Postfix Mail Server Administration","Mysql Management","Postgresql Management"]}
...
```

