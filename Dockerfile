FROM golang:alpine

RUN apk add postgresql-client

WORKDIR $GOPATH/src/oloapi

COPY ./api .

RUN go mod download

RUN go build -o oloapi .
RUN mv ./oloapi /home
RUN mv ./wait_for_postgres.sh /home

EXPOSE 3000

CMD ["/home/wait_for_postgres.sh", "postgresdb", "/home/oloapi"]
