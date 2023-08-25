-- +goose Up
-- +goose StatementBegin
CREATE TABLE "customer" (
    "id" INTEGER GENERATED ALWAYS AS IDENTITY,
    "username" VARCHAR(255),
    "email" VARCHAR(360) NOT NULL,
    "email_confirmation" BOOLEAN NOT NULL DEFAULT false,
    "password_hash" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY ("id")
);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER customer_set_updated_at_trigger
BEFORE UPDATE ON "customer"
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER customer_set_updated_at_trigger ON customer;
DROP FUNCTION IF EXISTS set_updated_at;
DROP TABLE customer;
-- +goose StatementEnd
