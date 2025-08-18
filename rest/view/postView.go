package view

type PostView struct {
	View  `json:",inline"`
	Title string `json:"title" validate:"required,min=5,is-unique-post-title"`
	Text  string `json:"text" validate:"required"`
}

func (v PostView) GetID() uint { return v.ID }
