# Start with a base image
FROM golang:1.16-alpine

# Set the working directory
RUN mkdir /build
WORKDIR /build

# Copy the source code to the container
RUN export GO111Module=on
RUN go get github.com/verolie/test-skyshi
RUN cd /build && git clone https://github.com/verolie/test-skyshi.git

# Build the Go application
RUN go /build/test-skyshi/main && go build

# Expose the port on which the application will listen
EXPOSE 8080

# Run the application
CMD ["/build/test-skyshi"]