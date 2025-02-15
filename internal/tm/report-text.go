package tm

import (
	"fmt"
	"strings"
	"text/template"
	"time"
	"tm/internal/data"
	"tm/pkg/contracts"
)

const tmpl = `---- Report for {{formatMonth .Month}} ----

-- Summary:
Total spending: {{formatCents .Summary.Spending}} ({{formatCents .SummaryComparison.Spending}})
Total small spending: {{formatCents .Summary.SmallSpending}} ({{formatCents .SummaryComparison.SmallSpending}})
Total earning: {{formatCents .Summary.Earning}} ({{formatCents .SummaryComparison.Earning}})
Net: {{formatCents .Summary.Net}} ({{formatCents .SummaryComparison.Net}})
{{- range .Periods}}


-- {{formatPeriodDate .EndDate}} ({{formatCents .Summary.SmallSpendingPerDay}} / {{formatCents .Summary.SpendingPerDay}}):

{{range .SmallSpends}}
{{formatTransaction .}}
{{- end}}
{{- end}}
`

func formatCents(cents int) string {
  dollars:= float64(cents) / 100
  if dollars < 0 {
    return fmt.Sprintf("-$%.2f", -dollars)
  
  }
  return fmt.Sprintf("$%.2f", dollars)
}

func formatMonth(month time.Time) string {
	return month.Format("January 2006")
}

func formatPeriodDate(date time.Time) string {
  date = date.AddDate(0, 0, -1)
	return date.Format("01-02")
}

func formatTransaction(t data.TmTransaction) string {
	return fmt.Sprintf("%s: %s", formatCents(int(t.AmountCents)), t.Description)
}

func textReport(report contracts.MonthReport) string {
	funcMap := template.FuncMap{
		"formatCents":       formatCents,
		"formatMonth":       formatMonth,
		"formatPeriodDate":  formatPeriodDate,
		"formatTransaction": formatTransaction,
	}

  tmpl, err := template.New("report").Funcs(funcMap).Parse(tmpl)
  if err != nil {
    panic(err)
  }

  var b strings.Builder
  err = tmpl.Execute(&b, report)

  if err != nil {
    panic(err)
  }
  return b.String()
}
