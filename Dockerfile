FROM golang:1.18 AS build
WORKDIR /pxier_web
COPY . .
RUN go mod tidy &&  \
    go mod vendor && \
    go build -o pxier_web && \
    cp config.example.yaml config.yaml

FROM ubuntu:22.04 AS run
COPY --from=build /pxier_web/pxier_web .
COPY --from=build /pxier_web/config.yaml .
CMD ["./pxier_web"]