package repository

import (
  "github.com/riita10069/echo_base_code/db"
  "github.com/riita10069/echo_base_code/entity"
)

type (
  IManual interface {
    All(entity *entity.Manuals) (*entity.Manuals, error)
    ByID(entity *entity.Manual, id int64) (*entity.Manual, error)
    Create(entity *entity.Manual) (*entity.Manual, error)
    Update(e *entity.Manual) (*entity.Manual, error)
    Delete(e *entity.Manual) error
  }
  Manual struct {}
)

func NewManual() IManual {
  return &Manual{}
}

func (repo *Manual)All(entities *entity.Manuals) (*entity.Manuals, error) {
  err :=db.New().Find(entities).Error
  if err != nil {
    return nil, err
  }
  return entities, err
}

func (repo *Manual)ByID(entity *entity.Manual, id int64) (*entity.Manual, error) {
  err := db.New().Find(entity, id).Error
  if err != nil {
    return nil, err
  }
  return entity, nil
}

func (repo *Manual)Create(entity *entity.Manual) (*entity.Manual, error) {
  err := db.New().Create(entity).Error
  if err != nil {
    return nil, err
  }
  return entity, nil
}

func (repo *Manual)Update(entity *entity.Manual) (*entity.Manual, error) {
  err := db.New().Save(entity).Error
  if err != nil {
    return nil, err
  }
  return entity, nil
}

func (repo *Manual)Delete(entity *entity.Manual) error {
  err := db.New().Delete(entity).Error
  if err != nil {
    return err
  }
  return nil
}
