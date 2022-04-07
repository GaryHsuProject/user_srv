# build stage
FROM golang:alpine AS build-stage 
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# final stage
FROM alpine 
WORKDIR /app 
COPY --from=build-stage /app /app 
CMD [ "./main" ]
EXPOSE 8001


