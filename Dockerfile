FROM golang:1.20-buster as builder
WORKDIR "/builder"
COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -mod=vendor -o appp ./

FROM debian:stable-slim

WORKDIR /app/
RUN mkdir docs

COPY --from=builder /builder/scripts  ./scripts
COPY --from=builder /builder/appp  .
COPY --from=builder /builder/.env  .
COPY --from=builder /builder/config.yml  .
CMD ["./appp"]
