package mapper

import (
	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
	"vitaliesvet.com/post-rest-app/rest/view"
)

func MapToPostView(model model.Post, index int) view.PostView {
	return view.PostView{
		View:  view.View{ID: model.ID},
		Text:  model.Text,
		Title: model.Title,
	}
}
