-- tubi.sql

INSERT INTO mpd (base_url) VALUES ('https://nc-aka2.tubi.video/bb758e57-cc67-46a9-ad4e-69a611c04491/qgfb4rnwib.mpd?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjZG5fcHJlZml4IjoiaHR0cHM6Ly9uYy1ha2EyLnR1YmkudmlkZW8iLCJleHAiOjE3OTA0NDU2MDAsIm1lZGlhX3NpZyI6Mjg2ODUwOTl9.rI1jXqm3HfAeD771C3BN0xK9-C8UbUWy8pB-hc8h1L8');

INSERT INTO period (id, duration) VALUES
('0', 'PT6788.82373046875S');

INSERT INTO adaptation_set (id, period_id, lang, label) VALUES
('0', '0', 'en', 'English'),
('1', '0', NULL, 'video');

INSERT INTO representation (id, period_id, adaptation_set_id, codecs, bandwidth, mime_type, width, height, base_url) VALUES
('0','0','0','mp4a.40.2',       133907, 'audio/mp4', NULL, NULL, 'vk2ecg9c.mp4?hdntl=exp=1790461396~acl=/bb758e57-cc67-46a9-ad4e-69a611c04491/*~hmac=a469584eeac43de89ce733cca1ece9e4445bf704beb00e1b28f9d1b5b1779932'),
('1','0','1','hev1.1.6.L63.90',539558, 'video/mp4',  640,  360, 'ppal3ksp.mp4?hdntl=exp=1790461396~acl=/bb758e57-cc67-46a9-ad4e-69a611c04491/*~hmac=a469584eeac43de89ce733cca1ece9e4445bf704beb00e1b28f9d1b5b1779932'),
('2','0','1','hev1.1.6.L90.90', 1059786, 'video/mp4',  854,  480, '6pbqq9jd.mp4?hdntl=exp=1790461396~acl=/bb758e57-cc67-46a9-ad4e-69a611c04491/*~hmac=a469584eeac43de89ce733cca1ece9e4445bf704beb00e1b28f9d1b5b1779932'),
('3','0','1','hev1.1.6.L93.90', 1701215, 'video/mp4', 1024,  576, 'pc94r74v.mp4?hdntl=exp=1790461396~acl=/bb758e57-cc67-46a9-ad4e-69a611c04491/*~hmac=a469584eeac43de89ce733cca1ece9e4445bf704beb00e1b28f9d1b5b1779932'),
('4','0','1','hev1.1.6.L93.90', 1977285, 'video/mp4', 1280,  720, 'vd9d55ks.mp4?hdntl=exp=1790461396~acl=/bb758e57-cc67-46a9-ad4e-69a611c04491/*~hmac=a469584eeac43de89ce733cca1ece9e4445bf704beb00e1b28f9d1b5b1779932');

INSERT INTO content_protection (adaptation_set_id, scheme_id_uri, pssh) VALUES
('0','urn:mpeg:dash:mp4protection:2011',NULL),
('0','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('0','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAPJu/7tFE1Boxqf0cs+vnZI49yVmwY='),
('1','urn:mpeg:dash:mp4protection:2011',NULL),
('1','urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95','AAACJnBzc2gAAAAAmgTweZhAQoarkuZb4IhflQAAAgYGAgAAAQABAPwBPABXAFIATQBIAEUAQQBEAEUAUgAgAHgAbQBsAG4AcwA9'),
('1','urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed','AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAPJu/7tFE1Boxqf0cs+vnZI49yVmwY=');

INSERT INTO role (adaptation_set_id, value) VALUES
('0', 'main');
