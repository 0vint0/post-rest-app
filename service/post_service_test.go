package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	model "vitaliesvet.com/post-rest-app/persistent/db/model/post"
	"vitaliesvet.com/post-rest-app/persistent/repository/mocks"
	"vitaliesvet.com/post-rest-app/rest/view"
	"vitaliesvet.com/post-rest-app/service"
)

func TestPostService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	postRepositoryMock := mocks.NewMockPostRepository(ctrl)

	toTest := service.NewPostService(postRepositoryMock)

	t.Run("Test find all paginated Post result.", func(t *testing.T) {
		postRepositoryMock.EXPECT().FindAll(int64(0), int64(10)).Return(expectedPostsInDd(), int64(2), nil)
		expectedResult := expectedPostViews()
		pagedRequest := view.PagedRequest{
			PageNumber: 0,
			PageSize:   10,
		}
		actualResult, error := toTest.FindAll(pagedRequest)

		assert.Nil(t, error)
		assert.Equal(t, int64(2), actualResult.TotalSize, "Invalid FindAll Post TotalSize Result !!!")
		assert.Equal(t, expectedResult, actualResult.Items, "Invalid FindAll Post list Result !!!")
	})

}

func expectedPostsInDd() []model.Post {
	return []model.Post{
		model.Post{
			Title: "Post1",
			Text:  "Text1",
		},
		model.Post{
			Title: "Post2",
			Text:  "Text2",
		},
	}
}

func expectedPostViews() []view.PostView {
	return []view.PostView{
		view.PostView{
			Title: "Post1",
			Text:  "Text1",
		},
		view.PostView{
			Title: "Post2",
			Text:  "Text2",
		},
	}
}
