package sayhi

import (
	"fmt"
	"net/http"
)

type SayHi struct {
	Name     string
	Birthday string
	URL      string
}

func (s SayHi) HI() string {
	return fmt.Sprintf("hi %s, your birthday is %s\n", s.Name, s.Birthday)
}

func (s SayHi) GetURL(req *http.Request) string {
	s.URL = fmt.Sprintf("%s", req.URL)
	return fmt.Sprintf("the url is %v\n", s.URL)
}

func (s SayHi) SendRequest(req *http.Request) (string, error) {
	// do a vcenter query request

	return fmt.Sprint("the behavior is work"), nil
}
