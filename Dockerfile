FROM golang:1.14-alpine3.12 as builder

RUN mkdir -p /build/src

ENV GOPATH=/build
ENV GOBIN=/usr/local/go/bin

ENV GO111MODULE=on
WORKDIR $GOPATH/src

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o out . \
	&& cp out $GOPATH/.

CMD [ "/build/out" ]

FROM alpine:3.13.3
COPY --from=builder /build/out .
COPY loadserver.sh /usr/bin/
RUN apk add --no-cache stress-ng && \
	chmod +x /usr/bin/loadserver.sh
CMD [ "./out" ]