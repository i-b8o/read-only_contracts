DROP TABLE IF EXISTS absent_reg;
DROP TABLE IF EXISTS pseudo_chapter;
DROP TABLE IF EXISTS pseudo_doc;
DROP TABLE IF EXISTS link;
DROP MATERIALIZED VIEW IF EXISTS reg_search;
DROP INDEX IF EXISTS idx_search;
DROP TABLE IF EXISTS paragraph;
DROP TABLE IF EXISTS chapter;
DROP TABLE IF EXISTS doc;
CREATE TABLE doc (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (NAME != '') UNIQUE,
    abbreviation TEXT,
    title TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
ALTER TABLE doc
ADD COLUMN ts tsvector GENERATED ALWAYS AS (
        setweight(to_tsvector('russian', coalesce(name, '')), 'A') || setweight(to_tsvector('russian', coalesce(title, '')), 'B')
    ) STORED;
CREATE INDEX reg_ts_idx ON doc USING GIN (ts);
CREATE TABLE chapter (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (name != ''),
    order_num SMALLINT NOT NULL CHECK (order_num >= 0),
    num TEXT,
    doc_id integer REFERENCES doc,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
ALTER TABLE chapter
ADD COLUMN ts tsvector GENERATED ALWAYS AS (to_tsvector('russian', name)) STORED;
CREATE INDEX ch_ts_idx ON chapter USING GIN (ts);
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
ALTER TABLE paragraph
ADD COLUMN ts tsvector GENERATED ALWAYS AS (to_tsvector('russian', content)) STORED;
CREATE INDEX p_ts_idx ON paragraph USING GIN (ts);
CREATE MATERIALIZED VIEW reg_search AS
SELECT r.id AS "doc_id",
    r.name AS "r_name",
    NULL AS "c_id",
    NULL AS "c_name",
    CAST(NULL AS integer) AS "p_id",
    NULL AS "p_text",
    r.name AS "text",
    to_tsvector('russian', r.name) AS ts
FROM doc AS r
UNION
SELECT NULL AS "doc_id",
    r.name AS "r_name",
    c.id AS "c_id",
    c.name AS "c_name",
    NULL AS "p_id",
    NULL AS "p_text",
    c.name AS "text",
    to_tsvector('russian', c.name) AS ts
FROM chapter AS c
    INNER JOIN doc AS r ON r.id = c.doc_id
UNION
SELECT NULL AS "doc_id",
    r.name AS "r_name",
    c.id AS "c_id",
    c.name AS "c_name",
    p.paragraph_id AS "p_id",
    p.content AS "p_text",
    p.content AS "text",
    to_tsvector('russian', content) AS ts
FROM paragraph AS p
    INNER JOIN chapter AS c ON p.c_id = c.id
    INNER JOIN doc AS r ON c.doc_id = r.id;
create index idx_search on reg_search using GIN(ts);
CREATE TABLE pseudo_doc (
    doc_id integer,
    pseudo TEXT NOT NULL CHECK (pseudo != '')
);
CREATE TABLE pseudo_chapter (
    c_id integer,
    pseudo TEXT NOT NULL CHECK (pseudo != '')
);
CREATE TABLE absent_reg (
    id SERIAL PRIMARY KEY,
    pseudo TEXT NOT NULL CHECK (pseudo != ''),
    done BOOLEAN NOT NULL DEFAULT false,
    paragraph_id integer
);
CREATE TABLE link (
    id INT NOT NULL UNIQUE,
    paragraph_num INT NOT NULL CHECK (paragraph_num >= 0),
    c_id integer,
    doc_id integer
);
INSERT INTO doc ("name", "abbreviation", "title", "created_at")
VALUES (
        '?????? ???????????? ????????????',
        '???????????????????????? ???????????? ????????????',
        '?????????????????? ???????????? ????????????',
        '2023-01-01 00:00:00'
    );
INSERT INTO chapter (
        "name",
        "num",
        "order_num",
        "doc_id",
        "updated_at"
    )
VALUES (
        '?????? ???????????? ????????????',
        'I',
        1,
        1,
        '2023-01-01 00:00:00'
    ),
    (
        '?????? ???????????? ????????????',
        'II',
        2,
        1,
        '2023-01-01 00:00:00'
    ),
    (
        '?????? ?????????????? ????????????',
        'III',
        3,
        1,
        '2023-01-01 00:00:00'
    );
INSERT INTO paragraph (
        "paragraph_id",
        "order_num",
        "is_table",
        "is_nft",
        "has_links",
        "class",
        "content",
        "c_id"
    )
VALUES (
        1,
        1,
        false,
        false,
        true,
        'any-class',
        '???????????????????? <a id="dst101675"></a> ?????????????? <a href=''11111/a3a3a3/111''>??????????????????</a>',
        1
    ),
    (
        2,
        2,
        true,
        true,
        true,
        'any-class',
        '???????????????????? ?????????????? <a href=''372952/4e92c731969781306ebd1095867d2385f83ac7af/335104''>???????????? 5.14</a> ??????????????????',
        1
    ),
    (
        3,
        3,
        false,
        false,
        true,
        'any-class',
        '<a id=''335050''></a>???????????????????? ???????????????? ??????????????????<a href=''/document/cons_doc_LAW_2875/''>???????????????? N 2</a>.',
        1
    );
INSERT INTO pseudo_doc ("doc_id", "pseudo")
VALUES (1, 11111);
INSERT INTO pseudo_chapter ("c_id", "pseudo")
VALUES (3, 'a3a3a3');
INSERT INTO absent_reg ("pseudo", "done", "paragraph_id")
VALUES ('aaaaa', false, 1),
    ('bbbbb', true, 2),
    ('ccccc', false, 3);