package window

type Screen interface {
	Initialize()
	Update(delta float64)
	Render(delta float64)
}

