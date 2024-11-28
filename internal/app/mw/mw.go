package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(roleAdmin, val) {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
}
