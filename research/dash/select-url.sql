-- select-url.sql
WITH RECURSIVE
-- Strip the manifest filename off mpd.base_url:
-- ('https://.../edcbee.mpd' → 'https://.../')
walk(pos, last_slash) AS (
   SELECT 1, 0 FROM mpd
   UNION ALL
   SELECT pos + 1,
          CASE WHEN substr((SELECT base_url FROM mpd), pos, 1) = '/'
               THEN pos ELSE last_slash END
   FROM walk
   WHERE pos < length((SELECT base_url FROM mpd))
),
prefix AS (
   SELECT substr((SELECT base_url FROM mpd), 1, last_slash) AS url_prefix
   FROM walk
   ORDER BY pos DESC LIMIT 1
),
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
   -- hbo max:
   -- WHERE r.id = 'a1'
   -- WHERE r.id = 't0'
   --WHERE r.id = 'images'
   -- tubi:
   WHERE r.id = '0'
),
numbers(path, n, last) AS (
   SELECT path, start_number, start_number + cnt - 1 FROM seg
   UNION ALL
   SELECT path, n + 1, last
   FROM numbers
   WHERE n < last
)
SELECT
   IFNULL((SELECT url_prefix FROM prefix), '')
   || REPLACE(path, '$Number$', n) AS url
FROM numbers
GROUP BY url
ORDER BY MIN(n);
