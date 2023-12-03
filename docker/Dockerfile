FROM alpine:3.18

ARG APP_NAME="app_api"


WORKDIR /app

RUN apk -U --no-cache add tzdata tini \
    && mkdir -p /app/storage/logs \
    && chown 1000:1000 -R /app \
    && rm -rf /root /tmp/* /var/cache/apk/* && mkdir /root

COPY ./build/${APP_NAME} /app/${APP_NAME}


# Set TimeZone
ENV TZ=Asia/Bangkok
ENV ENVOLOPMENT=production
ENV APP_NAME="${APP_NAME}"
ENV FB_PORT=8888

EXPOSE 8888



USER "1000:1000"

ENTRYPOINT ["tini", "--"]
CMD /app/$APP_NAME serv
