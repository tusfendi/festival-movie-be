# Festival Movie Backend Service

This backend service is for Festival movie apps

## Contact
| Name                   | Email                           |
| :--------------------: |:-------------------------------:|
| Tusfendi               | tusfendi@gmail.com              |

## Onboarding and Development Guide

### Documentations
- [Api Docs - Postman](https://documenter.getpostman.com/view/7094340/2s93XyShz7)
### Prerequisite
- Git (See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git))
- Go 1.20 or later (See [Golang Installation](https://golang.org/doc/install))
- MySQL (See [MySQL Installation](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/))


### Installation
- Clone this repo

    ```sh
        git clone https://github.com/tusfendi/festival-movie-be.git
    ```

- Copy `.env.example` to `.env`

    ```sh
        cp .env.example .env
    ```
- Setup local database, you can import database from database/festival_movie_db.sql

- Start service API
    ```sh
        go run cmd/api/main.go
    ```