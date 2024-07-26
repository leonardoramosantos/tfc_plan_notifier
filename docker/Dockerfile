# Build the application binary
FROM golang:alpine AS builder

WORKDIR /app
COPY src /app

RUN go get -d -v
RUN go build -o /tfc_plan_notifier

# Create a smaller image with only the generated binary
FROM scratch
COPY --from=builder /tfc_plan_notifier /tfc_plan_notifier
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/tfc_plan_notifier"]
