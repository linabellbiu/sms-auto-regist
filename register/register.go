package register

import "time"

type AppRegister interface {
	register()
}

func Run(apps ...AppRegister) {
	for _, app := range apps {
		go func() {
			tick := time.NewTicker(3 * time.Second)
			app.register()
			<-tick.C
		}()
	}
}
