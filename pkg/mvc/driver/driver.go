package driver

type MvcDriver struct {
	controller Controller
}

func (d *MvcDriver) SetController(controller Controller) {
	d.controller = controller
	d.Initialize()
}

func (d *MvcDriver) Update(delta float64) {
	if d.controller != nil {
		d.controller.Update(delta)
	}
}

func (d *MvcDriver) Initialize() {
	if d.controller != nil {
		d.controller.Initialize()
	}
}

func (d *MvcDriver) Render(delta float64) {
	if d.controller != nil && d.controller.GetView() != nil {
		d.controller.GetView().Render(delta)
	}
}
