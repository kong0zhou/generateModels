package models

type Test struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Score int    `json:"score"`
}
