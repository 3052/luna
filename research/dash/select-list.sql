-- select-list.sql
SELECT 
    AVG(r.bandwidth) AS avg_bandwidth,
    MIN(r.width) AS min_width,
    MIN(r.height) AS min_height,
    MIN(r.codecs) AS min_codecs,
    MIN(r.mime_type) AS min_mime_type,
    MIN(a.label) AS min_label,
    r.id
FROM representation r
JOIN adaptation_set a
  ON a.id = r.adaptation_set_id
GROUP BY r.id
ORDER BY
   min_height,
   avg_bandwidth;
