-- name: CreateDistributor :one
INSERT INTO distributors (id, name, owner_id)
VALUES ($1, $2, $3)
    RETURNING *;

-- name: GetDistributorByID :one
SELECT * FROM distributors
WHERE id = $1;

-- name: UpdateDistributorName :one
UPDATE distributors
SET name = $2, updated_at = NOW()
WHERE id = $1
    RETURNING *;

-- name: ListDistributors :many
SELECT * FROM distributors
ORDER BY created_at DESC;