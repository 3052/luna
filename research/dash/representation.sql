-- .mode line

SELECT 
    id,
    AVG(bandwidth)          AS avg_bandwidth,
    GROUP_CONCAT(period_id) AS period_ids,
    MIN(adaptation_set_id)  AS adaptation_set_id,
    MIN(base_url)           AS base_url,
    MIN(codecs)             AS codecs,
    MIN(height)             AS height,
    MIN(mime_type)          AS mime_type,
    MIN(width)              AS width
FROM representation
GROUP BY id;
