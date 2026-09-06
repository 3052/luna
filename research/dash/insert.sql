INSERT INTO segment_template (adaptation_set_id, period_id, duration, media, presentation_time_offset, start_number, timescale) VALUES
('7','0',NULL,'t/aa517e/t0/$Number$.vtt',NULL,1,1000),
('8','0',5,'i/1971b3/images-$Number$.jpg',NULL,0,NULL),
('7','1',NULL,'t/aa517e/t0/$Number$.vtt',1020519,2,1000),
('8','1',5,'i/1971b3/images-$Number$.jpg',NULL,204,NULL),
('7','2',NULL,'t/aa517e/t0/$Number$.vtt',1756629,3,1000),
('8','2',5,'i/1971b3/images-$Number$.jpg',NULL,351,NULL),
('7','3',NULL,'t/aa517e/t0/$Number$.vtt',2794291,4,1000),
('8','3',5,'i/1971b3/images-$Number$.jpg',NULL,558,NULL),
('7','4',NULL,'t/aa517e/t0/$Number$.vtt',3758045,5,1000),
('8','4',5,'i/1971b3/images-$Number$.jpg',NULL,751,NULL),
('7','5',NULL,'t/aa517e/t0/$Number$.vtt',4637674,6,1000),
('8','5',5,'i/1971b3/images-$Number$.jpg',NULL,927,NULL);

INSERT INTO mpd (media_presentation_duration_sec) VALUES (5660.613291666667);

INSERT INTO period (id, duration_sec) VALUES
('0', 1020.5195),
('1', 736.1103750000001),
('2', 1037.6616249999997),
('3', 963.7544583333333),
('4', 879.6287499999999),
('5', 1022.9385833333336);

INSERT INTO adaptation_set (id, period_id, lang, label) VALUES
('0','0','en-US',NULL), ('1','0','en-US',NULL), ('2','0',NULL,NULL), ('3','0',NULL,NULL),
('4','0',NULL,NULL), ('5','0',NULL,NULL), ('6','0',NULL,NULL), ('7','0','en-US','en-US CC'), ('8','0',NULL,NULL),
('0','1','en-US',NULL), ('1','1','en-US',NULL), ('2','1',NULL,NULL), ('3','1',NULL,NULL),
('4','1',NULL,NULL), ('5','1',NULL,NULL), ('6','1',NULL,NULL), ('7','1','en-US','en-US CC'), ('8','1',NULL,NULL),
('0','2','en-US',NULL), ('1','2','en-US',NULL), ('2','2',NULL,NULL), ('3','2',NULL,NULL),
('4','2',NULL,NULL), ('5','2',NULL,NULL), ('6','2',NULL,NULL), ('7','2','en-US','en-US CC'), ('8','2',NULL,NULL),
('0','3','en-US',NULL), ('1','3','en-US',NULL), ('2','3',NULL,NULL), ('3','3',NULL,NULL),
('4','3',NULL,NULL), ('5','3',NULL,NULL), ('6','3',NULL,NULL), ('7','3','en-US','en-US CC'), ('8','3',NULL,NULL),
('0','4','en-US',NULL), ('1','4','en-US',NULL), ('2','4',NULL,NULL), ('3','4',NULL,NULL),
('4','4',NULL,NULL), ('5','4',NULL,NULL), ('6','4',NULL,NULL), ('7','4','en-US','en-US CC'), ('8','4',NULL,NULL),
('0','5','en-US',NULL), ('1','5','en-US',NULL), ('2','5',NULL,NULL), ('3','5',NULL,NULL),
('4','5',NULL,NULL), ('5','5',NULL,NULL), ('6','5',NULL,NULL), ('7','5','en-US','en-US CC'), ('8','5',NULL,NULL);

