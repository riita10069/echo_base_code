package response

import "github.com/riita10069/echo_base_code/entity"

type(
  Manual struct {
    ID           int64       `json:"id"`
    Title        string      `json:"title" validate:"required"`
    Description  string      `json:"description"`
    TagID        int64       `json:"tag_id"`
  }

  Manuals []Manual
)

func NewManual(entity *entity.Manual) *Manual {
  if entity == nil {
    return nil
  } else {
    return &Manual{
      ID:          entity.ID,
      Title:       entity.Title,
      Description: entity.Description,
      TagID:       entity.TagID,
    }
  }
}

func NewManuals(entities *entity.Manuals)  *Manuals{
  ret := &Manuals{}
  for _, v := range *entities {
    *ret = append(*ret, *NewManual(&v))
  }
  return ret
}
