package impl

import (
	"github.com/labstack/echo/v4"

	"github.com/goforbroke1006/go-msvc-tpl/internal/component/api/internal_admin/spec"
)

func NewHandlers() spec.ServerInterface {
	return &handlersImpl{}
}

var _ spec.ServerInterface = (*handlersImpl)(nil)

type handlersImpl struct{}

func (h handlersImpl) PostGenerateReport(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
