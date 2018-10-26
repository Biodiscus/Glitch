package driver

type Controller interface {
	GetView() View
	Initialize()
	Update(delta float64)
}
