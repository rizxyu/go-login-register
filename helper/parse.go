package helper

import (
	"net/http"

	"github.com/gorilla/schema"
)

func ParseForm(r *http.Request, dll interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(dll, r.PostForm); err != nil {
		return err
	}
	return nil
}
