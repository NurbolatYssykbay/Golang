package Bookkeeper

type Bookkeeper struct {
	position string
	salary   float64
	address  string
}

func NewBookkeeper(position string, salary float64, address string) *Bookkeeper {
	return &Bookkeeper{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (m *Bookkeeper) GetPosition() string {
	return m.position
}

func (m *Bookkeeper) SetPosition(position string) {
	m.position = position
}

func (m *Bookkeeper) GetSalary() float64 {
	return m.salary
}

func (m *Bookkeeper) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Bookkeeper) GetAddress() string {
	return m.address
}

func (m *Bookkeeper) SetAddress(address string) {
	m.address = address
}
