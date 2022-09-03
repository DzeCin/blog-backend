FROM alpine:3.16.2 as certs
RUN apk add -U --no-cache ca-certificates

FROM golang:1.18.5-alpine AS build
WORKDIR /go/src/blog-backend
COPY . .

ENV CGO_ENABLED=0

RUN go build -a -installsuffix cgo -o blog-backend .

FROM scratch AS runtime
COPY --from=build /go/src/blog-backend ./
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080/tcp
USER 1000
ENTRYPOINT ["./blog-backend"]
