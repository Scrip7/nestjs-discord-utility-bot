FROM golang:1.18-alpine as build

# Set the working directory
WORKDIR /go/src/app

# The cgo tool is enabled by default for native builds on systems where it is expected to work.
# It is disabled by default when cross-compiling
ENV CGO_ENABLED=0

# Controls the source of Go module downloads
# Can help assure builds are deterministic and secure.
ENV GOPROXY=https://proxy.golang.org

# Executable filename (binary file)
ENV APP_NAME=app

# Cache dependencies
COPY ["go.mod", "go.sum", "./"]

# Download dependencies
RUN ["go", "mod", "download"]

# Copy project files
COPY . .

# Build binary file
RUN ["go", "build", "-o", "build/${APP_NAME}"]

#
# Production build
#
FROM alpine:3.16.0 as prod

# file and iconv commands
RUN apk add --no-cache gnu-libiconv file

# By default, Docker runs container as root which inside of the container can pose as a security issue.
# RUN addgroup -S app && adduser -S -G app app
# USER app

# Set the working directory
WORKDIR /go/src/app

COPY --from=build /go/src/app/build/${APP_NAME} ./

# Execute the binary file
CMD ["./${APP_NAME}"]
