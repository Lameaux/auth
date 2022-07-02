-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "email" text,
    "email_verified_at" timestamp,
    "msisdn" bigint,
    "msisdn_verified_at" timestamp,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    "deleted_at" timestamp
);

CREATE UNIQUE INDEX "users_email" ON "users" ( "email" )
    WHERE "email" IS NOT NULL AND "deleted_at" IS NULL;

CREATE UNIQUE INDEX "users_msisdn" ON "users" ( "msisdn" )
    WHERE "msisdn" IS NOT NULL AND "deleted_at" IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
