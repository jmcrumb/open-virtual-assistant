package database

const (
	DBClear = `
DELETE FROM report;
DELETE FROM review;
DELETE FROM plugin;
DELETE FROM profile;
DELETE FROM account;
`

	DBInit = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Creating the tables
CREATE TABLE account (
    id UUID DEFAULT uuid_generate_v4 (),
    password varchar(60) NOT NULL,
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    email varchar(50) NOT NULL UNIQUE,
    last_login timestamp DEFAULT CURRENT_TIMESTAMP,
    date_joined date DEFAULT CURRENT_DATE,
    is_admin boolean DEFAULT false,
    PRIMARY KEY(id)
);

CREATE TABLE profile (
    account_id UUID,
    bio text DEFAULT '',
    photo bytea DEFAULT NULL,
    PRIMARY KEY(account_id),
    CONSTRAINT fk_profile_account_id
        FOREIGN KEY(account_id)
        REFERENCES account(id)
        ON DELETE CASCADE
);

CREATE TABLE plugin (
    id UUID DEFAULT uuid_generate_v4 (),
    name varchar(150) NOT NULL,
    publisher UUID NOT NULL,
    source_link varchar(150) NOT NULL UNIQUE,
    about text DEFAULT '',
    download_count integer DEFAULT 0,
    published_on date DEFAULT CURRENT_DATE,
    tsv_name tsvector,
    PRIMARY KEY(id),
    CONSTRAINT plugin_publisher_fk
        FOREIGN KEY(publisher)
        REFERENCES account(id)
);

CREATE TRIGGER plugin_name_tsv_update BEFORE INSERT OR UPDATE
ON plugin FOR EACH ROW EXECUTE PROCEDURE
tsvector_update_trigger(
	tsv_name, 'pg_catalog.english', name
);

CREATE INDEX index_plugin_name_tsv ON plugin USING gin (tsv_name);

CREATE TABLE review (
    id UUID DEFAULT uuid_generate_v4 (),
    source_review UUID DEFAULT NULL,
    account UUID NOT NULL,
    plugin UUID NOT NULL,
    rating real NOT NULL,
    content text DEFAULT NULL,
    PRIMARY KEY(id),
    CONSTRAINT review_src_review_fk
        FOREIGN KEY(source_review)
        REFERENCES review(id),
    CONSTRAINT review_account_fk
        FOREIGN KEY(account)
        REFERENCES account(id),
    CONSTRAINT review_plugin_fk
        FOREIGN KEY(plugin)
        REFERENCES plugin(id)
);

CREATE TABLE report (
    id UUID DEFAULT uuid_generate_v4 (),
    account UUID NOT NULL,
    plugin UUID NOT NULL,
    content text NOT NULL,
    is_resolved boolean DEFAULT false,
    PRIMARY KEY(id),
    CONSTRAINT report_account_fk
        FOREIGN KEY(account)
        REFERENCES account(id),
    CONSTRAINT report_plugin_fk
        FOREIGN KEY(plugin)
        REFERENCES plugin(id)
);
`
)
