import sys
import xml.etree.ElementTree as ET

MPD_NS = "urn:mpeg:dash:schema:mpd:2011"

# Restrict segment_timeline to these langs, or None to import timelines
# for every adaptation set.
AS_LANGS = None  # e.g. {"eng", "spa"}


def q(tag: str) -> str:
    return f"{{{MPD_NS}}}{tag}"


def sql_text(v: str | None) -> str:
    return "NULL" if v is None else "'" + v.replace("'", "''") + "'"


def sql_int(v: str | None) -> str:
    # int() fails loudly on a malformed numeric attribute
    return "NULL" if v is None else str(int(v))


def emit(table: str, columns: list, rows: list, note: str = "") -> None:
    print()
    if not rows:
        print(f"-- {table}: no rows" + (f" ({note})" if note else ""))
        return
    print(f"-- {table}")
    print(f"INSERT INTO {table} ({', '.join(columns)}) VALUES")
    print(",\n".join("  (" + ", ".join(row) + ")" for row in rows) + ";")


def main(path: str) -> None:
    root = ET.parse(path).getroot()

    # mpd: the <MPD> element itself -- always exactly one row per manifest.
    # base_url = MPD-level <BaseURL> text if present, else NULL.
    base_el = root.find(q("BaseURL"))
    base_url = ((base_el.text or "").strip() or None) if base_el is not None else None
    mpd_rows = [[sql_text(base_url)]]

    period_rows = []
    as_rows, rep_rows, st_rows, tl_rows, role_rows = [], [], [], [], []

    as_position = 0  # document-wide, 1-based
    for period_el in root.findall(q("Period")):
        pid_sql = sql_text(period_el.get("id"))
        period_rows.append([pid_sql, sql_text(period_el.get("duration"))])

        for as_el in period_el.findall(q("AdaptationSet")):
            as_position += 1
            lang = as_el.get("lang")

            label_el = as_el.find(q("Label"))
            label = ("".join(label_el.itertext()).strip() or None) if label_el is not None else None
            as_rows.append([str(as_position), pid_sql, sql_text(lang), sql_text(label)])

            for role_el in as_el.findall(q("Role")):
                role_rows.append([str(as_position), sql_text(role_el.get("value"))])

            st_el = as_el.find(q("SegmentTemplate"))
            if st_el is not None:
                st_rows.append([
                    str(as_position), pid_sql,
                    sql_int(st_el.get("duration")),
                    sql_text(st_el.get("media")),
                    sql_int(st_el.get("presentationTimeOffset")),
                    sql_int(st_el.get("startNumber")),
                    sql_int(st_el.get("timescale")),
                ])
                timeline = st_el.find(q("SegmentTimeline"))
                if timeline is not None and (AS_LANGS is None or lang in AS_LANGS):
                    for pos, s_el in enumerate(timeline.findall(q("S")), 1):
                        tl_rows.append([
                            str(as_position), pid_sql, str(pos),
                            sql_int(s_el.get("d")), sql_int(s_el.get("r")),
                        ])

            as_mime = as_el.get("mimeType")
            for rep_el in as_el.findall(q("Representation")):
                rep_base = rep_el.find(q("BaseURL"))
                base_url = ((rep_base.text or "").strip() or None) if rep_base is not None else None
                rep_rows.append([
                    sql_text(rep_el.get("id")), pid_sql, str(as_position),
                    sql_text(rep_el.get("codecs")), sql_int(rep_el.get("bandwidth")),
                    sql_text(rep_el.get("mimeType") or as_mime),
                    sql_int(rep_el.get("width")), sql_int(rep_el.get("height")),
                    sql_text(base_url),
                ])

    emit("mpd", ["base_url"], mpd_rows)
    emit("period", ["id", "duration"], period_rows)
    emit("adaptation_set", ["position", "period_id", "lang", "label"], as_rows)
    emit("representation", ["id", "period_id", "adaptation_set_position", "codecs",
                            "bandwidth", "mime_type", "width", "height", "base_url"], rep_rows)
    emit("segment_template", ["adaptation_set_position", "period_id", "duration", "media",
                               "presentation_time_offset", "start_number", "timescale"], st_rows)
    emit("segment_timeline", ["adaptation_set_position", "period_id", "position", "d", "r"], tl_rows)
    emit("role", ["adaptation_set_position", "value"], role_rows)


if __name__ == "__main__":
    main(sys.argv[1] if len(sys.argv) > 1 else "kanopy.mpd")
