package bfx

import (
	"errors"
	"log"
	"net/http"
)

const retriesRESTLimit = 5

var (
	ErrNotEnoughParameters  = errors.New("not enough parameters for auth")
	ErrWrongArgumentsPassed = errors.New("wrong argument type was passed")
)

type Client struct {
	logger log.Logger

	Debug bool

	Auth clientAuth

	RetriesLimit int

	httpClient *http.Client
}

func New(args ...interface{}) (*Client, error) {

	c := &Client{
		httpClient:   &http.Client{},
		RetriesLimit: retriesRESTLimit,
	}

	if args == nil {
		return c, nil
	}

	for _, arg := range args {

		switch v := arg.(type) {

		case []string:
			if len(v) != 2 {
				return nil, ErrNotEnoughParameters
			}

			c.Auth.key = v[0]
			c.Auth.secret = v[1]

		default:
			return nil, ErrWrongArgumentsPassed
		}
	}

	return c, nil
}

type clientAuth struct {
	key    string
	secret string
}

func (ca *clientAuth) Key() string {
	return ca.key
}

func (ca *clientAuth) Secret() string {
	return ca.secret
}
