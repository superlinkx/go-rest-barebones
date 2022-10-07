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

```console
foo@bar:~$ go mod download
```

In order to get the default database up and running, use `docker-compose up database -d` (this works in the devcontainer thanks to
docker-in-docker). This will run Postgres with the default parameters and forward port 5432.

The following commands will need to be run when working with a fresh database:

```console
foo@bar:~$ <add command for migrating db>
foo@bar:~$ <add command for faking db data>
```

<!--This will run the current migrations and seed the database with Faker data.-->

Start the app:

```
foo@bar:~$ go run .
```

- Default server: localhost:8080
- Can either run Postman against the endpoints or use the `example.http` with the VS Code "Rest Client" extension to try out some example endpoints.
