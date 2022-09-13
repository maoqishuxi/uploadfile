# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY /public ./public
COPY /uploadfile ./uploadfile
COPY .gitignore ./gitignore
COPY Dockerfile ./Dockerfile
COPY *.go ./

RUN go build -o /uploadfile

EXPOSE 8000

CMD [ "/uploadfile" ]