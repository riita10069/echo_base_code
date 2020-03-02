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
    Update(id int64, form *form.Manual) (*entity.Manual, error)
    Delete(id int64) error
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
  entity := &entity.Manual{
    ID:          form.ID,
    Title:       form.Title,
    Description: form.Description,
    TagID:       form.TagID,
  }

  manual, err := usecase.ManualRepo.Create(entity)
  if err != nil {
    return nil, err
  }
  return manual, nil
}

func (usecase *Manual)Update(id int64, form *form.Manual) (*entity.Manual, error) {
  entity, err := usecase.ManualRepo.ByID(&entity.Manual{}, id)
  if err != nil {
    return nil, err
  }
  entity.ID = form.ID
  entity.Title = form.Title
  entity.Description = form.Description
  entity.TagID = form.TagID

  manual, err := usecase.ManualRepo.Update(entity)
  if err != nil {
    return nil, err
  }
  return manual, nil
}

func (usecase *Manual)Delete(id int64) error {
  entity, err := usecase.ManualRepo.ByID(&entity.Manual{}, id)
  if err != nil {
    return err
  }
  err = usecase.ManualRepo.Delete(entity)
  return err
}
