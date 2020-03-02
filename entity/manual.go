package entity

import "time"

type (
  Manual struct {
    ID          int64     `gorm:"column:id;primary_key"`
    Title       string    `gorm:"title"`
    Description string    `gorm:"description"`
    TagID       int64     `gorm:"tag_id"`
    CreatedAt   time.Time `gorm:"column:created_at"`
    UpdatedAt   time.Time `gorm:"column:updated_at`
  }

  // entityにManualsの配列を持たせてあげる
  Manuals []Manual
)

// entity constructorは持たせていない。usecaseで作成。
