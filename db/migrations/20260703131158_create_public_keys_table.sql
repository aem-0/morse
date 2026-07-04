-- +goose Up
CREATE TABLE public_keys (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    public_key      BYTEA NOT NULL UNIQUE,   
    key_fingerprint VARCHAR(128) NOT NULL UNIQUE,  
    key_type        VARCHAR(32) NOT NULL,    
    device_name     VARCHAR(128) NOT NULL,
    is_active       BOOLEAN NOT NULL DEFAULT true,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_used_at    TIMESTAMPTZ,
    revoked_at      TIMESTAMPTZ
);

CREATE INDEX idx_public_keys_user_id ON public_keys(user_id);
-- +goose Down
DROP TABLE IF EXISTS public_keys;
