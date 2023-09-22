package Accountant

type Accountant struct {
	position string
	salary   float64
	address  string
}

func NewAccountant(position string, salary float64, address string) *Accountant {
	return &Accountant{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (m *Accountant) GetPosition() string {
	return m.position
}

func (m *Accountant) SetPosition(position string) {
	m.position = position
}

func (m *Accountant) GetSalary() float64 {
	return m.salary
}

func (m *Accountant) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Accountant) GetAddress() string {
	return m.address
}

func (m *Accountant) SetAddress(address string) {
	m.address = address
}
