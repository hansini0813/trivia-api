FROM golang:1.21.0
WORKDIR  /usr/src/app

# so our app refreshes 
RUN go install github.com/air-verse/air@latest

# to start the app form the host machine add necessary dependenices 
COPY . . 
RUN go mod tidy 