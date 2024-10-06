package app

type Application struct {
	model Model
}

func (a *Application) Run() error {
	err := a.eventLoop()
	return err
}
func (a *Application) eventLoop() error {
	for {}
	return nil
}
func (a *Application) render() error {


	return nil
}
