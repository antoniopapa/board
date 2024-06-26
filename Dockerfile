FROM golang:1.22-alpine as build

ENV TIER=starter

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o=./fed-dashboard ./cmd/web

FROM gcr.io/distroless/static-debian11

COPY --from=build /build/fed-dashboard .
COPY --from=build /build/ui/ /ui/
EXPOSE 8080
ENTRYPOINT ["/fed-dashboard"]