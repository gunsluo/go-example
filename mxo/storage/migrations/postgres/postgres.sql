CREATE TABLE IF NOT EXISTS "account"
(
    id                 serial                              NOT NULL,
    subject            varchar(256)                        NOT NULL,
    email              varchar(256)                        NOT NULL,
    name               varchar(256)                        NOT NULL,
    label              varchar(256)                        NULL,

    -- timestamps
    created_date            timestamp                      DEFAULT now(),
    changed_date            timestamp                      DEFAULT now(),
    deleted_date            timestamp                      NULL,

    CONSTRAINT account_pk PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS account_subject_unique_index ON "account"("subject");

CREATE TABLE IF NOT EXISTS "user"
(
    id                 serial                              NOT NULL,
    subject            varchar(256)                        NOT NULL,
    name               varchar(256)                        NULL,

    -- timestamps
    created_date            timestamp                      DEFAULT now(),
    changed_date            timestamp                      DEFAULT now(),
    deleted_date            timestamp                      NULL,

    CONSTRAINT user_pk PRIMARY KEY ("id"),
    CONSTRAINT user_account_subject_fk FOREIGN KEY ("subject") REFERENCES "account"("subject")
);
