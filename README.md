# sturdy-robot

Overview
A robust Go-based invoice management application with advanced database and API capabilities.

## Technology Stack

- Language: Go (Golang)
- Database: PostgreSQL
- ORM/Query Generation: SQLC
- Database Driver: pgx
- API Framework: Gin
- PDF Generation: gofpdf

## Project Architecture

### Database Design
![Nume](https://github.com/user-attachments/assets/421a4038-83a0-4829-a7f8-1573c2e98301)



- Normalized relational schema
- Supports multiple currencies
- Designed for scalability and data integrity

- Versioned SQL migration files
  Supports forward and rollback migrations
- Compile-time query validation

Key Features

Invoice creation and management
Currency stored and retrieved as DECIMAL for precise financial calculations
Detailed invoice tracking
PDF invoice generation
Comprehensive API endpoints


## Setup and Installation
Prerequisites

- Go 1.21+
- PostgreSQL 13+
- golang-migrate
- docker 


Environment Configuration

Copy .env.example to .env and populate

Configure database connection

A `makefile` has been provided to make (pun intended) things easier to setup. 
Feel free to replace the credentials in the make file to suit your needs

To create a new postgres container with credentials
```shell
    make postgres
```


To Create the database
```shell
    make createdb
```

To run UP migrations
```shell
    make migrateup
```
N/B the above step requires you to have golang-migrate installed

Populate your db with testdata
```shell
    make test
```
Run Application
```shell
    cd cmd && go run main.go
```

### Testing

- Integration tests for database interactions (the tests also double as a way to quickly insert data into the DB)
- API endpoint testing


### Performance Considerations

Connection pooling with pgx
Prepared statements
Efficient query design

### TODO
- pdf generation
- cron jobs for reminders
- caching for frequently fetched data (drafts)
- extensive data validation
- payment confirmation