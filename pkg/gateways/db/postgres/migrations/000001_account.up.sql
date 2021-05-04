BEGIN;

CREATE TABLE IF NOT EXISTS account
(
	id uuid PRIMARY KEY,
	name varchar(255) NOT NULL,
	cpf varchar(14) NOT NULL UNIQUE,
	secret varchar(100) NOT NULL,
	balance int default 0 NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;