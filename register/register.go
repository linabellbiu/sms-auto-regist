package register

type AppRegister interface {
	Register()
}

func Run(apps ...AppRegister) {
	for _, app := range apps {
		go func() {
			app.Register()
		}()
	}
}
