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
