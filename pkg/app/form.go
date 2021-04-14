package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key string
	Value string
}

type ValidErrors []*ValidError

func (vs ValidErrors) Errors() []string {
	var errs []string
	for _, err := range vs {
		errs = append(errs, err.Value)
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (ValidErrors, bool) {
	var errs ValidErrors
	err := c.ShouldBind(&v)
	if err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return errs, false
		}
		trans := c.Value("trans").(ut.Translator)
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:   key,
				Value: value,
			})
		}
		return errs, false
	}

	return errs, true
}
