version: "0.1"
database:
  dsn : "host=localhost user=postgres password=password dbname=tm_dev port=5432 sslmode=disable"
  db  : "postgres"
  tables:
    - "tm_transaction"
  onlyModel : true
  outPath: "./internal/orm/model"
  fieldNullable: true
