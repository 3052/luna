-- select-t0.sql
SELECT
    REPLACE(st.media, '$Number$', st.start_number) AS media
FROM representation r
LEFT JOIN segment_template st
    -- PK is (adaptation_set_id, period_id): both conditions required
    ON st.adaptation_set_id = r.adaptation_set_id
   AND st.period_id         = r.period_id
WHERE r.id = 't0';
