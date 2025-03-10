-- 000001_bank_accounts.up.sql
CREATE TABLE IF NOT EXISTS bank_accounts (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bic VARCHAR(255) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    correspondent_account VARCHAR(255),
    account_number VARCHAR(255) NOT NULL,
    currency VARCHAR(10),
    comment TEXT,
    legal_entity_uuid UUID REFERENCES legal_entities(uuid) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE
);
