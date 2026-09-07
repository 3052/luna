.mode line
SELECT
    r.id,
    r.base_url,
    st.media
FROM representation r
LEFT JOIN segment_template st
       ON st.adaptation_set_id = r.adaptation_set_id
GROUP BY r.id;
