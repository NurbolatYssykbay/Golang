package Supervisor

type Supervisor struct {
	position string
	salary   float64
	address  string
}

func NewSupervisor(position string, salary float64, address string) *Supervisor {
	return &Manager{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (m *Supervisor) GetPosition() string {
	return m.position
}

func (m *Supervisor) SetPosition(position string) {
	m.position = position
}

func (m *Supervisor) GetSalary() float64 {
	return m.salary
}

func (m *Supervisor) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Supervisor) GetAddress() string {
	return m.address
}

func (m *Supervisor) SetAddress(address string) {
	m.address = address
}
