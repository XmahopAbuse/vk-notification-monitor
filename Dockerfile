FROM golang:1.20-alpine as backend-builder

WORKDIR /app/backend

COPY ./backend . 

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM node:20-alpine3.17 as frontend-builder

WORKDIR /app/frontend

COPY ./frontend . 
RUN npm install --only=production
WORKDIR /app/frontend

RUN npm run build 

FROM alpine:3

WORKDIR /app

COPY --from=backend-builder /app/backend/main /app
COPY --from=frontend-builder /app/frontend/build/ /app/frontend/


CMD ["./main"]