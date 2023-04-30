# WebApp
Social network for "Travel Dream"

## DataBase:

In my project I use MySql so you must have mysql driver on your PC


## Config.yml:

If you want work with my code, you need create new file "config.yml" and input something as:
```
port: 8181

data:
  url: "login:password@/testdb" //default login: root
  table: "testdb.tablename"
```

P.S Why I save ```data``` in .yml? I don't know:) But, maybe, I will change it

## Requirements
- go 1.18
- MySql driver (installation instruction on <b>RUSSIAN</b>: <a href="http://dev.mysql.com/downloads/mysql/">Click</a>)

## Run Project

Use ```go run cmd/main.go``` to run application and enter in the address bar ```localhost:8181```
