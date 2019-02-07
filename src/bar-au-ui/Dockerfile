FROM node:11.9-alpine as builder

WORKDIR /app

COPY . .

# Apparently python is needed by the Ninja build system
RUN apk add --update \
    build-base \
    python

# Install dependencies
RUN npm install

# Build release
RUN npm run build:prod

# Build our tiny custom http server
FROM golang:alpine as gobuilder

COPY ./serve /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o serve .

RUN ls -al .

FROM scratch

# Copy website content
COPY  --from=builder /app/build/ /www/

# Copy http server and make it executable
COPY --from=gobuilder /app/serve /

EXPOSE 8080

ENTRYPOINT [ "/serve" ]
