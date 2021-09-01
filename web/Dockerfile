FROM alpine:latest

ENV DATABASE_URL="${DATABASE_URL}"
ENV DATABASE_USER="${DATABASE_USER}"
ENV DATABASE_PASSWORD="${DATABASE_PASSWORD}"
ENV DATABASE_NAME="${DATABASE_NAME}"
ENV DATABASE_PORT="${DATABASE_PORT}"
ENV RAPIDAPI_KEY="${RAPIDAPI_KEY}"

WORKDIR /app

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

ADD wtmWeb .

COPY views/ /app/views/

CMD /wait && /app/wtmWeb
