FROM golang:1.18-alpine3.15 as build
WORKDIR /app
COPY . .
RUN go build

FROM alpine:3.15
WORKDIR /app
COPY --from=build /app/rest .
ENTRYPOINT [ "./rest" ]