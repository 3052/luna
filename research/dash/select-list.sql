-- select-list.sql
SELECT
   AVG(r.bandwidth) AS bandwidth,
   MIN(r.width) AS width,
   MIN(r.height) AS height,
   MIN(r.codecs) AS codecs,
   MIN(r.mime_type) AS mime_type,
   MIN(a.label) AS label,
   r.id
FROM representation r
JOIN adaptation_set a
   -- PK is (id, period_id): both conditions required
   ON a.id        = r.adaptation_set_id
  AND a.period_id = r.period_id
GROUP BY r.id
ORDER BY
   height,
   bandwidth;
