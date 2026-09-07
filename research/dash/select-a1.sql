SELECT
   DISTINCT r.base_url,
   init.range,
   sb.index_range
FROM representation r
LEFT JOIN segment_base sb
    ON sb.representation_id = r.id
LEFT JOIN initialization init
    ON init.representation_id = r.id
WHERE r.id = 'a1';

