
-- +migrate Up
CREATE TABLE IF NOT EXISTS "account"
(
    id                 serial                              NOT NULL,
    subject            varchar(256)                        NOT NULL,
    email              varchar(256)                        NOT NULL,

    -- timestamps
    created_date            timestamp                      DEFAULT now(),
    changed_date            timestamp                      DEFAULT now(),
    deleted_date            timestamp                      NULL,

    CONSTRAINT account_pk PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS account_subject_unique_index ON "account"("subject");

-- +migrate Down
DROP TABLE IF EXISTS "account";
