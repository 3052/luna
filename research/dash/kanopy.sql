-- kanopy.sql

-- Source <Period> is attribute-less: no @id, no @duration.
-- MPD-level @mediaPresentationDuration="PT2H17M30S" has no column in
-- this schema (mpd table only holds base_url, absent here) -> not stored.
INSERT INTO period (id, duration) VALUES
(NULL, NULL);

-- Source AdaptationSets carry no @id (id is optional in DASH; this MPD
-- omits it on all four). No period id exists to reference either.
-- Consequence: child tables cannot reference these rows by any
-- source-derived value; their adaptation_set_id is NULL below.
-- Linkage by document order is lost in this schema.
INSERT INTO adaptation_set (id, period_id, lang, label) VALUES
(NULL, NULL, NULL, NULL),
(NULL, NULL, 'eng', 'English'),
(NULL, NULL, 'spa', 'Spanish (Español)'),
(NULL, NULL, NULL, NULL);

-- representation @id values ARE present in source and stored verbatim.
-- All other association columns are NULL: no source ids exist to put there.
INSERT INTO representation (id, period_id, adaptation_set_id, codecs, bandwidth, mime_type, width, height, base_url) VALUES
('a5c648a8',NULL,NULL,'avc1.4D400C',207184,'video/mp4',256,144,NULL),
('e63ccbea',NULL,NULL,'avc1.4D401E',775738,'video/mp4',640,360,NULL),
('63337af0',NULL,NULL,'avc1.4D401E',1323215,'video/mp4',852,480,NULL),
('3c388bb9',NULL,NULL,'avc1.4D401E',1646859,'video/mp4',960,540,NULL),
('1e5450a9',NULL,NULL,'avc1.64001F',2465115,'video/mp4',1280,720,NULL),
('4648a436',NULL,NULL,'avc1.640029',6014064,'video/mp4',1920,1080,NULL),
('2336ae5e',NULL,NULL,'mp4a.40.2',203437,'audio/mp4',NULL,NULL,NULL),
('866e1e85',NULL,NULL,'mp4a.40.2',204629,'audio/mp4',NULL,NULL,NULL),
('dd5ce011-7c09-4ecb-b023-75e3a7048d1c/142c7c1d-ed31-449e-8491-6b624b135ad1',NULL,NULL,NULL,415,'image/jpeg',720,324,NULL),
('dd5ce011-7c09-4ecb-b023-75e3a7048d1c/c2ea61a0-cfdc-4070-94f5-6c7edea9fd63',NULL,NULL,NULL,1460,'image/jpeg',1600,720,NULL);

-- pssh values truncated to 99 bytes (non-URL rule).
-- Rows 1-6: sets 1 & 2 share KID da31b8cb-f5a6-4226-bb7b-18568d54df49
-- (no column for KIDs). Row 7-9: Spanish set, KID 9af58f7d-5907-4221-acf0-526ffb3fe366.
-- PlayReady mspr:pro payloads stored in pssh column (only payload column
-- available). mp4protection entries carry no pssh payload in source -> NULL.
INSERT INTO content_protection (adaptation_set_id, scheme_id_uri, pssh) VALUES
(NULL,'urn:mpeg:dash:mp4protection:2011',NULL),
(NULL,'urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','wgMAAAEAAQC4AzwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwA'),
(NULL,'urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAV3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADcIARIQ2jG4y/WmQia7exhWjVTfSRoLYnV5ZHJta2V5b3MiEB9RCCuJ6UF'),
(NULL,'urn:mpeg:dash:mp4protection:2011',NULL),
(NULL,'urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','wgMAAAEAAQC4AzwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwA'),
(NULL,'urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAV3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADcIARIQ2jG4y/WmQia7exhWjVTfSRoLYnV5ZHJta2V5b3MiEB9RCCuJ6UF'),
(NULL,'urn:mpeg:dash:mp4protection:2011',NULL),
(NULL,'urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','pAMAAAEAAQCaAzwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwA'),
(NULL,'urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAWXBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADkIARIQmvWPfVkHQiGs8FJv+z/jZhoLYnV5ZHJta2V5b3MiEJr1j31ZB0I');

INSERT INTO role (adaptation_set_id, value) VALUES
(NULL,'main'),
(NULL,'main');

-- URLs verbatim. @initialization exists on all four SegmentTemplates in
-- source; no column exists for it -> dropped, not fabricated.
-- Thumbnail set has no @timescale in source -> NULL (genuinely absent,
-- unlike the 1000 on the three media sets).
INSERT INTO segment_template (adaptation_set_id, period_id, duration, media, presentation_time_offset, start_number, timescale) VALUES
(NULL,NULL,NULL,'https://chunks.kanopy.com/file/prod-kcmschunks-001/35c20b70/29863579/$RepresentationID$/seg-$Number$.m4s?e=1789960590&token=1789874190-gomg14yqlzsIIL6CkYijdmaaa%2Fe%2BwZFOKo4szhBzJjA%3D',NULL,1,1000),
(NULL,NULL,NULL,'https://chunks.kanopy.com/file/prod-kcmschunks-001/35c20b70/29863579/$RepresentationID$/seg-$Number$.m4s?e=1789960590&token=1789874190-gomg14yqlzsIIL6CkYijdmaaa%2Fe%2BwZFOKo4szhBzJjA%3D',NULL,1,1000),
(NULL,NULL,NULL,'https://chunks.kanopy.com/file/prod-kcmschunks-001/23148ea6/e24aa06c/$RepresentationID$/seg-$Number$.m4s?e=1789960590&token=1789874190-gomg14yqlzsIIL6CkYijdmaaa%2Fe%2BwZFOKo4szhBzJjA%3D',NULL,1,1000),
(NULL,NULL,200,'https://chunks.kanopy.com/file/prod-kcmschunks-001/thumbnails/$RepresentationID$/sprite-$Number$.jpg?e=1789917390&token=1789874190-hihQFEkXrjk8D3FMQrqx8Cj9mdEbrETXWqlF0jZI2eo%3D',NULL,1,NULL);

-- segment_timeline: skipped per instruction (no <S> rows emitted).
-- When generated: one row per <S> in document order, position from 1,
-- d and r verbatim, r NULL/0 where source omits it.

-- kanopy.sql
