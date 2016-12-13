# startups-referential

## Usage

### Local database
Launch a mysql container : sudo docker run -d -p 3306:3306 -e MYSQL_PASS="password" -e MYSQL_USER="admin" tutum/mysql
Test connection : mysql -uadmin -ppassword -h127.0.0.1 -P3306

Go to the scripts location (sql) and then run
`for script in $(ls *.sql); do mysql -uadmin -ppassword -h127.0.0.1 -P3306 startupdb < $script ; done` to create data structure
`for script in $(ls testing/*.sql); do mysql -uadmin -ppassword -h127.0.0.1 -P3306 startupdb < $script ; done` to populate test data

Launch server : `./startups-referential-api -config sample.toml` where toml is a configuration file like this

```
[database]
host = "127.0.0.1"
port = 3306
user = "admin"
password = "password"
database = "startupdb"
```
