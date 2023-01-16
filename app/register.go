package app

import "time"

type Register interface {
	Register()
}

func Run(apps ...Register) {
	for _, app := range apps {
		go func(app Register) {
			tick := time.NewTicker(3 * time.Second)
			app.Register()
			<-tick.C
		}(app)
	}
}
