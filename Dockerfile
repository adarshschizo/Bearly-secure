
# Build stage
FROM golang:1.27.0-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /out/bearly-secure ./cmd/server
RUN CGO_ENABLED=0 go build -o /out/bearly-attacker-lab ./cmd/attackerlab


# Runtime stage
FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && adduser -s -g bearly bearly -D

WORKDIR /app

COPY --from=build /out/bearly-secure ./bearly-secure
COPY --from=build /out/bearly-attacker-lab ./bearly-attacker-lab
COPY --from=build /src/attacker-lab ./attacker-lab
COPY --from=build /src/web ./web
COPY --from=build /src/data/fixtures ./data/fixtures

RUN mkdir -p /app/data \
    && chown -R bearly:bearly ./data

USER bearly

CMD ["./bearly-secure"]