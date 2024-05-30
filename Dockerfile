FROM --platform=linux/amd64 golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o main .
RUN apk update && apk add upx
RUN upx --best --lzma /app/main

FROM scratch
COPY --from=builder /app/main /main
ENV PORT 9000 
ENV DATABASE_URL "postgresql://dont_root:my_secret_pw@where.are.you.from/dont_forget_production?sslmode=require"
EXPOSE 9000
ENTRYPOINT ["/main"]