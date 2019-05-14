FROM  golang:1.12.5-alpine3.9
LABEL maintainer="Nick Pocock <pocockn@hotmail.co.uk>"

ENV GO111MODULE=on

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .

ENTRYPOINT ["/app/database-migrator"]
