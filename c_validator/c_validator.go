package c_validator

import (
	"github.com/go-playground/validator/v10"
	"vitaliesvet.com/post-rest-app/persistent/repository"
)

func RegisterCustomValidators(validate *validator.Validate, r repository.PostRepository) {
	validate.RegisterValidation("is-unique-post-title", func(fl validator.FieldLevel) bool {
		title := fl.Field().String()
		count, _ := r.CountByTitle(title)
		return count == 0
	})

}
