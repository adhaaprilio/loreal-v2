CREATE SEQUENCE tag_id_seq;

CREATE TABLE IF NOT EXISTS Tag (
    tag_id VARCHAR(255) NOT NULL PRIMARY KEY DEFAULT nextval('tag_id_seq'),
    name  VARCHAR(255)
);