package persistence

import "github.com/saidaydogan/chi-poc/pkg/errors"

var (
	NotFoundError = errors.Error("1", "Not found get by id")
)
