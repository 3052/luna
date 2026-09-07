# images

I dont think that is correct - looking at the input:
```
<SegmentTemplate media="i/1971b3/images-$Number$.jpg" duration="5" startNumber="0"></SegmentTemplate>
```

the Period looks like this:
```
<Period id="0" start="PT0.000000S" duration="PT1020.5195S">
```

SegmentTemplate timescale is the default 1 so count should be:
```
ceil(
   1020.5195 / (5 / 1)
)

ceil(1020.5195 / 5)

ceil(204.1039)

205
```

so the result should be 0 to 204 - but the query only returns 5 URL for the first Period:
```
   rep_id: images
period_id: 0
   number: 0
      url: i/1971b3/images-0.jpg

   rep_id: images
period_id: 0
   number: 1
      url: i/1971b3/images-1.jpg

   rep_id: images
period_id: 0
   number: 2
      url: i/1971b3/images-2.jpg

   rep_id: images
period_id: 0
   number: 3
      url: i/1971b3/images-3.jpg

   rep_id: images
period_id: 0
   number: 4
      url: i/1971b3/images-4.jpg
```

update the query or prove me wrong
