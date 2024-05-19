# App-Registry

# Init

```bash
go mod init github.com/gal16v8d/app-registry.git
```
then:
```bash
go mod tidy
```
add dependency:
```bash
go get -u {dependency_name_here}
```
ie:
```bash
go get -u github.com/gin-gonic/gin
```

# Db setup

Download mysql docker image:
```bash
docker pull mysql
```
Run instance like:
```bash
docker run -p 3306:3306 --name mysqldb -e MYSQL_ROOT_PASSWORD={your_pass_here} -d mysql
```

Conditional steps in case mysql db don't connect:
```bash
docker exec -it mysqldb /bin/bash
mysql -uroot -p -A
select user, host from mysql.user;
update mysql.user set host='%' where user='root';
flush privileges;
exit
```

# Run

1. Please be sure of create the mysql db using something similar to the `init_db.sql`
file located into the `dev_tools` folder
2. Then, go to `cmd/server` folder inspect the `.env.sample` file, you might need a `.env` file with your real user, pass and db name in the same format as you find in the sample.
3. Then

```bash
cd cmd/server
```
then:
```bash
go run .
```