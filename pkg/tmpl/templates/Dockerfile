# tmpl/Dockerfile.tmpl
FROM go:1.17 as env
ENV GOPROXY https://goproxy.io,direct

FROM go:1.17 as builder
WORKDIR /go/src
ADD . .
RUN CGO_ENABLED=0 go build -o {{ .Name }} {{ .Workdir }}/main.go

FROM alpine:3.12 as runtime
COPY --from /go/src/demo {{ .Name }}
WORKDIR /root/workdir
ENTRYPOINT [ "/go/bin/{{ .Name }}" ]