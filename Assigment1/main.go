package main

import (
	"awesomeProject1/Assistant"
	"awesomeProject1/Accountant"
	"awesomeProject1/Creator"
	"awesomeProject1/Mechanic"
	"awesomeProject1/Supervisor"
	"fmt"
)

func main() {
	supervisor := Supervisor.NewSupervisor("Supervisor", 260000, "Gogolya, 32")

	creator := Creator.NewCreator("Creator", 370000, "AbylayKhan, 140")

	accountant := Accountant.NewAccountant("Accountant", 280000, "Raiymbekova, 42")

	mechanic := Mechanic.NewMechanic("Mechanic", 300000, "Penzenskaya, 27")

	assistant := Assistant.NewAssistant("Assistant", 170000, "Sarsenbaeva, 21")

	fmt.Println("Supervisor Position:", supervizor.GetPosition())
	supervisor.SetPosition("Senior Supervisor")
	fmt.Println("Supervisor Position (Updated):", supervisor.GetPosition())

	fmt.Println("Creator Salary:", creator.GetSalary())
	creator.SetSalary(385000)
	fmt.Println("Creator Salary (Updated):", creator.GetSalary())

	fmt.Println("Accountant Position:", accountant.GetPosition())

	fmt.Println("Mechanic Position:", mechanic.GetPosition())

	fmt.Println("Assistant Address:", assistant.GetAddress())

}
