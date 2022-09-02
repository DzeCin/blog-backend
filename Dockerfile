FROM golang:1.18.5-alpine AS build
WORKDIR /go/src/blog-backend
COPY . .

ENV CGO_ENABLED=0

RUN go build -a -installsuffix cgo -o blog-backend .

FROM scratch AS runtime
COPY --from=build /go/src/blog-backend ./
EXPOSE 8080/tcp
ENTRYPOINT ["./blog-backend"]
