BEGIN;

CREATE TABLE IF NOT EXISTS transfer
(
	id uuid PRIMARY KEY,
	account_origin_id uuid REFERENCES account,
	account_destination_id uuid REFERENCES account,
	amount int NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;