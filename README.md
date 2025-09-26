# Project Builder API

* Dependences

- Gin (HTTP server)
- Gorm (golang ORM)
- postgres driver (database)
- uuid
- godotenv (environment variables)

- make
- go
- postgres

## Install & Setup

* Setup your `.env` vars

> tip: use `openssl rand -base64 32` to generate a strong `JWT_SECRET`


* Setup your postgres database

```sql

  CREATE DATABASE YOUR_DATABASE_NAME;

  CREATE TYPE user_role AS ENUM ('USER', 'ADMIN');

```

* Run

```sh

  make build

  make run

```

## Run Tests

* Unitary Tests

```sh

  make test

```

* Stress Tests (JMeter)

  * Run `cd jmeter`:

  ```sh

    jmeter -n -t <TEST-FILE> -l results.jtl -e -o ./report

  ```

  * Cleanup `cd jmeter`:

  ```sh

    rm -rf report *.jtl *.log

  ```
