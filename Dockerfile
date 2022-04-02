FROM alpine:latest
ARG config
COPY ./logs /data/teacupapi/logs
COPY ./config /data/teacupapi/config
COPY ./data /data/teacupapi/data
COPY ./teacupapi /data/teacupapi/teacupapi
WORKDIR /data/teacupapi
VOLUME ["/data/teacupapi/logs"]
VOLUME ["/data/teacupapi/data"]

RUN apk update \
    && apk upgrade \
    && apk add --no-cache  \
    && apk add ca-certificates \
    && apk add update-ca-certificates 2>/dev/null || true \
    && apk add tzdata \
    && ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && mv ${config} ./config/active.ini

CMD ["/data/teacupapi/teacupapi", "--config=./config/active.ini"]