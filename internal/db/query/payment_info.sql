-- Insert payment info
-- name: CreatePaymentInfo
INSERT INTO payment_info (id, user_id, account_no, routing_no, account_name, bank_name)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5);

-- Get payment info by user ID
-- name: GetPaymentInfoByUserID
SELECT * FROM payment_info WHERE user_id = $1;

-- Update payment info
-- name: UpdatePaymentInfo
UPDATE payment_info 
SET account_no = $2, routing_no = $3, account_name = $4, bank_name = $5
WHERE user_id = $1;