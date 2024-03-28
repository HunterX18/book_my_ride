CREATE TABLE accounts (
    id bigserial PRIMARY KEY,
    owner varchar(20) NOT NULL,
    balance bigint NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transfers ( 
    id bigserial PRIMARY KEY,
    from_account_id bigint NOT NULL,
    to_account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP 
);

-- foreign key constraints create deadlock with concurrent requests (with FOR UPDATE command)
-- deadlocks are shown in postgresql but not in mysql
ALTER TABLE transfers ADD CONSTRAINT fk_transfers_from_account_id FOREIGN KEY (from_account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD CONSTRAINT fk_transfers_to_account_id FOREIGN KEY (to_account_id) REFERENCES accounts (id);

CREATE INDEX owner_index ON accounts (owner);
CREATE INDEX from_account_id_index ON transfers (from_account_id);
CREATE INDEX to_account_id_index ON transfers (to_account_id);

CREATE INDEX compound_index ON transfers (from_account_id, to_account_id);




