#######################################
###############  BUILD  ###############
#######################################

FROM golang:1.25 AS build-stage
WORKDIR /app
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build ./main.go

#######################################
##############  RUNTIME  ##############
#######################################

FROM gcr.io/distroless/static-debian12 AS release-stage
COPY --from=build-stage /app/main ./main
CMD ["/main"]
