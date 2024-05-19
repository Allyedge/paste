package handlers

import "github.com/Allyedge/paste/db/repositories"

type Handlers struct {
	HomeHandler *HomeHandler
	CodeHandler *CodeHandler
}

func NewHandlers(r *repositories.Repository) *Handlers {
	return &Handlers{
		HomeHandler: NewHomeHandler(r),
		CodeHandler: NewCodeHandler(r),
	}
}
