package form

import (
  "github.com/labstack/echo"
  "strconv"
)

type (
  Manual struct {
    ID           int64       `json:"id"`
    Title        string      `json:"title" validate:"required"`
    Description  string      `json:"description"`
    TagID        int64       `json:"tag_id"`
  }

  ManualDescription struct {
    Description string `json:"description"`
  }
)

func InitManual() *Manual {
  return &Manual{}
}

func NewManual(c echo.Context) (*Manual, error){
  form := InitManual()
  if err := c.Bind(form); err != nil {
    return nil, err
  }

  if err := c.Validate(form); err != nil {
    return nil, err
  }

  return form, nil
}

func GetID(c echo.Context) (id int64, err error){
  id, err = strconv.ParseInt(c.Param("id"), 10, 64)
  return id, err
}

