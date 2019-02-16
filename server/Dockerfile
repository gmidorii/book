ARG GO_VERSION=1.11.5

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
  echo 'gopher:x:65534:65534:gopher:/:' > /user/passwd && \
  echo 'gopher:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git make

WORKDIR /tool

RUN git clone https://github.com/google/wire.git
RUN cd ./wire/cmd/wire && \
  go build -o wire_bin ./
ENV WIRE_BIN=/tool/wire/cmd/wire/wire_bin

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

ENV CGO_ENABLED=0

RUN make build WIRE_BIN=${WIRE_BIN} BIN=/app

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app

EXPOSE 8080

USER gopher:gopher

ENTRYPOINT [ "/app" ]
