-- name: AddTransaction :execrows
insert into tm_transaction
  (id, date, description, amount_cents, category_id)
values ($1, $2, $3, $4, $5)
on conflict do nothing;

-- name: ListTransactions :many
select *
from tm_transaction
where date >= $1 and date <= $2
order by date desc;

-- name: YearlyTimeline :many
select
  date_trunc('month', date)::date as month,
  sum(case when amount_cents > 0 then amount_cents else 0 end)::int as earnings,
  (-1 * sum(case when amount_cents < 0 then amount_cents else 0 end))::int as spendings
from tm_transaction
where date >= $1 and date < $2
group by month
order by month asc;

-- name: SummariseTransactions :one
select
  sum(case when amount_cents > 0 then amount_cents else 0 end) as earnings,
  -1 * sum(case when amount_cents < 0 then amount_cents else 0 end) as spendings
from tm_transaction
where date >= $1 and date < $2;

-- name: SummariseTransactionsU100 :one
select
  sum(case when amount_cents > 0 then amount_cents else 0 end) as earnings,
  -1 * sum(case when amount_cents < 0 then amount_cents else 0 end) as spendings
from tm_transaction
where date >= $1 and date < $2 and amount_cents >= -10000;

-- name: TopSpendings :many
select *
from tm_transaction
where date >= $1 and date < $2 and amount_cents < 0
order by amount_cents asc;

-- name: TopEarnings :many
select *
from tm_transaction
where date >= $1 and date < $2 and amount_cents > 0
order by amount_cents desc;

-- name: TopSpendingsU100 :many
select *
from tm_transaction
where date >= $1 and date < $2 and amount_cents < 0 and amount_cents >= -10000
order by amount_cents asc;

