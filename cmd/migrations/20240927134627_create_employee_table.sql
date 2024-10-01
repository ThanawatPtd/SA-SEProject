-- +goose Up
-- +goose StatementBegin
-- EMPLOYEE Table (Inherits from USER)
CREATE TABLE "employee" (
    id UUID PRIMARY KEY,
    salary REAL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    FOREIGN KEY (id) REFERENCES "user"(id) ON DELETE CASCADE
);

ALTER TABLE "employee" ALTER COLUMN id SET DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd