package handler

import "testUmbrella/internal/middleware"

func (s *server) InitRoutes() {

	s.echo.Use(middleware.CheckAdmin)
	s.echo.GET("/days", s.getTime)
}
