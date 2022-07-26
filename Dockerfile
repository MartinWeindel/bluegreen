#############      builder       #############
FROM golang:1.18.3 AS builder

WORKDIR /build
COPY . .

RUN make release

############# base
FROM gcr.io/distroless/static-debian11:nonroot AS base

#############      bluegreen     #############
FROM base AS bluegreen
WORKDIR /

COPY --from=builder /build/bluegreen /bluegreen

ENTRYPOINT ["/bluegreen"]