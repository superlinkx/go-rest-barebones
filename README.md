# Go Rest Barebones Example

This is a simple, barebones REST example, a part of my series of barebones REST examples for research purposes.

This is by no means a recommended way to build a CRUD app. It is for research purposes only and has
several already known poor security. May be something I'll look at fixing in the future, but
isn't part of my current research. It's a quick and dirty implementation designed to get an MVP.

## Getting Started

I've created this repo with a VS Code Devcontainer definition. It's relatively easy to see the dependencies this way,
and if you're using VS Code, you can simply run this project in a dev container with the VS Code Remote Dev Containers
extension. This will create a reproducible environment with all dependencies already configured.

### Starting The App

To get the app's dependencies, run:

```bash session
foo@bar:~$ go mod download
```

In order to get the default database up and running, use:
```bash session
foo@bar:~$ docker-compose up database -d
```
(this works in the devcontainer thanks to
docker-in-docker). This will run Postgres with the default parameters and forward port 5432.

To handle migrations, you'll need to install `sql-migrate` if not using the dev container:

```bash session
foo@bar:~$ go install github.com/rubenv/sql-migrate/...@latest
```

The following commands will need to be run when working with a fresh database:

```bash session
# Have to read .env file because sql-migrate doesn't support .env out of the box :(
# Migrate
foo@bar:~$ export $(< .env xargs); sql-migrate up
# Seed
foo@bar:~$ go run ./cmd/faker
```

Start the app:

```bash session
foo@bar:~$ go run ./cmd/go-rest-barebones
```

- Default server: localhost:8080
- Can either run Postman against the endpoints or use the `example.http` with the VS Code "Rest Client" extension to try out some example endpoints.

## Changing the queries/schema

Schema and raw queries are in the `sql` directory split into separate subfolders. These are used to generate the actual
models and strongly typed queries in the app through `sqlc`. To regenerate, you'll need to install sqlc and then run the
generate command:

```bash session
# Only to install sqlc when you don't have it already
foo@bar:~$ go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
# Generate any time you change the schema or queries files
foo@bar:~$ sqlc generate
```

## Running as a production-style test

The point of this repo is to understand some of the dynamics at play when deploying as a "production-ready" app.
To this end, the `docker-compose.yml` defines three services:  
- database: Postgres database used for both production and development
- go-rest-barebones: Runs our Go app inside a distroless container
- web: runs an Nginx reverse proxy to our Go app. It's generally considered good practice to use some kind of reverse
  proxy when running a Go web server. It's possible to go without, but you do need to understand the implications and write
  more code and error handling to make sure it's up to spec.

To test out the "production-ready" environment, simply run:
```bash session
foo@bar:~$ docker-compose up -d
```

By default, all ports are forwarded to localhost in the following way:
- Postgres Port: 5432
- Direct Go App Port: 8080
- Web Port: 8888

They should also be autoforwarded out of the dev container if using VS Code