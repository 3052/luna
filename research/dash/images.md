# images

OK that is better but still issue:
```
period_id: 0
      url: i/1971b3/images-204.jpg

period_id: 1
      url: i/1971b3/images-204.jpg
```

reference

```
<h3 class="heading settled" data-level="12.1" id="connectivity-duplicates"><span class="secno">12.1. </span><span class="content">Segment reference duplication during connected period transitions</span><a class="self-link" href="#connectivity-duplicates"></a></h3>
<p>As a <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧⓪">period</a> may start and/or end in the middle of a <a data-link-type="dfn" href="#media-segment" id="ref-for-media-segment③⑨">media segment</a>, the same <a data-link-type="dfn" href="#media-segment" id="ref-for-media-segment④⓪">media segment</a> may simultaneously be referenced by two <a data-link-type="dfn" href="#period-connected" id="ref-for-period-connected①①">period-connected</a> adaptation sets, with one part of it scheduled for playback during the first <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧①">period</a> and the other part during the second <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧②">period</a>. This is likely to be the case when no <a data-link-type="dfn" href="#sample-timeline" id="ref-for-sample-timeline②②">sample timeline</a> discontinuity is introduced by the transition.</p>
<figure>
  <img height="160" src="Images/Timing/SegmentOverlapOnPeriodConnectivity.png" width="635"> 
 <figcaption>The same <a data-link-type="dfn" href="#media-segment" id="ref-for-media-segment④①">media segment</a> will often exist in two <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧③">periods</a> at a <a data-link-type="dfn" href="#period-connected" id="ref-for-period-connected①②">period-connected</a> transition. On the diagram, this is segment 4.</figcaption>
</figure>
<p>Clients SHOULD NOT present a <a data-link-type="dfn" href="#media-segment" id="ref-for-media-segment④②">media segment</a> twice when it occurs on both sides of a <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧④">period</a> transition in a <a data-link-type="dfn" href="#period-connected" id="ref-for-period-connected①③">period-connected</a> adaptation set.</p>
<p>Clients SHOULD ensure seamless playback of <a data-link-type="dfn" href="#period-connected" id="ref-for-period-connected①④">period-connected</a> adaptation sets in consecutive <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧⑤">periods</a>. Clients unable to ensure seamless playback MAY incur some amount of <a data-link-type="dfn" href="#time-shift" id="ref-for-time-shift②">time shift</a> at the <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧⑥">period</a> transition point provided that the resulting <a data-link-type="dfn" href="#time-shift" id="ref-for-time-shift③">time shift</a> is permitted by the timing model.</p>
<p class="note" role="note"><span class="marker">Note:</span> The exact mechanism that ensures seamless playback depends on client capabilities and will be implementation-specific. Any shared <a data-link-type="dfn" href="#media-segment" id="ref-for-media-segment④③">media segment</a> overlapping the <a data-link-type="dfn" href="#periods" id="ref-for-periods⑧⑦">period</a> boundary may need to be detected and deduplicated to avoid presenting it twice.</p>
```
