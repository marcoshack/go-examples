FROM golang:1.23

# Install dependencies
RUN apt-get update && apt-get install -y wget unzip tmux

# Install protoc
RUN cd /tmp && \
    wget https://github.com/protocolbuffers/protobuf/releases/download/v29.2/protoc-29.2-linux-x86_64.zip && \
    unzip protoc-29.2-linux-x86_64.zip bin/protoc -d /usr/local && \
    rm protoc-29.2-linux-x86_64.zip

# Install protoc-gen-go and protoc-gen-twirp
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@latest

# Set PATH
ENV PATH="/root/go/bin:${PATH}"

# Set working directory
WORKDIR /twirp

# Copy the rest of your application code
COPY . .

# Build the application
RUN go mod tidy
RUN make
