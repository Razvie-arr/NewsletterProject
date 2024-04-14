- create url for your postgre 
```jsx
export POSTGRESQL_URL='postgres://login:password@localhost:5432/postgres?sslmode=disable' 
```
- run up all migrations

```jsx
 migrate -database ${POSTGRESQL_URL} -path repository/sql/migrations up 
```

- run down all migrations

```jsx
 migrate -database ${POSTGRESQL_URL} -path repository/sql/migrations down 
```

- after alternating table structure create migration and fill up and down files
```jsx
migrate create -ext sql -dir repository/sql/migrations -seq table_name
```