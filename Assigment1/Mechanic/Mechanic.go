package Mechanic

type Mechanic struct {
	position string
	salary   float64
	address  string
}

func NewMechanic(position string, salary float64, address string) *Mechanic {
	return &Engineer{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (d *Mechanic) GetPosition() string {
	return d.position
}

func (d *Mechanic) SetPosition(position string) {
	d.position = position
}

func (d *Mechanic) GetSalary() float64 {
	return d.salary
}

func (d *Mechanic) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Mechanic) GetAddress() string {
	return d.address
}

func (d *Mechanic) SetAddress(address string) {
	d.address = address
}
