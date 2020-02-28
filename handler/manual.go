package handler

import (
  "github.com/labstack/echo"
  "github.com/riita10069/echo_base_code/form"
  "github.com/riita10069/echo_base_code/response"
  "github.com/riita10069/echo_base_code/usecase"
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

func NewManual() *Manual {
  return &Manual{}
}


func (h *Manual) List(c echo.Context) error {
  // form, err := form.NewPaging(c)
  manuals, err := usecase.NewManual().GetList()
  if err := nil {

  }
  return c.JSON(http.StatusOK, response.NewManuals(manuals))
}

func (h *Manual) Get(c echo.Context) error {
  id, err := form.GetID(c)
  if err != nil {
    return err
  }

  manual, err := usecase.NewManual().GetByID(id)

  return c.JSON(http.StatusOK, response.NewManual(manual))
}

func (h *Manual) Create(c echo.Context) error {
  form, err := form.NewManual(c)
  if err != nil {
    return err
  }

  manual, err := usecase.NewManual().Create(form)

  return c.JSON(http.StatusCreated, response.NewManual(manual))
}

func (h *Manual) Update(c echo.Context) error{
  return c.JSON(http.StatusOK, Manual{})
}

func (h *Manual) Destroy(c echo.Context) error{
  return c.JSON(http.StatusOK, Manual{})
}

func (h *Manual) UploadMainImage(c echo.Context) error{
  return c.JSON(http.StatusOK, Manual{})
}

