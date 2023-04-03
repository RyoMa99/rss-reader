FROM golang:1.20 as build

WORKDIR /app/src

COPY ./src/go.mod ./
COPY ./src/go.sum ./
ENV CGO_ENABLED=0
RUN go mod download -x

COPY ./src .
RUN go build -o /feeder .

FROM scratch AS runtime

USER 1000
WORKDIR /app

ENV CONFIG_FILE=/config/config.yaml

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /feeder .

ENTRYPOINT [ "./feeder" ]
