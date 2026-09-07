SELECT 
    AVG(r.bandwidth)    ,
    MIN(r.width)        ,
    MIN(r.height)       ,
    MIN(r.codecs)       ,
    MIN(r.mime_type)    ,
    MIN(a.label)        ,
    r.id
FROM representation r
JOIN adaptation_set a
  ON a.id = r.adaptation_set_id
GROUP BY r.id
ORDER BY
   height,
   avg_bandwidth;
