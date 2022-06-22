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
//
/// Example of parsing different arguments including an array of values

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Command line arguments
type arrayVars []string

func (i *arrayVars) String() string {
	return "String of parameters"
}

func (i *arrayVars) Set(s string) error {
	if strings.Contains(s, "/") {
		//
		//log.Fatalln("The path to the file")
		tagFile, err := os.Open(s)
		if err != nil {
			return err
		}
		defer tagFile.Close()
		scanner := bufio.NewScanner(tagFile)
		for scanner.Scan() {
			*i = append(*i, scanner.Text())
		}
	} else if strings.Contains(s, "-") {
		if strings.Contains(s, ",") {
			//Comma and dash presented as several ranges
			// TO DO

		} else {
			if len(strings.Split(s, "-")) > 2 {
				//exit and warning
				log.Fatalln("Only max and min values for a tags range separated by commas!!!")
			}
			var r_min int
			var r_cnt int
			for _, r := range strings.Split(s, "-") {
				r_str := strings.Split(r, ".")[1]
				r_num, err := strconv.Atoi(r_str)
				if err != nil {
					panic(err)
				}
				r_cnt = r_num - r_min
				r_min = r_num
			}
			//fmt.Println(strconv.Itoa(r_cnt))
			k := strings.Split(strings.Split(s, "-")[0], ".")
			//fmt.Println(k[1])
			for j := 0; j <= r_cnt; j++ {
				//
				m, err := strconv.Atoi(k[1])
				if err != nil {
					panic(err)
				}
				n := strconv.Itoa(m + j)
				l := k[0] + "." + n + "." + k[2]
				//fmt.Println(l)
				*i = append(*i, l)
			}
		}

	} else {
		*i = strings.Split(s, ",")
	}
	//*i = strings.Split(s, "-")
	return nil
}

var (
	isCreate    = flag.Bool("create", false, "Create app manifests")
	isDelete    = flag.Bool("delete", false, "Delete app manifests")
	numCopies   = flag.Int("numcopies", 0, "Number of generated manifests")
	tmplPath    = flag.String("tmplpath", "templates/guestbook-helm.tmpl", "Full path to the template file")
	namePattern = flag.String("namepattern", "guestbook-helm-", "Name pattern for generated manifests")
	destDir     = flag.String("destdir", "./", "Destination directory in the repository")
	imageTags   arrayVars
)

func main() {
	//var imageTag arrayVars
	flag.Var(&imageTags, "imagetags", "Comma separated list of image tags")
	flag.Parse()
	//fmt.Println(len(imageTags))
	//fmt.Println(*tmplPath)
	x := 0
	//
	for i := 1; i < *numCopies; i++ {
		for _, imageTag := range imageTags {
			if x < *numCopies {
				fmt.Println(strconv.Itoa(x) + " " + imageTag)
			}
			x++
		}
	}

	// for _, imageTag := range imageTags {
	// 	//fmt.Println(imageTag)
	// 	for i := 1; i < *numCopies; i++ {
	// 		if x < *numCopies {
	// 			fmt.Println(strconv.Itoa(x) + " " + imageTag)
	// 		}
	// 		x++
	// 	}
	// }

}
