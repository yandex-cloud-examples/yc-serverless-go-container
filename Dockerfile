FROM golang:1.16-alpine AS build

WORKDIR /app

COPY app/* ./

RUN go mod download
RUN go build -o /server

FROM alpine

COPY --from=build /server /server
CMD ["/server"]