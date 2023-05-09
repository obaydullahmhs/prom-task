FROM golang:alpine AS builder
RUN mkdir /prom-task
COPY . /prom-task
WORKDIR /prom-task
RUN go build .

FROM alpine
WORKDIR /prom-task
COPY --from=builder /prom-task/ /prom-task/

ENTRYPOINT ["./prom-task", "server"]