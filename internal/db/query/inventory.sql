-- Insert inventory record
-- name: CreateInventoryRecord :one
INSERT INTO inventory (
    id, version_id, branch_id, quantity, unit_cost, last_counted
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5
) RETURNING *;

-- Get inventory by ID
-- name: GetInventoryByID :one
SELECT * FROM inventory WHERE id = $1 LIMIT 1;

-- Get inventory by version ID
-- name: GetInventoryByVersionID :one
SELECT * FROM inventory WHERE version_id = $1 LIMIT 1;

-- Get inventory by branch ID
-- name: GetInventoryByBranchID :many
SELECT * FROM inventory WHERE branch_id = $1;

-- Update inventory
-- name: UpdateInventory :one
UPDATE inventory 
SET 
    quantity = $2, unit_cost = $3, last_counted = $4
WHERE id = $1 RETURNING *;