INSERT INTO representation (id, period_id, adaptation_set_id, codecs, bandwidth, mime_type, width, height, base_url) VALUES
-- period 0
('a0','0','0','ec-3',258322,'audio/mp4',NULL,NULL,'a/527443/a0.mp4'),
('a1','0','1','mp4a.40.5',66945,'audio/mp4',NULL,NULL,'a/527443/a1.mp4'),
('v0','0','2','avc1.64001e',524004,'video/mp4',768,416,'v/2fda2a/v0.mp4'),
('v1','0','2','avc1.64001e',887999,'video/mp4',768,416,'v/2fda2a/v1.mp4'),
('v2','0','2','avc1.64001e',1492560,'video/mp4',768,416,'v/2fda2a/v2.mp4'),
('v4','0','2','avc1.64001e',2488910,'video/mp4',768,416,'v/2fda2a/v4.mp4'),
('v5','0','2','avc1.64001f',4025121,'video/mp4',1024,554,'v/2fda2a/v5.mp4'),
('v3','0','3','avc1.64001f',6047005,'video/mp4',1280,692,'v/2fda2a/v3.mp4'),
('v6','0','4','hvc1.2.4.L93.90',442916,'video/mp4',1024,554,'v/2fda2a/v6.mp4'),
('v9','0','5','hvc1.2.4.L93.90',838074,'video/mp4',1280,692,'v/2fda2a/v9.mp4'),
('v7','0','6','hvc1.2.4.L120.90',1650628,'video/mp4',1920,1038,'v/2fda2a/v7.mp4'),
('v8','0','6','hvc1.2.4.L120.90',3048962,'video/mp4',1920,1038,'v/2fda2a/v8.mp4'),
('t0','0','7',NULL,13,'text/vtt',NULL,NULL,NULL),
('images','0','8',NULL,7341,'image/jpeg',352,190,NULL),
-- period 1
('a0','1','0','ec-3',258322,'audio/mp4',NULL,NULL,'a/527443/a0.mp4'),
('a1','1','1','mp4a.40.5',65968,'audio/mp4',NULL,NULL,'a/527443/a1.mp4'),
('v0','1','2','avc1.64001e',410484,'video/mp4',768,416,'v/2fda2a/v0.mp4'),
('v1','1','2','avc1.64001e',706578,'video/mp4',768,416,'v/2fda2a/v1.mp4'),
('v2','1','2','avc1.64001e',1194042,'video/mp4',768,416,'v/2fda2a/v2.mp4'),
('v4','1','2','avc1.64001e',1961229,'video/mp4',768,416,'v/2fda2a/v4.mp4'),
('v5','1','2','avc1.64001f',3107353,'video/mp4',1024,554,'v/2fda2a/v5.mp4'),
('v3','1','3','avc1.64001f',4484690,'video/mp4',1280,692,'v/2fda2a/v3.mp4'),
('v6','1','4','hvc1.2.4.L93.90',455717,'video/mp4',1024,554,'v/2fda2a/v6.mp4'),
('v9','1','5','hvc1.2.4.L93.90',868520,'video/mp4',1280,692,'v/2fda2a/v9.mp4'),
('v7','1','6','hvc1.2.4.L120.90',1687631,'video/mp4',1920,1038,'v/2fda2a/v7.mp4'),
('v8','1','6','hvc1.2.4.L120.90',3140036,'video/mp4',1920,1038,'v/2fda2a/v8.mp4'),
('t0','1','7',NULL,9,'text/vtt',NULL,NULL,NULL),
('images','1','8',NULL,7341,'image/jpeg',352,190,NULL),
-- period 2
('a0','2','0','ec-3',258322,'audio/mp4',NULL,NULL,'a/527443/a0.mp4'),
('a1','2','1','mp4a.40.5',66070,'audio/mp4',NULL,NULL,'a/527443/a1.mp4'),
('v0','2','2','avc1.64001e',535991,'video/mp4',768,416,'v/2fda2a/v0.mp4'),
('v1','2','2','avc1.64001e',937521,'video/mp4',768,416,'v/2fda2a/v1.mp4'),
('v2','2','2','avc1.64001e',1560554,'video/mp4',768,416,'v/2fda2a/v2.mp4'),
('v4','2','2','avc1.64001e',2492206,'video/mp4',768,416,'v/2fda2a/v4.mp4'),
('v5','2','2','avc1.64001f',4228300,'video/mp4',1024,554,'v/2fda2a/v5.mp4'),
('v3','2','3','avc1.64001f',6372980,'video/mp4',1280,692,'v/2fda2a/v3.mp4'),
('v6','2','4','hvc1.2.4.L93.90',469027,'video/mp4',1024,554,'v/2fda2a/v6.mp4'),
('v9','2','5','hvc1.2.4.L93.90',904778,'video/mp4',1280,692,'v/2fda2a/v9.mp4'),
('v7','2','6','hvc1.2.4.L120.90',1727847,'video/mp4',1920,1038,'v/2fda2a/v7.mp4'),
('v8','2','6','hvc1.2.4.L120.90',3255897,'video/mp4',1920,1038,'v/2fda2a/v8.mp4'),
('t0','2','7',NULL,12,'text/vtt',NULL,NULL,NULL),
('images','2','8',NULL,7341,'image/jpeg',352,190,NULL),
-- period 3
('a0','3','0','ec-3',258345,'audio/mp4',NULL,NULL,'a/527443/a0.mp4'),
('a1','3','1','mp4a.40.5',66170,'audio/mp4',NULL,NULL,'a/527443/a1.mp4'),
('v0','3','2','avc1.64001e',467035,'video/mp4',768,416,'v/2fda2a/v0.mp4'),
('v1','3','2','avc1.64001e',869263,'video/mp4',768,416,'v/2fda2a/v1.mp4'),
('v2','3','2','avc1.64001e',1460546,'video/mp4',768,416,'v/2fda2a/v2.mp4'),
('v4','3','2','avc1.64001e',2417315,'video/mp4',768,416,'v/2fda2a/v4.mp4'),
('v5','3','2','avc1.64001f',4050882,'video/mp4',1024,554,'v/2fda2a/v5.mp4'),
('v3','3','3','avc1.64001f',6180550,'video/mp4',1280,692,'v/2fda2a/v3.mp4'),
('v6','3','4','hvc1.2.4.L93.90',420946,'video/mp4',1024,554,'v/2fda2a/v6.mp4'),
('v9','3','5','hvc1.2.4.L93.90',822560,'video/mp4',1280,692,'v/2fda2a/v9.mp4'),
('v7','3','6','hvc1.2.4.L120.90',1690248,'video/mp4',1920,1038,'v/2fda2a/v7.mp4'),
('v8','3','6','hvc1.2.4.L120.90',3004800,'video/mp4',1920,1038,'v/2fda2a/v8.mp4'),
('t0','3','7',NULL,16,'text/vtt',NULL,NULL,NULL),
('images','3','8',NULL,7341,'image/jpeg',352,190,NULL),
-- period 4
('a0','4','0','ec-3',258357,'audio/mp4',NULL,NULL,'a/527443/a0.mp4'),
('a1','4','1','mp4a.40.5',65881,'audio/mp4',NULL,NULL,'a/527443/a1.mp4'),
('v0','4','2','avc1.64001e',438360,'video/mp4',768,416,'v/2fda2a/v0.mp4'),
('v1','4','2','avc1.64001e',793631,'video/mp4',768,416,'v/2fda2a/v1.mp4'),
('v2','4','2','avc1.64001e',1372900,'video/mp4',768,416,'v/2fda2a/v2.mp4'),
('v4','4','2','avc1.64001e',2296372,'video/mp4',768,416,'v/2fda2a/v4.mp4'),
('v5','4','2','avc1.64001f',3833673,'video/mp4',1024,554,'v/2fda2a/v5.mp4'),
('v3','4','3','avc1.64001f',5781709,'video/mp4',1280,692,'v/2fda2a/v3.mp4'),
('v6','4','4','hvc1.2.4.L93.90',425355,'video/mp4',1024,554,'v/2fda2a/v6.mp4'),
('v9','4','5','hvc1.2.4.L93.90',828698,'video/mp4',1280,692,'v/2fda2a/v9.mp4'),
('v7','4','6','hvc1.2.4.L120.90',1614424,'video/mp4',1920,1038,'v/2fda2a/v7.mp4'),
('v8','4','6','hvc1.2.4.L120.90',2983909,'video/mp4',1920,1038,'v/2fda2a/v8.mp4'),
('t0','4','7',NULL,8,'text/vtt',NULL,NULL,NULL),
('images','4','8',NULL,7341,'image/jpeg',352,190,NULL),
-- period 5
('a0','5','0','ec-3',258433,'audio/mp4',NULL,NULL,'a/527443/a0.mp4'),
('a1','5','1','mp4a.40.5',66070,'audio/mp4',NULL,NULL,'a/527443/a1.mp4'),
('v0','5','2','avc1.64001e',464242,'video/mp4',768,416,'v/2fda2a/v0.mp4'),
('v1','5','2','avc1.64001e',887217,'video/mp4',768,416,'v/2fda2a/v1.mp4'),
('v2','5','2','avc1.64001e',1517167,'video/mp4',768,416,'v/2fda2a/v2.mp4'),
('v4','5','2','avc1.64001e',2519277,'video/mp4',768,416,'v/2fda2a/v4.mp4'),
('v5','5','2','avc1.64001f',4151759,'video/mp4',1024,554,'v/2fda2a/v5.mp4'),
('v3','5','3','avc1.64001f',6166100,'video/mp4',1280,692,'v/2fda2a/v3.mp4'),
('v6','5','4','hvc1.2.4.L93.90',422652,'video/mp4',1024,554,'v/2fda2a/v6.mp4'),
('v9','5','5','hvc1.2.4.L93.90',835237,'video/mp4',1280,692,'v/2fda2a/v9.mp4'),
('v7','5','6','hvc1.2.4.L120.90',1635479,'video/mp4',1920,1038,'v/2fda2a/v7.mp4'),
('v8','5','6','hvc1.2.4.L120.90',3088870,'video/mp4',1920,1038,'v/2fda2a/v8.mp4'),
('t0','5','7',NULL,14,'text/vtt',NULL,NULL,NULL),
('images','5','8',NULL,7341,'image/jpeg',352,190,NULL);

