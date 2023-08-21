# account-service

## Description

Spongebob merupakan nama repository dari account service

aturan clean arsitektur menurut uncle bob

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

Referensi : https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

#### Diagram:

![golang clean architecture](https://github.com/DueIt-Jasanya-Aturuang/spongebob/blob/rama/REFACTOR/testing-usecase/docs/arsitektur/clean-arsitektur.png)

## Directory

```
|-- delivery
|   |-- rest-api
|-- deploy
|   |-- Dockerfile
|   |-- docker-compose.yml
|   |-- deploy.sh
|-- docs
|   |-- arsitektur
|-- domain
|   |-- domain-name
|   |-- mocks
|-- infrastructures
|   |-- config
|   |-- db
|   |-- logs
|   |-- redis
|   |-- init.go
|-- internal
|   |-- helpers
|   |-- repositories
|   |-- usecase
|   |-- utils
|-- migrations
|-- tests
|   |-- integration
|   |   |-- utils
|   |-- unit
```
