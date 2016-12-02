package main

import (
	"github.com/kkserver/kk-cache/cache"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.Llongfile | log.LstdFlags)

	env := "./config/env.ini"

	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	a := cache.CacheApp{}

	err := app.Load(&a, "./app.ini")

	if err != nil {
		log.Panicln(err)
	}

	err := app.Load(&a, env)

	if err != nil {
		log.Panicln(err)
	}

	app.Obtain(&a)

	app.Handle(&a, &app.InitTask{})

	kk.DispatchMain()

	app.Recycle(&a)

}
