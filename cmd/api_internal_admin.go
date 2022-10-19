package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"

	"github.com/goforbroke1006/go-msvc-tpl/internal/component/api/internal_admin/impl"
	"github.com/goforbroke1006/go-msvc-tpl/internal/component/api/internal_admin/spec"
)

func NewAPIInternalAdminCmd() *cobra.Command {
	return &cobra.Command{
		Use: "internal-admin",
		Run: func(cmd *cobra.Command, args []string) {
			router := echo.New()
			spec.RegisterHandlers(router, impl.NewHandlers())
		},
	}
}

func init() {
	apiCmd.AddCommand(NewAPIInternalAdminCmd())
}
