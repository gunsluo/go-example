

CREATE TABLE IF NOT EXISTS country (
    id bigserial NOT NULL,
    country_name varchar(128) NOT NULL,
    country_logo varchar(256),
    country_active boolean,
    country_snum smallint,
    country_num  integer,
    country_bnum  bigint,
    created_date timestamp DEFAULT now(),
    changed_date timestamp DEFAULT now(),
    deleted_date timestamp,
    country_code varchar(2) NOT NULL,

    CONSTRAINT country_country_name_key UNIQUE (country_name),
    CONSTRAINT country_country_code_key UNIQUE (country_code),
    CONSTRAINT country_pk PRIMARY KEY (id)
);

