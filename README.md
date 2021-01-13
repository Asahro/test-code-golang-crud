# Test Code Golang CRUD

This project is for Golang Developer Test Recruitment in Alfa Corp. This project build using Beego, MySql, and third party API (https://api-docs.alfadigital.id/)

## What You Need

This project required :
- golang v 1.15
- mySql v 10.1.38 (mariaDb)

## Installation

first clone project in a folder named crud. 
After clone the project, you will get on folder that contant db, postman colection and golang code.

### Installation db

create db project wirth name "crud_golang_test" and inport file crud_golang_test.sql

### Installation postman 

you just need to open postman and export collection and select "Testing Code Dog.postman_collection" file in the cloning project folder 

### Installation Golang 

first runing golang  

```bash
go run main.go
```

if there any dependency needed, install them using

```bash
go get git-url
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.
