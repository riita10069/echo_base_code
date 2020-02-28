package usecase

import (
  "github.com/riita10069/echo_base_code/entity"
  "github.com/riita10069/echo_base_code/form"
  "github.com/riita10069/echo_base_code/repository"
)

type (
  IManual interface {
    GetList() (*entity.Manuals, error)
    Create(form *form.Manual) (*entity.Manual, error)
    GetByID(id int64) (*entity.Manual, error)
  }

  Manual struct {
    ManualRepo repository.IManual
  }
)

func NewManual() IManual {
  return &Manual{
    ManualRepo: repository.NewManual(),
  }
}

func (usecase *Manual)GetList() (*entity.Manuals, error) {
  entities, err := usecase.ManualRepo.All(&entity.Manuals{})
  if err != nil {
    return nil, err
  }
  return entities, nil
}

func (usecase *Manual)GetByID(id int64) (*entity.Manual, error) {
  entity, err := usecase.ManualRepo.ByID(&entity.Manual{}, id)
  if err != nil {
    return nil, err
  }
  return entity, nil
}

func (usecase *Manual)Create(form *form.Manual)(*entity.Manual, error){
  manual := &entity.Manual{
    ID:          form.ID,
    Title:       form.Title,
    Description: form.Description,
    TagID:       form.TagID,
  }

  manual, err := usecase.ManualRepo.Create(manual)
  if err != nil {
    return nil, err
  }
  return manual, nil
}
