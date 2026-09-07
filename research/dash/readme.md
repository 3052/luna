# dash

~~~
# 1. load schema first (creates the tables)
tursodb edcbee.db ".read schema.sql"

# 2. then load data (tables must already exist)
tursodb edcbee.db ".read data.sql"
~~~
