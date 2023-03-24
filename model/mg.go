package model

type Record struct {
	Name  string `bson:"name"`
	Title string `bson:"title"`
}
