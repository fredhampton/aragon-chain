FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES make git curl build-base

# Set working directory for the build
WORKDIR /go/src/github.com/aragon/aragon-chain

# Install dependencies
RUN apk add --update $PACKAGES

# Add source files
COPY . .

# Make the binary
RUN make tools deps build

# Final image
FROM alpine

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/aragon/aragon-chain/build/arad /usr/bin/arad

# Run arad by default
CMD ["arad"]
