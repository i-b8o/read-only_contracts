CREATE TABLE "absent_reg" (
	"id" INTEGER,
	"pseudo" TEXT NOT NULL,
	"done" INTEGER NOT NULL DEFAULT 0,
	"paragraph_id" INTEGER,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE "pseudo_chapter" (
	"c_id" INTEGER UNIQUE,
	"pseudo" TEXT NOT NULL
);
CREATE TABLE "pseudo_doc" (
	"d_id" INTEGER UNIQUE,
	"pseudo" TEXT NOT NULL
);