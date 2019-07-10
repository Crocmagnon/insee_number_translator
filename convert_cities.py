# This expects a file named "correspondance-code-insee-code-postal.json"
# You can find one here : https://public.opendatasoft.com/explore/dataset/correspondance-code-insee-code-postal/export/

import json

with open("correspondance-code-insee-code-postal.json", "r") as f:
    data = json.load(f)
    mapping = dict()
    failed = []
    duplicates = []
    for base_item in data:
        item = base_item.get("fields")
        if not item:
            failed.append(base_item)
        insee_com = item.get("insee_com")
        if insee_com:
            if insee_com in mapping:
                duplicates.append(base_item)
            mapping[insee_com] = {
                "name": item.get("nom_comm"),
                "zip_code": item.get("postal_code"),
            }
        else:
            failed.append(base_item)
    import pprint

    with open("cities.py", "w") as writef:
        writef.write(pprint.pformat(mapping, indent=4, compact=True))
        writef.flush()
