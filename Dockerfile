FROM golang:1.15.7-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o ./bin ./cmd/pim-service

FROM alpine:3.13
LABEL maintainer="Vladimir Chekholin chekholin@outlook.com"
RUN apk add ca-certificates
COPY --from=build_base /app /app
EXPOSE 7070
CMD ["/app/bin"]