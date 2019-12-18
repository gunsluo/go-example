-- user
CREATE TABLE IF NOT EXISTS "user"
(
    id                 serial                   	        NOT NULL,
    subject            varchar(256)                        NOT NULL,

    -- timestamps
    created_date            timestamp                      DEFAULT now(),
    changed_date            timestamp                      DEFAULT now(),
    deleted_date            timestamp                               NULL,

    CONSTRAINT user_pk PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS user_subject_unique_index ON "user"("subject");

