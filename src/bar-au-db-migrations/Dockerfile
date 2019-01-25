FROM alpine:3.8 as downloader

ENV VERSION=4.2.2

RUN apk add --no-cache curl tar

ENV DOWNLOAD_URL =  https://github.com/golang-migrate/migrate/releases/download/v$VERSION/migrate.linux-amd64.tar.gz

RUN curl -L $DOWNLOAD_URL | tar -xvz -C /tmp/

FROM alpine:3.8

COPY --from=downloader /tmp/migrate.linux-amd64 /usr/local/bin/migrate 

ENTRYPOINT ["/migrate"]