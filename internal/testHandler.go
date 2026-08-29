package internal

import (
	"net/http"
)

type TestHandler struct{
	
}

func NewTestHandle () *TestHandler{
	return &TestHandler{}


}



func (t *TestHandler) ServeHTTP (w http.ResponseWriter,r *http.Request) {

	w.Write([]byte("Hello"))





} 



