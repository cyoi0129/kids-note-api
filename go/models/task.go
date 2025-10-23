package models

type TaskType struct {
	Id     uint
	Name   string
	Family int
}

type Task struct {
	Id     uint
	Name   string
	Detail string
	Types  []int
	Status string
	Update string
	Due    string
	Items  []int
	Kid    int
	UserId int
	Family int
}
