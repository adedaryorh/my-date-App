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

# Import the compiled executable.
COPY --from=builder /app/bin/celebut-api /app/

EXPOSE 8080

CMD ["/app/celebut-api"]
