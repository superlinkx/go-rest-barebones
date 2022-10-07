# argument for Go version
ARG GO_VERSION=1.19.1

# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build

RUN apk add --no-cache git
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
# Build the executable
RUN CGO_ENABLED=0 go build \
  -installsuffix 'static' \
  -o /app \
  /src/cmd/go-rest-barebones/

# STAGE 2: build the container to run
FROM gcr.io/distroless/static AS final

ENV APP_HOSTNAME=localhost
ENV APP_PORT=8080
ENV PSQL_HOST=localhost
ENV PSQL_PORT=5432
ENV PSQL_DB=app
ENV PSQL_USER=app
ENV PSQL_PASSWORD=!ChangeMe!
ENV PSQL_SSLMODE=default

USER nonroot:nonroot

# copy compiled app
COPY --from=build --chown=nonroot:nonroot /app /app
EXPOSE 8080

# run binary
ENTRYPOINT ["/app"]