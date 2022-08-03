package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов
в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	// ....
	Name    string
	Surname string
}

func (h *Human) sayNameSurname() {
	fmt.Printf("My name is %s and surname is %s\n", h.Name, h.Surname)
}

type Action struct {
	Human
	// В структуре Action доступны все поля и методы структуры Human
	Sex bool
}

func (a *Action) makeAction() {
	// используем метод встроенной структуры Human
	a.sayNameSurname()
	switch a.Sex {
	case false:
		fmt.Print("Hi girl!\n")
	case true:
		fmt.Print("Hi boy!\n")
	}
}

func main() {
	a := Action{
		Human: Human{
			Name:    "Nikita",
			Surname: "Ponomarev",
		},
		Sex: true,
	}
	a.makeAction()
	// доступен метод встроенного типа
	a.sayNameSurname()
	// также к методом встроенного типа можно обращаться и так
	a.Human.sayNameSurname()
}

/*
Вышележащий тип может переопределить метод.
Причем метод перекрывается по имени, без учета полной сигнатуры метода:

1

В языке Go нет наследования — есть только композиция (встраивание):

type X struct {
    Value int
}

func (x X) GetValue() int { return x.Value }

type Y struct {
    X
}

type Z struct {
    X X
}
При этом обратите внимание, что встраивание в структурах Y и Z разное — в
первом случае оно "безымянное", благодаря чему можно вызывать методы X прямо
из структуры Y (пример: Y{}.GetValue() вернёт 0 как значение int по умолчанию),
 а во втором случае создаётся поле X типа X, из-за чего вызывать методы нельзя
 (пример: Z{}.GetValue() вызовет ошибку при компиляции).
*/
