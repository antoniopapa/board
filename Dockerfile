FROM golang:1.22-alpine as build

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o=./fed-dashboard ./cmd/web

FROM gcr.io/distroless/static-debian11

COPY --from=build /build/ .
COPY --from=build /build/ui/ /ui/
EXPOSE 8080
ENTRYPOINT ["/fed-dashboard"]