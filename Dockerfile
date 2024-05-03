FROM golang:1.21.1

ENV GOPATH=/

WORKDIR /app

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# copy SQL script for creating tables
COPY ./pkg/database/postgres/create_database.sql /docker-entrypoint-initdb.d/

# wait for db to initialize
RUN chmod +x scripts/wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o med-app  ./cmd/app/main.go

CMD [ "./med-app" ]
