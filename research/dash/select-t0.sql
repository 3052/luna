SELECT
    r.base_url,
    REPLACE(st.media, '$Number$', st.start_number) AS media
FROM representation r
LEFT JOIN segment_template st
    ON st.adaptation_set_id = r.adaptation_set_id
   AND st.period_id         = r.period_id
WHERE r.id = 't0';

