# api.kig.land

## Description

Power by go, deploy on docker, a tmpl backend.

## Deps

```shell
go mod download
```

## Env Vars

```shell
PORT=8080 DATABSE_URL=xxx go run .
```

## Build

```shell
docker buildx build --platform linux/amd64 -t image_name:tag --load .
```

## Docker Run

```shell
docker run --restart unless-stopped -d -p 9321:9321 -e PORT=9321 -e DATABASE_URL="xxx" image_name:tag
```