-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE RULE "customer_soft_delete" AS ON DELETE TO "customer" DO INSTEAD (
    UPDATE "customer" SET "deleted_at" = NOW() WHERE "id" = old.id AND "deleted_at" IS NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP RULE "customer_soft_delete" ON "customer";
-- +goose StatementEnd
