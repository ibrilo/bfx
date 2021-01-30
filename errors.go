package bfx

import "fmt"

// RestError represetnt struct for RESP API errors
type RESTError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func restError(d abstractSlice) *RESTError {
	e, err := parseRESTError(d)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return e
}

func parseRESTError(d []interface{}) (*RESTError, error) {
	if len(d) != 3 {
		return nil, fmt.Errorf("failed to parse REST error")
	}
	return &RESTError{parseNumber(d[1]).Int(), parseString(d[2]).String()}, nil
}

func (re *RESTError) print() {
	fmt.Println("API ERROR:", re.String())
}

func (re *RESTError) String() string {
	return fmt.Sprintf("<(%d) %s>", re.Code, re.Message)
}
