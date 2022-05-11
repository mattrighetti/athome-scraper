base     := justfile_directory()
json_out := "/tmp/res.json"
links    := base + "/homes.txt"
db_path  := base + "/db.sqlite"

scrape:
    LINKS_PATH={{links}} \
    JSON_OUT={{json_out}} \
    node scraper/main.js

gobuild:
    cd {{base}}/loader; go build cmd/main.go

load: gobuild
    CONFIG_PATH={{base}}/loader/config.yaml \
    JSON_OUT={{json_out}} \
    DB_PATH={{db_path}} \
    {{base}}/loader/main

open:
    DB_PATH={{db_path}} \
    {{base}}/analyzer/venv/bin/jupyter notebook \
    {{base}}/analyzer/apartments.ipynb

fetch: scrape load

show: fetch open