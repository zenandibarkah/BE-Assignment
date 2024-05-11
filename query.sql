CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE "users" (
    id bigserial PRIMARY KEY,
    username varchar NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL,
	phone varchar(15),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


CREATE TABLE "account_bank" (
    id bigserial PRIMARY KEY,
	user_id bigint NOT NULL,
	acc_name varchar NOT NULL,
	bank_name varchar NOT NULL,
	acc_number varchar NOT NULL,
	balance numeric(15) default 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	CONSTRAINT FK_user_id_account_bank FOREIGN KEY (user_id)
		REFERENCES users (id)
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON account_bank
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE "history_trans" (
    id bigserial PRIMARY KEY,
	user_id bigint NOT NULL,
	source_bank varchar NOT NULL,
	source_account_number varchar NOT NULL,
	destination_bank varchar NOT NULL,
	destination_account_number varchar NOT NULL,
	amount numeric(15) NOT NULL,
	trans_type varchar NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	CONSTRAINT FK_user_id_history_bank FOREIGN KEY (user_id)
		REFERENCES users (id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON history_trans
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
