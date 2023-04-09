CREATE TABLE "tsarka" (
                          "id" serial PRIMARY KEY,
                          "string" VARCHAR NOT NULL,
                          "max_substring" VARCHAR NOT NULL
);


CREATE INDEX ON "tsarka" ("string");