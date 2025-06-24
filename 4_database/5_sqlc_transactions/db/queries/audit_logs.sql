-- name: CreateAuditLog :one
INSERT INTO audit_logs (
    user_id,
    action
) VALUES (
    $1,
    $2
)
RETURNING *;