package handler

import (
  "github.com/labstack/echo"
  "net/http"
)

type (
  IManual interface {
    List(c echo.Context) error
    Get(c echo.Context) error
    Create(c echo.Context) error
    Update(c echo.Context) error
    Destroy(c echo.Context) error
  }

  Manual struct {

  }
)


func (h *Manual) List(c echo.Context) error {
  return c.JSON(http.StatusCreated, Manual{})
}

func (h *Manual) Get(c echo.Context) error {
  return c.JSON(http.StatusCreated, Manual{})
}

func (h *Manual) Create(c echo.Context) error {
  return c.JSON(http.StatusCreated, Manual{})
}

func (h *Manual) Update(c echo.Context) error{
  return c.JSON(http.StatusCreated, Manual{})
}

func (h *Manual) Destroy(c echo.Context) error{
  return c.JSON(http.StatusCreated, Manual{})
}


