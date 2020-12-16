FROM golang:1.13-alpine as builder
ARG VERSION=0.0.1

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# build
WORKDIR /go/src/qis-crawler-queue-api
COPY go.* ./
RUN GO111MODULE=on go mod download
COPY . .
RUN go build -o _output/bin/qis-crawler-queue-api .

# runtime image
FROM scratch
COPY --from=builder /go/src/qis-crawler-queue-api/_output/bin/qis-crawler-queue-api /opt/qis-crawler-queue-api
EXPOSE 80
ENTRYPOINT ["opt/qis-crawler-queue-api"]