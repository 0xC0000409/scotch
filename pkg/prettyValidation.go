package pkg

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	errorInternalError = errors.New("whoops something went wrong")
)

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func lcFirst(str string) string {
	return strings.ToLower(str)
}

func split(src string) string {
	// don't split invalid utf8
	if !utf8.ValidString(src) {
		return src
	}
	var entries []string
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}

	for index, word := range entries {
		if index == 0 {
			entries[index] = ucFirst(word)
		} else {
			entries[index] = lcFirst(word)
		}
	}
	justString := strings.Join(entries, " ")
	return justString
}

func validationErrorToText(e validator.FieldError) string {
	word := split(e.Field())

	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", word)
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", word, e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", word, e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", word, e.Param())
	case "unique_email":
		return fmt.Sprintf("Email should be unique")
	}
	return fmt.Sprintf("%s is not valid", word)
}

func prettyValidationErrors(c *gin.Context) {
	if len(c.Errors) > 0 {
		for _, e := range c.Errors {
			// Find out what type of error it is
			switch e.Type {
			case gin.ErrorTypePublic:
				// Only output public errors if nothing has been written yet
				if !c.Writer.Written() {
					c.AbortWithStatusJSON(c.Writer.Status(), gin.H{"Error": e.Error()})
				}
			case gin.ErrorTypeBind:
				errs := e.Err.(validator.ValidationErrors)
				list := make(map[string]string)
				for _, err := range errs {
					list[err.Field()] = validationErrorToText(err)
				}

				// Make sure we maintain the preset response status
				status := http.StatusBadRequest
				if c.Writer.Status() != http.StatusOK {
					status = c.Writer.Status()
				}
				c.AbortWithStatusJSON(status, gin.H{"Errors": list})
			}

		}
		// If there was no public or bind error, display default 500 message
		if !c.Writer.Written() {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": errorInternalError.Error()})
		}
	}
}
