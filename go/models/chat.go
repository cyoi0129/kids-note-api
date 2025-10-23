package models

type Chat struct {
	Id     uint
	Member []int
}

type ChatData struct {
	Id         uint
	TargetId   int
	TargetName string
	Message    string
	Date       string
}
