-- DROP TABLE IF EXISTS absent_reg;
-- DROP TABLE IF EXISTS pseudo_chapter;
-- DROP TABLE IF EXISTS pseudo_regulation;

-- CREATE TABLE public.pseudo_regulation (
--     r_id integer,
--     pseudo TEXT NOT NULL CHECK (pseudo != '')
-- );

-- CREATE TABLE public.pseudo_chapter (
--     c_id integer,
--     pseudo TEXT NOT NULL CHECK (pseudo != '')
-- );

-- CREATE TABLE public.absent_reg (
--     id SERIAL PRIMARY KEY,
--     pseudo TEXT NOT NULL CHECK (pseudo != ''),
--     done BOOLEAN NOT NULL DEFAULT false,
--     paragraph_id integer  
-- );

-- CREATE USER master_user WITH ENCRYPTED PASSWORD '031501';
-- GRANT CONNECT ON DATABASE master_db TO master_user;
-- GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO master_user;
-- GRANT ALL PRIVILEGES ON TABLE public.pseudo_regulation TO master_user;
-- GRANT ALL PRIVILEGES ON TABLE public.pseudo_chapter TO master_user;
-- GRANT ALL PRIVILEGES ON TABLE public.absent_reg TO master_user;

CREATE TABLE "absent_reg" (
	"id"	INTEGER,
	"pseudo"	TEXT NOT NULL,
	"done"	INTEGER NOT NULL DEFAULT 0,
	"paragraph_id"	INTEGER,
	PRIMARY KEY("id" AUTOINCREMENT)
);

CREATE TABLE "pseudo_chapter" (
	"c_id"	INTEGER UNIQUE,
	"pseudo"	TEXT NOT NULL
);

CREATE TABLE "pseudo_regulation" (
	"r_id"	INTEGER UNIQUE,
	"pseudo"	TEXT NOT NULL
);