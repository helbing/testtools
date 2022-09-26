FROM golang:alpine

WORKDIR /www
COPY . .
RUN go build -o app .
CMD [ "/www/app" ]
