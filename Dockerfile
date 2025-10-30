# Start from the lightweight Go image
FROM golang:1.25-alpine

# Install git and any necessary build tools
RUN apk add --no-cache git build-base

# Set working directory
WORKDIR /app

CMD ["air", "-c", ".air.toml"]