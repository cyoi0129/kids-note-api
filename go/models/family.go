package models

type Family struct {
	Id   uint
	Name string
}

type FamilyMember struct {
	Id   int
	Name string
}

type FamilyResponse struct {
	Id      uint
	Name    string
	Members []FamilyMember
}
