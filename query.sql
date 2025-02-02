-- name: AddTransaction :execrows
insert into tm_transaction
  (id, date, description, amount_cents, category_id)
values ($1, $2, $3, $4, $5)
on conflict do nothing;
