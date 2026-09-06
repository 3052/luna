CREATE TABLE segment_template (
    duration          INTEGER,
    media             TEXT,
    presentation_time_offset INTEGER,
    start_number      INTEGER,
    timescale         INTEGER,
    adaptation_set_id TEXT NOT NULL,
    period_id         TEXT NOT NULL,
    PRIMARY KEY (adaptation_set_id, period_id)
);

CREATE TABLE mpd (
    media_presentation_duration_sec REAL
);

CREATE TABLE period (
    id           TEXT PRIMARY KEY,
    duration_sec REAL
);

CREATE TABLE adaptation_set (
    id        TEXT NOT NULL,
    period_id TEXT NOT NULL,
    lang      TEXT,
    label     TEXT,
    PRIMARY KEY (id, period_id)
);

CREATE TABLE representation (
    id                  TEXT NOT NULL,
    period_id           TEXT NOT NULL,
    adaptation_set_id   TEXT NOT NULL,
    codecs              TEXT,
    bandwidth           INTEGER NOT NULL,
    mime_type           TEXT,
    width               INTEGER,
    height              INTEGER,
    base_url            TEXT,
    PRIMARY KEY (id, period_id)
);

CREATE TABLE segment_base (
    representation_id TEXT PRIMARY KEY,
    index_range       TEXT
);

CREATE TABLE initialization (
    representation_id TEXT PRIMARY KEY,
    range             TEXT
);
