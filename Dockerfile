FROM golang:1.8 AS builder

RUN apt-get update && apt-get install -y unzip --no-install-recommends && \
    apt-get autoremove -y && apt-get clean -y && \
    wget -O dep.zip https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-amd64.zip && \
    echo '96c191251164b1404332793fb7d1e5d8de2641706b128bf8d65772363758f364  dep.zip' | sha256sum -c - && \
    unzip -d /usr/bin dep.zip && rm dep.zip

RUN mkdir -p /go/src/github.com/pocockn/mono-repo/database-migrator
WORKDIR /go/src/github.com/mono-repo/database-migrator

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o database-migrator

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/mono-repo/database-migrator  .
ENTRYPOINT ["./database-migrator"]
