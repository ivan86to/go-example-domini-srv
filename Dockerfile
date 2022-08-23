FROM golang:1.17-bullseye as go-build

WORKDIR /app

COPY . .

RUN addgroup --system panda && adduser --system --shell /sbin/nologin --home /var/cache/panda  --ingroup panda panda 
RUN chown panda:panda /app
RUN go build -o app/app-srv .

#FROM gcr.io/distroless/base-debian11:latest
FROM golang:1.17-bullseye

COPY --from=go-build /etc/passwd /etc/group /etc/shadow /etc/
COPY --from=go-build /app /app

EXPOSE 8080

WORKDIR /app

USER panda

CMD [ "./app/app-srv"]