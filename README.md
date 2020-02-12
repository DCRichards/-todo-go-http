# todo-go-http

> An example Go HTTP microservice.

## Setup

**Docker**

```bash
docker-compose build
docker-compose up
```

**No Docker**

1. Spin up an instance of Postgres.
2. Set up your environment variables (see below).
3. Run the following:

    ```bash
    go get -u github.com/cosmtrek/air
    go mod download
    air -c .air.conf
    ```

## Development

#### Live Reload

This project [cosmtrek/air](https://github.com/cosmtrek/air) for live reloading.

Config file: `.air.conf`.


## Testing

Run the tests in the usual way.

```bash
go test ./...
```
Ensure you have `CGO_ENABLED=0` in your env vars to remove the warning (some packages cause this, such as `net/http`). Test utilities can be found in the `internal/testutils` package.

## Configuration

For Docker development, environment variables are managed through the `env_file` option of `docker-compose.yml`. To configure, simply create a file `.env` in the project root. For example:

```env
PORT=8080
POSTGRES_HOST=db
POSTGRES_PORT=5432
POSTGRES_DB=dev
POSTGRES_USER=dev
POSTGRES_PASSWORD=mysecurepass
```

## Motivations

This project aims to explore the standards, conventions and general best practices from the Golang community, as well as my own implementation. In my research I was fortunate to come across a number of great speakers and resources, which include, but are not limited to:

* [🍿 Kat Zien - How Do You Structure Your Go Apps](https://www.youtube.com/watch?v=oL6JBUk6tj0)
* [🍿 Mat Ryer - How I Write HTTP Web Services after Eight Years](https://www.youtube.com/watch?v=rWBSMsLG8po)
* [📘 Dave Cheney: Practical Go](https://dave.cheney.net/practical-go)
