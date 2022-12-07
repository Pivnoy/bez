package usecase

import (
	"context"
	"net/http"
)

type (
	UserRp interface {
	}

	GoogleAPI interface {
		CreateRegLink() string
		CreateClient(context.Context, string) (*http.Client, error)
	}
)
