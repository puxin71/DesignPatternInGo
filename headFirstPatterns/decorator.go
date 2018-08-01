package headFirstPatterns

type component struct {
	description string
	cost        float64
}

type Component interface {
	Description() string
	Cost() float64
}

func NewComponent(description string, cost float64) Component {
	return &component{description: description, cost: cost}
}
func (c *component) Description() string {
	return c.description
}
func (c *component) Cost() float64 {
	return c.cost
}

type whipDecorator struct {
	wrappedObj Component
}

type soyDecorator struct {
	wrappedObj Component
}

type Decorator interface {
	Description() string
	Cost() float64
}

func NewWhipDecorator(c Component) Decorator {
	return &whipDecorator{wrappedObj: c}
}

func NewSoyDecorator(c Component) Decorator {
	return &soyDecorator{wrappedObj: c}
}

func (wd *whipDecorator) Description() string {
	return wd.wrappedObj.Description() + ", whip"
}
func (wd *whipDecorator) Cost() float64 {
	return wd.wrappedObj.Cost() + float64(0.2)
}

func (sd *soyDecorator) Description() string {
	return sd.wrappedObj.Description() + ", soy"
}
func (sd *soyDecorator) Cost() float64 {
	return sd.wrappedObj.Cost() + float64(0.5)
}
