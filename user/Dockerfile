FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR /app/

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
#COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/app .


# Run the binary program produced by `go install`
CMD ["./out/app"]