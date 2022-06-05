package server

import (
	"cesgo/config"
	cashier "cesgo/controllers/cashier"
	users "cesgo/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CesGo structur
type CesGo struct {
	e *echo.Echo
}

// NewServer Instance of Echo
func NewServer() *CesGo {

	return &CesGo{
		e: echo.New(),
	}
}

// Start server functionality
func (s *CesGo) Start(port string) {

	config.ConnectDB()
	// logger
	s.e.Use(middleware.Logger())
	// recover
	s.e.Use(middleware.Recover())

	// Start Server
	s.e.GET("/cashier/voucher", cashier.GetVoucher)
	s.e.GET("/cashier/journal", cashier.GetJournal)
	s.e.GET("/users/search", users.SearchUser)
	s.e.POST("/users/create", users.CreateUser)
	s.e.Logger.Fatal(s.e.Start(port))
}

// Close server functionality
func (s *CesGo) Close() {
	s.e.Close()
}
