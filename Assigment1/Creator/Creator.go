package Creator

type Creator struct {
	position string
	salary   float64
	address  string
}

func NewCreator(position string, salary float64, address string) *Creator {
	return &Creator{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (d *Creator) GetPosition() string {
	return d.position
}

func (d *Creator) SetPosition(position string) {
	d.position = position
}

func (d *Creator) GetSalary() float64 {
	return d.salary
}

func (d *Creator) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Creator) GetAddress() string {
	return d.address
}

func (d *Creator) SetAddress(address string) {
	d.address = address
}
