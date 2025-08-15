package view

type PostView struct {
	View  `json:",inline"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (v PostView) GetID() uint { return v.ID }
