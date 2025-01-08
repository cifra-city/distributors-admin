-- name: CreateDistributorEmployees :one
INSERT INTO distributors_employees (id, distributors_id, user_id, role)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetDistributorEmployeesByID :one
SELECT * FROM distributors_employees
WHERE id = $1;

-- name: GetDistributorEmployeesByDistributorIDAndUserID :one
SELECT * FROM distributors_employees
WHERE distributors_id = $1 AND user_id = $2;

-- name: GetDistributorEmployeesByDistributorID :many
SELECT * FROM distributors_employees
WHERE distributors_id = $1;

-- name: GetDistributorEmployeesByUserID :many
SELECT * FROM distributors_employees
WHERE user_id = $1;

-- name: UpdateDistributorEmployees :one
UPDATE distributors_employees
SET role = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateDistributorEmployeesByDistributorIDAndUserID :one
UPDATE distributors_employees
SET role = $3
WHERE distributors_id = $1 AND user_id = $2
    RETURNING *;

-- name: DeleteDistributorEmployees :exec
DELETE FROM distributors_employees
WHERE id = $1;

-- name: DeleteDistributorEmployeesByDistributorIDAndUserId :exec
DELETE FROM distributors_employees
WHERE distributors_id = $1 AND user_id = $2;

-- name: ListDistributorEmployees :many
SELECT * FROM distributors_employees
WHERE distributors_id = $1
ORDER BY created_at DESC;