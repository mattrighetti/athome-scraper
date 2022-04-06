json_out := "/tmp/res.json"
db_path := "./db.sqlite"

scrape:
    LINKS_PATH={{justfile_directory()}}/homes.txt \
    JSON_OUT={{json_out}} \
    node scraper/main.js

gobuild:
    cd {{justfile_directory()}}/loader; go build cmd/main.go

load: gobuild
    CONFIG_PATH={{justfile_directory()}}/loader/config.yaml \
    JSON_OUT={{json_out}} \
    DB_PATH={{db_path}} \
    {{justfile_directory()}}/loader/main

all: scrape load