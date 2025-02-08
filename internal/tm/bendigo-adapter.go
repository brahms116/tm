package tm

import (
	"fmt"
	"strconv"
	"time"
	"tm/internal/data"
)

func BendigoCsvRowAdapter(row []string) (data.AddTransactionParams, error) {
  dateStr := row[0]
  description := row[1]
  creditStr := row[2]
  debitStr := row[3]

  var amountStr string
  if creditStr != "" {
    amountStr = creditStr
  } else {
    amountStr = debitStr
  }

  id := dateStr + description

  time, err := time.Parse("02/01/2006", dateStr)
  if err != nil {
    return data.AddTransactionParams{}, fmt.Errorf("error parsing date: %w", err)
  }

  amount, err := strconv.ParseFloat(amountStr, 64)
  if err != nil {
    return data.AddTransactionParams{}, fmt.Errorf("error parsing amount: %w", err)
  }
  amountCents := int32(amount * 100)

  return data.AddTransactionParams{
    ID:          id,
    Date:        time,
    Description: description,
    AmountCents: amountCents,
  }, nil
}
