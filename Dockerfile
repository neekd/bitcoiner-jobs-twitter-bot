FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o bot .

FROM amazon/aws-lambda-go
ADD config.yaml /app/
COPY --from=builder /build/bot /app/
WORKDIR /app
CMD ["./bot"]