INSERT INTO segment_base (representation_id, index_range) VALUES
('a0','658-17705'), ('a1','723-17770'),
('v0','794-17793'), ('v1','794-17793'), ('v2','794-17793'), ('v4','794-17793'),
('v5','795-17794'), ('v3','795-17794'),
('v6','861-17860'), ('v9','861-17860'), ('v7','861-17860'), ('v8','861-17860');

INSERT INTO initialization (representation_id, range) VALUES
('a0','0-657'), ('a1','0-722'),
('v0','0-793'), ('v1','0-793'), ('v2','0-793'), ('v4','0-793'),
('v5','0-794'), ('v3','0-794'),
('v6','0-860'), ('v9','0-860'), ('v7','0-860'), ('v8','0-860');

INSERT INTO segment_timeline (period_id, d) VALUES
('0',1020519), ('1',736110), ('2',1037661), ('3',963754), ('4',879628), ('5',1022938);

INSERT INTO content_protection (adaptation_set_id, scheme_id_uri, default_kid, pssh) VALUES
('0','urn:mpeg:dash:mp4protection:2011','01009967-8fa4-9384-f365-aba6a95d4a9f',NULL),
('0','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95',NULL,'AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('0','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed',NULL,'AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAEAmWePpJOE82WrpqldSp9I49yVmwY='),
('1','urn:mpeg:dash:mp4protection:2011','01009967-8fa4-9384-f365-aba6a95d4a9f',NULL),
('1','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95',NULL,'AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('1','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed',NULL,'AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAEAmWePpJOE82WrpqldSp9I49yVmwY='),
('2','urn:mpeg:dash:mp4protection:2011','01019967-8fa4-9384-f365-aba6a95d4a9f',NULL),
('2','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95',NULL,'AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('2','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed',NULL,'AAAAbnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAE4SEAEAmWePpJOE82WrpqldSp8SEAEFmWePpJOE82WrpqldSp8SEAECmWeP'),
('3','urn:mpeg:dash:mp4protection:2011','01029967-8fa4-9384-f365-aba6a95d4a9f',NULL),
('3','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95',NULL,'AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('3','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed',NULL,'AAAAbnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAE4SEAEAmWePpJOE82WrpqldSp8SEAEFmWePpJOE82WrpqldSp8SEAECmWeP'),
('4','urn:mpeg:dash:mp4protection:2011','01019967-8fa4-9384-f365-aba6a95d4a9f',NULL),
('4','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95',NULL,'AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('4','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed',NULL,'AAAAbnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAE4SEAEAmWePpJOE82WrpqldSp8SEAEFmWePpJOE82WrpqldSp8SEAECmWeP'),
('5','urn:mpeg:dash:mp4protection:2011','01029967-8fa4-9384-f365-aba6a95d4a9f',NULL),
('5','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95',NULL,'AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('5','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed',NULL,'AAAAbnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAE4SEAEAmWePpJOE82WrpqldSp8SEAEFmWePpJOE82WrpqldSp8SEAECmWeP'),
('6','urn:mpeg:dash:mp4protection:2011','01059967-8fa4-9384-f365-aba6a95d4a9f',NULL),
('6','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95',NULL,'AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('6','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed',NULL,'AAAAbnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAAE4SEAEAmWePpJOE82WrpqldSp8SEAEFmWePpJOE82WrpqldSp8SEAECmWeP');

INSERT INTO role (adaptation_set_id, value) VALUES
('2','main'), ('3','main'), ('4','main'), ('5','main'), ('6','main'), ('7','caption');
