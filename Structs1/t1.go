package Structs1

import "fmt"

type Human struct {
	Sex  bool
	Age  int
	Name string
}

func (h Human) GetName() {
	fmt.Println(h.Name)
}

type Action struct {
	ActionName string
	Human      // аналог наследования
}

func (a Action) MakeAction() {
	fmt.Println(a.Human.Name + " make " + a.ActionName)
}

func Demo() {
	a := &Action{
		ActionName: "cook",
	}
	a.GetName() // from Human

}
