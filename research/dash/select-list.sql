-- select-list.sql
SELECT
   CAST(AVG(r.bandwidth) AS INTEGER)  bandwidth,
   MIN(r.width)  width,
   MIN(r.height)  height,
   MIN(r.codecs)  codecs,
   MIN(r.mime_type)  mime_type,
   MIN(a.label)  label,
   r.id representation
FROM representation r
JOIN adaptation_set a
   ON a.id        = r.adaptation_set_id
GROUP BY r.id
ORDER BY
   CASE r.mime_type
      WHEN 'video/mp4' THEN 3
      WHEN 'audio/mp4' THEN 2
      ELSE 1
   END,
   r.height;
