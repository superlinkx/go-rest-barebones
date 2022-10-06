-- name: CreateCustomer :one
INSERT INTO customers (first_name, last_name, email, phone)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetCustomer :one
SELECT *
FROM customers
WHERE id = $1
LIMIT 1;
-- name: UpdateCustomer :one
UPDATE customers
SET first_name = COALESCE(sqlc.narg(first_name), first_name),
  last_name = COALESCE(sqlc.narg(last_name), last_name),
  email = COALESCE(sqlc.narg(email), email),
  phone = COALESCE(sqlc.narg(phone), phone)
WHERE id = $1
RETURNING *;
-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;