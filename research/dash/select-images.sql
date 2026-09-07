WITH RECURSIVE
seg AS (
    SELECT
        r.id AS rep_id,
        st.period_id,
        st.media,
        COALESCE(st.start_number, 1) AS start_number,
        CAST(CEIL(
            CAST(SUBSTR(p.duration, 3) AS REAL)
            / (st.duration * 1.0 / COALESCE(st.timescale, 1))
        ) AS INTEGER) AS cnt
    FROM representation r
    JOIN segment_template st
      ON st.adaptation_set_id = r.adaptation_set_id
     AND st.period_id         = r.period_id
    JOIN period p
      ON p.id = st.period_id
    WHERE r.id = 'images'
),
numbers(i, rep_id, period_id, media, n) AS (
    SELECT 0, rep_id, period_id, media, start_number FROM seg
    UNION ALL
    SELECT i + 1, rep_id, period_id, media, n + 1
    FROM numbers
    WHERE i + 1 < (SELECT cnt FROM seg WHERE seg.period_id = numbers.period_id)
)
SELECT
    rep_id,
    period_id,
    n AS number,
    REPLACE(media, '$Number$', n) AS url
FROM numbers
ORDER BY period_id, n;
