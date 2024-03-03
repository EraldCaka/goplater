FROM golang:latest

WORKDIR /api

COPY . .

RUN make tidy

RUN make build

EXPOSE 5000

CMD make run
