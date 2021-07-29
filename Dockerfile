FROM golang:alpine as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o urlshortener .

WORKDIR /dist

RUN cp /build/urlshortener .

FROM scratch

COPY --from=builder /dist/urlshortener /app/

WORKDIR /app


CMD ["./urlshortener"]