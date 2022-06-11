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

	const (
		reportPath string = "public/report"
		indexPath  string = "public/index.html"
	)

	config.ConnectDB()
	// logger
	s.e.Use(middleware.Logger())
	// recover
	s.e.Use(middleware.Recover())
	// jwt
	// Unauthenticated route
	s.e.POST("/login", loginAuth)
	// Restricted group
	r := s.e.Group("/cashier")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	//Restricted Endpoints
	r.GET("/journal", cashier.GetJournal)
	r.POST("/journalcreate", cashier.CreateJournal)
	r.POST("/cashier/journalcreate", cashier.CreateJournal)
	r.POST("/users/create", users.CreateUser)

	//Public Endpoints
	s.e.GET("/cashier/voucher", cashier.GetVoucher)
	s.e.POST("/cashier/journalsearch", cashier.SearchJournal)
	s.e.POST("/cashier/journalreport", cashier.GetJournalReport)
	s.e.GET("/users/search", users.SearchUser)

	//serve static
	s.e.Static("/report", reportPath)
	s.e.File("/index.html", indexPath)

	s.e.Logger.Fatal(s.e.Start(port))
}

// Close server functionality
func (s *CesGo) Close() {
	s.e.Close()
}
