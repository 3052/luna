-- select-pssh.sql
SELECT DISTINCT
   r.id,
   cp.scheme_id_uri,
   cp.pssh
FROM representation r
JOIN content_protection cp
   ON cp.adaptation_set_id = r.adaptation_set_id
WHERE cp.pssh IS NOT NULL;
