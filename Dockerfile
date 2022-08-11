FROM golang:1.19.0-alpine3.15 as builder
ARG GITHUB_TOKEN

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

RUN git config --global http.extraheader "PRIVATE-TOKEN: ${GITHUB_TOKEN}"
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/noneexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $GOPATH/src/github.com/yolcu360/app-versions

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app-versions

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /go/bin/app-versions /go/bin/app-versions

USER appuser:appuser

EXPOSE 9091

ENTRYPOINT ["/go/bin/app-versions"]