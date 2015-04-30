// Copyright 2015 DrWrong
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
// the UI service for the composition system.
// the service is build on the web framework "Macaron"
package main

import (
	// "fmt"
	"flag"
	"github.com/DrWrong/finalProject_UI/controller"
	"github.com/DrWrong/finalProject_UI/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pongo2"
	"os"
	"runtime"
	// "tech_oa/middleware/binding"
)

// return a new Macaron instance
//  the configure to use pongo2 as a template engine
func newMacaron() *macaron.Macaron {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory: "views",
	}))
	return m

}

// the url mapping of the system, it specify the process function of each http request.
func runWeb() {
	m := newMacaron()
	// bindIgnErr := binding.BindIgnErr
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Left"] = "{{"
		ctx.Data["Right"] = "}}"
		ctx.HTML(200, "index")
	})
	m.Get("/plain", controller.Plain)
	m.Get("/encrypted", controller.Encrypted)
	m.Get("/decrypted", controller.Decrypted)
	m.Run()
}

// the entrance of the program, it read a configure file and initialize the configure.
// then it start the web service.
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pwd, _ := os.Getwd()
	exeDir := flag.String("d", pwd, "Execute Directory")
	flag.Parse()
	log.Infof("run server at: %s", *exeDir)
	utils.InitConfig(*exeDir + "/conf/uiserver.cfg")
	runWeb()
}
