-- name: AddTransaction :execrows
insert into tm_transaction
  (id, date, description, amount_cents, category_id)
values ($1, $2, $3, $4, $5)
on conflict do nothing;

-- name: ListTransactions :many
select id, date, description, amount_cents, category_id
from tm_transaction
where date >= $1 and date <= $2
order by date desc;
