import xml.etree.ElementTree as ET

NS = {"mpd": "urn:mpeg:dash:schema:mpd:2011"}
AS_ID = {"eng": "1", "spa": "2"}

def main(path: str) -> None:
    tree = ET.parse(path)
    for as_el in tree.getroot().iter(f"{{{NS['mpd']}}}AdaptationSet"):
        lang = as_el.get("lang")
        if lang not in AS_ID:
            continue
        timeline = as_el.find(f".//{{{NS['mpd']}}}SegmentTimeline")
        if timeline is None:
            continue
        as_id = AS_ID[lang]
        rows = [
            f"  ('{as_id}','0',{pos:>5},{s.get('d'):>6},{s.get('r','0')})"
            for pos, s in enumerate(timeline.findall(f"{{{NS['mpd']}}}S"), 1)
        ]
        print(f"\n-- AS {as_id} ({lang})")
        print("INSERT INTO segment_timeline (adaptation_set_id, period_id, position, d, r) VALUES")
        print(",\n".join(rows) + ";")

if __name__ == "__main__":
    import sys
    main(sys.argv[1] if len(sys.argv) > 1 else "kanopy.mpd")
