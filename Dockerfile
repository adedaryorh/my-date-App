# First stage: build the executable.
FROM golang:1.18.0-buster as builder

RUN apt-get -y update \
    && apt-get -y upgrade \
    && apt-get -y install git bash make

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Import the code from the context.
COPY . ./

# Build the statically-linked executable looking into '/vendor' folder
RUN make build-static

# Final stage: the running container.
FROM debian:buster-slim

# Import the Certificate-Authority certificates for enabling HTTPS.
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN apt update
RUN apt install -y curl

RUN curl -O https://dl.google.com/go/go1.18.4.linux-amd64.tar.gz
RUN tar xvf go1.18.4.linux-amd64.tar.gz

RUN chown -R root:root ./go
RUN mv go /usr/local

# Import the compiled executable.
COPY --from=builder /app/bin/celebut-api /app/
COPY --from=builder /app/migrations /app/migrations/

EXPOSE 8080

CMD ["/app/celebut-api"]
