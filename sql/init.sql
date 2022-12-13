DROP TABLE IF EXISTS paragraph;
DROP TABLE IF EXISTS chapter;
DROP TABLE IF EXISTS regulation;


CREATE TABLE regulation (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (NAME != '') UNIQUE,
    abbreviation TEXT,
    title TEXT,
    meta TEXT,
    keywords TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE chapter (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (name != ''),
    order_num SMALLINT NOT NULL CHECK (order_num >= 0),
    num TEXT,
    r_id integer REFERENCES regulation,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX chapter_regulation_idx ON chapter (r_id);

CREATE TABLE paragraph (
    id SERIAL PRIMARY KEY,
    paragraph_id INT NOT NULL CHECK (paragraph_id >= 0),
    order_num INT NOT NULL CHECK (order_num >= 0),
    is_table BOOLEAN NOT NULL,
    is_nft BOOLEAN NOT NULL,
    has_links BOOLEAN NOT NULL,
    class TEXT,
    content TEXT NOT NULL,
    c_id integer REFERENCES chapter
);
CREATE INDEX paragraph_chapter_idx ON paragraph (c_id);


CREATE USER reader WITH ENCRYPTED PASSWORD '031501';
GRANT CONNECT ON DATABASE main TO reader;
GRANT SELECT ON TABLE regulation TO reader;
GRANT SELECT ON TABLE chapter TO reader;
GRANT SELECT ON TABLE paragraph TO reader;

CREATE USER writer WITH ENCRYPTED PASSWORD '031501';
GRANT CONNECT ON DATABASE main TO writer;
GRANT ALL PRIVILEGES ON TABLE regulation TO writer;
GRANT ALL PRIVILEGES ON TABLE chapter TO writer;
GRANT ALL PRIVILEGES ON TABLE paragraph TO writer;
