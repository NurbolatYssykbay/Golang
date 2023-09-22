package Assistant

type Assistant struct {
	position string
	salary   float64
	address  string
}

func NewAssistant(position string, salary float64, address string) *Assistant {
	return &Assistant{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (d *Assistant) GetPosition() string {
	return d.position
}

func (d *Assistant) SetPosition(position string) {
	d.position = position
}

func (d *Assistant) GetSalary() float64 {
	return d.salary
}

func (d *Assistant) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Assistant) GetAddress() string {
	return d.address
}

func (d *Assistant) SetAddress(address string) {
	d.address = address
}
