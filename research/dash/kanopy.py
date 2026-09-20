import sys
import xml.etree.ElementTree as ET

NS = {"mpd": "urn:mpeg:dash:schema:mpd:2011"}
AS_LANGS = {"eng", "spa"}

def main(path: str) -> None:
    tree = ET.parse(path)
    root = tree.getroot()
    mpd_ns = NS["mpd"]
    for as_el in root.iter(f"{{{mpd_ns}}}AdaptationSet"):
        lang = as_el.get("lang")
        if lang not in AS_LANGS:
            continue
        timeline = as_el.find(f".//{{{mpd_ns}}}SegmentTimeline")
        if timeline is None:
            continue
        rows = []
        for pos, s in enumerate(timeline.findall(f"{{{mpd_ns}}}S"), 1):
            d = s.get("d")
            r = s.get("r")  # absent @r -> NULL
            r_sql = r if r is not None else "NULL"
            rows.append(f"  (NULL,NULL,{pos:>5},{d:>6},{r_sql})")
        print(f"\n-- AS lang={lang}")
        print("INSERT INTO segment_timeline (adaptation_set_id, period_id, position, d, r) VALUES")
        print(",\n".join(rows) + ";")

if __name__ == "__main__":
    main(sys.argv[1] if len(sys.argv) > 1 else "kanopy.mpd")
