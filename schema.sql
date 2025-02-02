create table if not exists tm_transaction_category (
  id text not null default gen_random_uuid() primary key, 
  name text not null,
  description text not null
);

create table if not exists tm_transaction_classifying_rule (
  id text not null default gen_random_uuid() primary key,
  name text not null,
  description text not null,
  category_id text not null references tm_transaction_category(id) on delete cascade on update cascade,
  rule_type text not null,
  pattern text not null
);

create table if not exists tm_transaction (
  id text not null primary key,
  date date not null,
  description text not null,
  amount_cents integer not null,
  category_id text references tm_transaction_category(id) on delete set null on update cascade
);

create index if not exists tm_transaction_date_idx on tm_transaction(date);
create index if not exists tm_transaction_category_idx on tm_transaction(category_id);
