DROP TABLE IF EXISTS public.paragraph;
DROP TABLE IF EXISTS public.chapter;
DROP TABLE IF EXISTS public.doc;
CREATE TABLE public.doc (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (NAME != '') UNIQUE,
    abbreviation TEXT,
    header TEXT,
    title TEXT,
    meta TEXT,
    keywords TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE public.chapter (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (name != ''),
    title TEXT,
    meta TEXT,
    keywords TEXT,
    order_num SMALLINT NOT NULL CHECK (order_num >= 0),
    num TEXT,
    doc_id integer REFERENCES doc,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX chapter_doc_idx ON public.chapter (doc_id);
CREATE TABLE public.paragraph (
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
CREATE INDEX paragraph_chaptedoc_idx ON public.paragraph (c_id);
CREATE USER reader WITH ENCRYPTED PASSWORD '031501';
GRANT CONNECT ON DATABASE main TO reader;
GRANT SELECT ON TABLE public.doc TO reader;
GRANT SELECT ON TABLE public.chapter TO reader;
GRANT SELECT ON TABLE public.paragraph TO reader;
CREATE USER writer WITH ENCRYPTED PASSWORD '031501';
GRANT CONNECT ON DATABASE main TO writer;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO writer;
GRANT ALL PRIVILEGES ON TABLE public.doc TO writer;
GRANT ALL PRIVILEGES ON TABLE public.chapter TO writer;
GRANT ALL PRIVILEGES ON TABLE public.paragraph TO writer;