// Copyright 2022 Gleb Otochkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/// Package contains code to multiply files based on a template
/// The example is created based on an ArgoCD application template

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

// Command line arguments
var (
	isCreate    = flag.Bool("create", false, "Create copies")
	isDelete    = flag.Bool("delete", false, "Delete copies")
	numCopies   = flag.Int("numcopies", 0, "Number of generated copies")
	tmplPath    = flag.String("tmplpath", "templates/guestbook-helm.tmpl", "Full path to the template file")
	namePattern = flag.String("namepattern", "guestbook-helm-", "Name pattern for generated manifests")
	destDir     = flag.String("destdir", "../argocd/argo-apps", "Destination directory in the repository")
)

func main() {

	flag.Parse()
	if *isCreate && *isDelete {
		fmt.Println("Error: you cannot specify --create and --delete simultaneously")
	} else if *isCreate {
		fmt.Println("Creating")
	} else if *isDelete {
		fmt.Println("Deleting")
	} else {
		fmt.Println("Please specify --create or --delete ")
	}
	// Template variables
	tmplvars := make(map[string]interface{})
	tmplvars["AppName"] = *namePattern + "1"
	tmplvars["ServicePort"] = "8001"
	appName := *namePattern
	numcopies := *numCopies + 1

	if *isCreate {
		if _, err := os.Stat(*tmplPath); err == nil {
			fmt.Println(*tmplPath)
			//Parse the template
			tmpl, _ := template.ParseFiles(*tmplPath)
			if *numCopies > 0 {
				fmt.Println(strconv.Itoa(*numCopies))
				for i := 1; i < numcopies; i++ {
					tmplvars["AppName"] = appName + strconv.Itoa(i)
					tmplvars["ServicePort"] = "8" + fmt.Sprintf("%03d", i)
					file, _ := os.Create(*destDir + "/" + appName + strconv.Itoa(i) + ".yaml")
					defer file.Close()
					err := tmpl.Execute(file, tmplvars)
					if err != nil {
						log.Fatalln(err)
					}
				}
			}
		} else if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The template " + *tmplPath + " doesn't exist!")
		}

	}
	if *isDelete {
		files, err := filepath.Glob(*destDir + "/" + "guestbook-helm*.yaml")
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			if err := os.Remove(f); err != nil {
				panic(err)
			}
		}
	}

}
