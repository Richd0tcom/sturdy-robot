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

## Database Design
![Nume](https://github.com/user-attachments/assets/421a4038-83a0-4829-a7f8-1573c2e98301)



Normalized relational schema

Supports multiple currencies
Designed for scalability and data integrity

Database Migrations
Location: migrations/

Versioned SQL migration files
Supports forward and rollback migrations

Query Generation
Using SQLC for type-safe database interactions:

Automatically generates Go structs
Compile-time query validation
Reduces manual SQL boilerplate

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
- SQLC
- golang-migrate


Environment Configuration

Copy .env.example to .env and populate

Configure database connection

Run migrations (with golang-migrate installed already)

Database Migration
```bash
    make migrateup
```
migrate -path migrations -database "postgresql://user:pass@localhost/dbname" up
Generate SQLC Queries
bashCopysqlc generate
Run Application
bashCopygo run cmd/api/main.go


Development Workflow

Write SQL queries in db/query/
Generate structs with SQLC
Implement service logic
Create API handlers

Testing


Integration tests for database interactions (the tests also double as a way toquickly insert data into the DB)
API endpoint testing


Performance Considerations

Connection pooling with pgx
Prepared statements
Efficient query design

Security

Input validation
Prepared statements prevent SQL injection
JWT authentication (recommended)
