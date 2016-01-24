package main

//test comment for git stuff

import (
	"fmt"
	"strconv"
)

type Factory struct {
	Name               string
	Production         float64
	ProductionModifier float64
}
type PlayField struct {
	Money               float64
	Round               int
	GlobalModifier      float64
	SmallFactoryCounter int
	FactoryContainer
}

//Container of Factory structs. this is a slice of []Factory available @ [i]
type FactoryContainer struct {
	Factories []*Factory
}

//Create a new custom factory not used by default --- testing purposes
//func NewFactory(name string, production float64, prodmod float64) *Factory {
//	f1 := new(Factory)
//	f1.Name = name
//	f1.Production = production
//	f1.ProductionModifier = prodmod
//	return f1
//}

//Create a new small factory with default settings
func (p *PlayField) NewSmallFactory() {
	p.SmallFactoryCounter++
	f1 := new(Factory)
	f1.Name = "Small Factory " + strconv.Itoa(p.SmallFactoryCounter)
	f1.Production = 1
	f1.ProductionModifier = 0.2
	p.Money = p.Money - 10
	p.AddFactory(f1)
}
func (p *PlayField) UpgradeFactories() {
	if p.Money >= 100 {

		for i, _ := range p.Factories {
			p.Factories[i].ProductionModifier = p.Factories[i].ProductionModifier + 0.10
		}

		p.Money = p.Money - 100
	} else {
		fmt.Printf("You only have: %f and a to upgrade factories you will need 100\n", p.Money)
		p.NextRound()
	}
}

//function to add a created factory to a factorycontainer slice
func (o *FactoryContainer) AddFactory(f *Factory) {
	o.Factories = append(o.Factories, f)
}

//Counts the output of factories in a FactoryContainer
func (o *FactoryContainer) CountOutput() float64 {
	var output float64 = 0.0
	for i, _ := range o.Factories {
		output = output + o.Factories[i].Production*o.Factories[i].ProductionModifier
	}
	return output
}

//List all factories and stats in a FactoryContainer
func (o *FactoryContainer) ListFactory() {
	for i, _ := range o.Factories {
		fmt.Printf("____________________________________\n")
		fmt.Printf("Factoryname: %s\nProduction: %f\nModifier:%f\n", o.Factories[i].Name, o.Factories[i].Production, o.Factories[i].ProductionModifier)
		fmt.Printf("____________________________________\n")
	}
	output := o.CountOutput()
	fmt.Printf("Total Output from factories: %f\n", output)
	fmt.Printf("____________________________________\n")
}

func (p *PlayField) ListAll() {
	fmt.Printf("Money: %f\nRound: %d\n", p.Money, p.Round)
}
func (p *PlayField) IncreaseRound() {
	p.Round++
}

func (p *PlayField) Menu() {
	var choice string
	fmt.Printf("What is your next move ?\n1)Build a factory\n2)wait turn\n3)Upgrade salary modifier - Costs 50\n4)Upgrade Factory modifier - Costs 100\n")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		p.NewSmallFactory()
		fmt.Println("Added a Small factory")
		p.NextRound()

	case "2":
		p.NextRound()

	case "3":
		p.Money = p.Money - 50
		p.GlobalModifier = p.GlobalModifier + 0.20
		p.NextRound()
	case "4":
		p.UpgradeFactories()
		p.NextRound()
	}
}

//NextRound should do everything needed to be done when on a round
func (p *PlayField) NextRound() {

	salary := p.CountOutput() * p.GlobalModifier
	p.Money = p.Money + salary
	fmt.Printf("Income this round: %f\n", salary)
	p.Round++
	p.ListFactory()
	p.ListAll()
	p.Menu()

}

// experimenting for now
func main() {
	p := PlayField{Money: 100, Round: 1, GlobalModifier: 1.0}
	p.Menu()

}

//TODO create somekind of struct that holds factory types that when selected in buying menu uses Newfactory to create selected type. Maybe even have a json file with factorytypes.
//TODO Add checks for money before building factory or upgrading salarymodifier
//Clear upp code maybe make several files ?
