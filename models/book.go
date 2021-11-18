package models

type BookManagement struct {
	ID           int    `json:"id" validate:"min=0"`
	Name         string `json:"name" validate:"min=0,max=15"`
	Author       string `json:"author" validate:"min=0,max=15"`
	Prices       int    `json:"prices" validate:"min=0"`
	Available    string `json:"available" validate:"min=0,max=15"`
	PageQuality  string `json:"pagequality" validate:"min=0,max=15"`
	LaunchedYear string `json:"launchedyear" validate:"min=0,max=15"`
	Isbn         string `json:"isbn" validate:"min=0,max=15"`
	Stock        int    `json:"stock" validate:"min=0"`
}
