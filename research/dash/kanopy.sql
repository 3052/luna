-- kanopy.sql

INSERT INTO period (id, duration) VALUES
('0', 'PT2H17M30S');

INSERT INTO adaptation_set (id, period_id, lang, label) VALUES
('0','0',NULL,NULL),
('1','0','eng','English'),
('2','0','spa','Spanish (Español)'),
('3','0',NULL,NULL);

INSERT INTO representation (id, period_id, adaptation_set_id, codecs, bandwidth, mime_type, width, height, base_url) VALUES
('a5c648a8','0','0','avc1.4D400C',207184,'video/mp4',256,144,NULL),
('e63ccbea','0','0','avc1.4D401E',775738,'video/mp4',640,360,NULL),
('63337af0','0','0','avc1.4D401E',1323215,'video/mp4',852,480,NULL),
('3c388bb9','0','0','avc1.4D401E',1646859,'video/mp4',960,540,NULL),
('1e5450a9','0','0','avc1.64001F',2465115,'video/mp4',1280,720,NULL),
('4648a436','0','0','avc1.640029',6014064,'video/mp4',1920,1080,NULL),
('2336ae5e','0','1','mp4a.40.2',203437,'audio/mp4',NULL,NULL,NULL),
('866e1e85','0','2','mp4a.40.2',204629,'audio/mp4',NULL,NULL,NULL),
('dd5ce011-7c09-4ecb-b023-75e3a7048d1c/142c7c1d-ed31-449e-8491-6b624b135ad1','0','3',NULL,415,'image/jpeg',720,324,NULL),
('dd5ce011-7c09-4ecb-b023-75e3a7048d1c/c2ea61a0-cfdc-4070-94f5-6c7edea9fd63','0','3',NULL,1460,'image/jpeg',1600,720,NULL);

INSERT INTO content_protection (adaptation_set_id, scheme_id_uri, pssh) VALUES
('0','urn:mpeg:dash:mp4protection:2011',NULL),
('0','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','wgMAAAEAAQC4AzwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuA'),
('0','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAV3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADcIARIQ2jG4y/WmQia7exhWjVTfSRoLYnV5ZHJta2V5b3MiEB9RCCuJ6U'),
('1','urn:mpeg:dash:mp4protection:2011',NULL),
('1','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','wgMAAAEAAQC4AzwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuA'),
('1','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAV3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADcIARIQ2jG4y/WmQia7exhWjVTfSRoLYnV5ZHJta2V5b3MiEB9RCCuJ6U'),
('2','urn:mpeg:dash:mp4protection:2011',NULL),
('2','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','pAMAAAEAAQCaAzwAVwBSAE0ASABFAEEARABFAFIAIAB4AG0AbABuAHMAPQAiAGgAdAB0AHAAOgAvAC8AcwBjAGgAZQBtAGEAcwAuA'),
('2','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAWXBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADkIARIQmvWPfVkHQiGs8FJv+z/jZhoLYnV5ZHJta2V5b3MiEJr1j31ZB0');

INSERT INTO role (adaptation_set_id, value) VALUES
('1','main'),
('2','main');

INSERT INTO segment_template (adaptation_set_id, period_id, duration, media, presentation_time_offset, start_number, timescale) VALUES
('0','0',NULL,'https://chunks.kanopy.com/file/prod-kcmschunks-001/35c20b70/29863579/$RepresentationID$/seg-$Number$.m4s?e=1789960590&token=1789874190-gomg14yqlzsIIL6CkYijdmaaa%2Fe%2BwZFOKo4szhBzJjA%3D',NULL,1,1000),
('1','0',NULL,'https://chunks.kanopy.com/file/prod-kcmschunks-001/35c20b70/29863579/$RepresentationID$/seg-$Number$.m4s?e=1789960590&token=1789874190-gomg14yqlzsIIL6CkYijdmaaa%2Fe%2BwZFOKo4szhBzJjA%3D',NULL,1,1000),
('2','0',NULL,'https://chunks.kanopy.com/file/prod-kcmschunks-001/23148ea6/e24aa06c/$RepresentationID$/seg-$Number$.m4s?e=1789960590&token=1789874190-gomg14yqlzsIIL6CkYijdmaaa%2Fe%2BwZFOKo4szhBzJjA%3D',NULL,1,1000),
('3','0',200,'https://chunks.kanopy.com/file/prod-kcmschunks-001/thumbnails/$RepresentationID$/sprite-$Number$.jpg?e=1789917390&token=1789874190-hihQFEkXrjk8D3FMQrqx8Cj9mdEbrETXWqlF0jZI2eo%3D',NULL,1,NULL);

INSERT INTO segment_timeline (adaptation_set_id, period_id, position, d, r) VALUES
('0','0',1,5000,1649),
('0','0',2,  33,   0);

-- kanopy.sql
