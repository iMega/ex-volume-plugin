FROM golang:1.8.3-alpine as builder
ARG CWD
WORKDIR $CWD
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o rel/app

FROM alpine:3.6
ARG CWD
WORKDIR /app
EXPOSE 80
COPY --from=builder $CWD/rel/app ./
CMD ["./app"]
