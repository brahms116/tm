package tm

import (
	"time"
	"tm/internal/orm/model"
)

type importTransactionParams struct {
	id          string
	date        time.Time
	description string
	amountCents int32
}

func (p importTransactionParams) toDbModel() model.TmTransaction {
  return model.TmTransaction{
    ID:          p.id,
    Date:        p.date,
    Description: p.description,
    AmountCents: p.amountCents,
  }
}
