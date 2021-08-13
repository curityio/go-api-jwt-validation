FROM golang:1.16.6-alpine AS GO_BUILD
COPY . /server
WORKDIR /server/api
RUN go mod download github.com/gorilla/mux
RUN go get github.com/gbrlsnchs/jwt/v3
RUN go get github.com/joho/godotenv
RUN go build -o /go/bin/server/server

FROM alpine:latest
WORKDIR app
COPY --from=GO_BUILD /go/bin/server/ ./
COPY /api/records.json /app/records.json
COPY /api/.env /app/.env
EXPOSE 8080
CMD ./server 
