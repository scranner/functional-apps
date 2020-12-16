FROM golang:1.13.3 as builder

WORKDIR /go/src/promotion/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o _output/bin/promotion .

FROM scratch
COPY --from=builder /go/src/promotion/_output/bin/promotion /bin/
EXPOSE 80
ENTRYPOINT ["/bin/promotion"]