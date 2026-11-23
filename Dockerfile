FROM --platform=$BUILDPLATFORM golang:alpine AS build
WORKDIR /build
COPY . .
RUN go mod download
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /tftp .

FROM scratch
COPY --from=build /tftp /tftp
ENTRYPOINT ["/tftp"]
