FROM alpine:latest

LABEL Company="DatNguyen."
LABEL Author="Dat Nguyen <it.tandat@gmail.com"

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

COPY dist/tokoin_test /bin/
RUN chmod +x /bin/tokoin_test
RUN mkdir -p /tokoin_test
WORKDIR /tokoin_test

VOLUME /tokoin_test
EXPOSE 8181

CMD [ "tokoin_test" ]
