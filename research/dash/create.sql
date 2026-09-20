-- create.sql

CREATE TABLE mpd (
   base_url TEXT
);

CREATE TABLE period (
   id           TEXT,
   duration     TEXT
);

CREATE TABLE adaptation_set (
   id        TEXT NOT NULL,
   period_id TEXT NOT NULL,
   lang      TEXT,
   label     TEXT
);

CREATE TABLE representation (
   id                TEXT NOT NULL,
   period_id         TEXT NOT NULL,
   adaptation_set_id TEXT NOT NULL,
   codecs            TEXT,
   bandwidth         INTEGER NOT NULL,
   mime_type         TEXT,
   width             INTEGER,
   height            INTEGER,
   base_url          TEXT
);

CREATE TABLE segment_template (
   adaptation_set_id        TEXT NOT NULL,
   period_id                TEXT NOT NULL,
   duration                 INTEGER,
   media                    TEXT,
   presentation_time_offset INTEGER,
   start_number             INTEGER,
   timescale                INTEGER
);

CREATE TABLE segment_timeline (
   adaptation_set_id TEXT NOT NULL,
   period_id         TEXT NOT NULL,
   position          INTEGER NOT NULL,
   d                 INTEGER NOT NULL,
   r                 INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE content_protection (
   adaptation_set_id TEXT NOT NULL,
   scheme_id_uri     TEXT NOT NULL,
   pssh              TEXT
);

CREATE TABLE role (
   adaptation_set_id TEXT,
   value             TEXT
);

-- create.sql
