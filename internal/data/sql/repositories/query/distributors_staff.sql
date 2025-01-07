-- name: CreateDistributorStaff :one
INSERT INTO distributors_staff (id, distributors_id, user_id, role)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetDistributorStaffByID :one
SELECT * FROM distributors_staff
WHERE id = $1;

-- name: GetDistributorStaffByDistributorIDAndUserID :one
SELECT * FROM distributors_staff
WHERE distributors_id = $1 AND user_id = $2;

-- name: GetDistributorStaffByDistributorID :many
SELECT * FROM distributors_staff
WHERE distributors_id = $1;

-- name: GetDistributorStaffByUserID :many
SELECT * FROM distributors_staff
WHERE user_id = $1;

-- name: UpdateDistributorStaff :one
UPDATE distributors_staff
SET role = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteDistributorStaff :exec
DELETE FROM distributors_staff
WHERE id = $1;

-- name: DeleteDistributorStaffByDistributorIDAndUserId :exec
DELETE FROM distributors_staff
WHERE distributors_id = $1 AND user_id = $2;

-- name: ListDistributorStaff :many
SELECT * FROM distributors_staff
WHERE distributors_id = $1
ORDER BY created_at DESC;