package model

//TODO: add ids for each struct

type Editor struct {
	Email    string
	Password string
}

type Subscriber struct {
	Email      string
	Newsletter []Newsletter
}

type Newsletter struct {
	Name        string
	Description *string
	Editor      Editor
	Subscriber  []Subscriber
}

type Post struct {
	Newsletter Newsletter
	Text       string
}
