package main

import (
	// "fmt"
	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pongo2"
	"os"
	"runtime"
	"ui/controller"
	"ui/utils"
	// "tech_oa/middleware/binding"
)

func newMacaron() *macaron.Macaron {
	m := macaron.Classic()
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory: "views",
	}))
	return m

}

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

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pwd, _ := os.Getwd()
	exeDir := flag.String("d", pwd, "Execute Directory")
	flag.Parse()
	log.Infof("run server at: %s", *exeDir)
	utils.InitConfig(*exeDir + "/conf/uiserver.cfg")
	runWeb()
}
