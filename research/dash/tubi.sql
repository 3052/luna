-- tubi.sql

-- mpd: the <MPD> root has no BaseURL of its own.
INSERT INTO mpd (base_url) VALUES (NULL);

-- period
INSERT INTO period (id, duration) VALUES
    ('0', 'PT3587.654296875S');

-- adaptation_set
INSERT INTO adaptation_set (id, period_id, lang, label) VALUES
    ('0', '0', NULL,     'video'),
    ('1', '0', 'ko',     'Korean');

-- representation  (BaseURLs truncated to 99 bytes)
INSERT INTO representation (id, period_id, adaptation_set_id, codecs, bandwidth, mime_type, width, height, base_url) VALUES
    ('0', '0', '0', 'avc1.4d401e',  270816, 'video/mp4',  426,  240, '9zn0re6y.mp4?hdntl=exp=1790470419~acl=/3674894f-90bb-48db-9223-58c0bb46bdb6/*~hmac=bd64d0d482fe8f16'),
    ('2', '0', '0', 'avc1.4d401e',  657858, 'video/mp4',  640,  360, 'dw2vbwmr.mp4?hdntl=exp=1790470419~acl=/3674894f-90bb-48db-9223-58c0bb46bdb6/*~hmac=bd64d0d482fe8f16'),
    ('3', '0', '0', 'avc1.640028', 1281236, 'video/mp4',  854,  480, 'ku6qn50r.mp4?hdntl=exp=1790470419~acl=/3674894f-90bb-48db-9223-58c0bb46bdb6/*~hmac=bd64d0d482fe8f16'),
    ('4', '0', '0', 'avc1.640028', 1925419, 'video/mp4', 1024,  576, 'u6c00wgc.mp4?hdntl=exp=1790470419~acl=/3674894f-90bb-48db-9223-58c0bb46bdb6/*~hmac=bd64d0d482fe8f16'),
    ('5', '0', '0', 'avc1.640028', 2554304, 'video/mp4', 1280,  720, '4rla65ji.mp4?hdntl=exp=1790470419~acl=/3674894f-90bb-48db-9223-58c0bb46bdb6/*~hmac=bd64d0d482fe8f16'),
    ('1', '0', '1', 'mp4a.40.2',   132236, 'audio/mp4', NULL, NULL, '7p8ybomm.mp4?hdntl=exp=1790470419~acl=/3674894f-90bb-48db-9223-58c0bb46bdb6/*~hmac=bd64d0d482fe8f16');

-- role: only the audio AdaptationSet declares a Role of "main".
INSERT INTO role (adaptation_set_id, value) VALUES
    ('1', 'main');
