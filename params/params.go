package params

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

// Type for http post body
type PostBody []byte // Byte array with request body

// Get params stands for "query params"
type GetParams map[string]string

type Params struct {
	Query      GetParams
	Body       PostBody
	PathParams PathParams
}

func NewParams(request *http.Request, pattern *regexp.Regexp, path string) (*Params, error) {
	var body []byte
	if request.Body != nil {
		newBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		body = newBody
	}

	return &Params{
		Query:      ValuesToGetParams(request.URL.Query()),
		Body:       body,
		PathParams: ExtractPathParams(pattern, path),
	}, nil
}

func ExtractParams(req *http.Request) *Params {
	return req.Context().Value("params").(*Params)
}

// Converts url.Url.Query() from "Values" (map[string][]string)
// to "getParams" (map[string]string)
func ValuesToGetParams(values url.Values) GetParams {
	params := make(map[string]string)
	for key := range values {
		params[key] = values.Get(key)
	}
	return params
}

// Example: url "/api/v1/users/599a49bacdf43b817eeea57b" and pattern `/api/v1/users/:id`
// path params = {"id": "599a49bacdf43b817eeea57b"}
type PathParams map[string]string

// Extract path params from path
func ExtractPathParams(pattern *regexp.Regexp, path string) PathParams {
	match := pattern.FindStringSubmatch(path)
	result := make(PathParams)

	for i, name := range pattern.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	return result
}
