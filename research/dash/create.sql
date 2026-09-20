-- create.sql

CREATE TABLE mpd (
   base_url TEXT
);

CREATE TABLE period (
   id           TEXT,
   duration     TEXT
);

CREATE TABLE adaptation_set (
   position  INTEGER,
   period_id TEXT,
   lang      TEXT,
   label     TEXT
);

CREATE TABLE representation (
   id                      TEXT,
   period_id               TEXT,
   adaptation_set_position INTEGER,
   codecs                  TEXT,
   bandwidth               INTEGER,
   mime_type               TEXT,
   width                   INTEGER,
   height                  INTEGER,
   base_url                TEXT
);

CREATE TABLE segment_template (
   adaptation_set_position INTEGER,
   period_id               TEXT,
   duration                INTEGER,
   media                   TEXT,
   presentation_time_offset INTEGER,
   start_number            INTEGER,
   timescale               INTEGER
);

CREATE TABLE segment_timeline (
   adaptation_set_position INTEGER,
   period_id               TEXT,
   position                INTEGER,
   d                       INTEGER,
   r                       INTEGER
);

CREATE TABLE content_protection (
   adaptation_set_position INTEGER,
   scheme_id_uri           TEXT,
   pssh                    TEXT
);

CREATE TABLE role (
   adaptation_set_position INTEGER,
   value                   TEXT
);

-- create.sql
