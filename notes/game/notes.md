// config and resources all in same language you do your code in. they have object schemas, config
// todo: ci/cd steps?

package main

func main() {
	if err := config.Register(); err != nil {
		log.Fatal(err)
	}

	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(conf)
}


// I think we need to laod context first
