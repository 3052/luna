-- select-images.sql
WITH RECURSIVE
seg AS (
    SELECT
        r.id AS rep_id,
        st.period_id,
        st.media,
        COALESCE(st.start_number, 1) AS start_number,
        CAST(CEIL(
            CAST(REPLACE(REPLACE(p.duration, 'PT', ''), 'S', '') AS REAL)
            / (st.duration * 1.0 / COALESCE(st.timescale, 1))
        ) AS INTEGER) AS cnt
    FROM representation r
    JOIN segment_template st
        -- PK is (adaptation_set_id, period_id): both conditions required
      ON st.adaptation_set_id = r.adaptation_set_id
     AND st.period_id         = r.period_id
    JOIN period p
      ON p.id = st.period_id
    WHERE r.id = 'images'
),
numbers(i, period_id, media, n, cnt) AS (
    SELECT 0, period_id, media, start_number, cnt FROM seg
    UNION ALL
    SELECT i + 1, period_id, media, n + 1, cnt
    FROM numbers
    WHERE i + 1 < cnt
)
SELECT
    MIN(period_id) AS period_id,
    REPLACE(media, '$Number$', n) AS url
FROM numbers
GROUP BY url
ORDER BY MIN(n);
