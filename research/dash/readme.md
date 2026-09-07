# dash

~~~
# 1. load schema first (creates the tables)
sqlite3 edcbee.db '.read create.sql'

# 2. then load data (tables must already exist)
sqlite3 edcbee.db '.read insert.sql'
~~~
