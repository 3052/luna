-- select-item.sql
WITH RECURSIVE
seg AS (
   SELECT
      -- path: template media, or plain base_url when there is no template
      COALESCE(st.media, r.base_url) AS path,
      -- single URL when there is no template row
      COALESCE(st.start_number, 1) AS start_number,
      -- single URL when template has no duration, else range from period duration
      COALESCE(CAST(CEIL(
         CAST(REPLACE(REPLACE(p.duration, 'PT', ''), 'S', '') AS REAL)
         / (st.duration * 1.0 / COALESCE(st.timescale, 1))
      ) AS INTEGER), 1) AS cnt
   FROM representation r
   LEFT JOIN segment_template st
      -- PK is (adaptation_set_id, period_id): both conditions required
      ON st.adaptation_set_id = r.adaptation_set_id
     AND st.period_id         = r.period_id
   JOIN period p
      ON p.id = r.period_id
   WHERE r.id = 'a1'
-- WHERE r.id = 't0'
-- WHERE r.id = 'images'
),
numbers(path, n, last) AS (
   SELECT path, start_number, start_number + cnt - 1 FROM seg
   UNION ALL
   SELECT path, n + 1, last
   FROM numbers
   WHERE n < last
)
SELECT
   rtrim(m.base_url, replace(m.base_url, '/', ''))
   || REPLACE(path, '$Number$', n) AS url
FROM numbers
CROSS JOIN mpd m
GROUP BY url
ORDER BY MIN(n);
