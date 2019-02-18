# Build our tiny custom http server
FROM golang:alpine as gobuilder

COPY ./serve /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o serve .

# Build ui front-end
FROM node:11.9-alpine as builder

WORKDIR /app

# Apparently python is needed by the Ninja build system
RUN apk add --update \
    build-base \
    python

# Copy only package.json to use build cache
COPY ./package.json ./package-lock.json ./

# Install dependencies
RUN npm install

# Copy everythings now
COPY . .

# Build release
RUN npm run build:prod

FROM scratch

# Copy website content
COPY  --from=builder /app/build/ /www/

# Copy http server and make it executable
COPY --from=gobuilder /app/serve /

EXPOSE 8080

ENTRYPOINT [ "/serve" ]
