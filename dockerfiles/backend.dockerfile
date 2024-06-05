FROM golang:1.22.3-alpine as builder
WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-ldflags="-w -s" -o build/main .


# Build final image
FROM gcr.io/distroless/static-debian12
LABEL authors="docker@alecodev.com"

COPY --from=builder /app/build/main /

EXPOSE 8000

CMD [ "/main" ]