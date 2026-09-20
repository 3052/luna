# dash

~~~
# 1. load schema first (creates the tables)
sqlite3 edcbee.db '.read create.sql'

# 2. then load data (tables must already exist)
sqlite3 edcbee.db '.read insert.sql'
~~~

then:

~~~
hboMax e=086662ec-0ac6-4603-8ceb-a62be62b3d1d
~~~
