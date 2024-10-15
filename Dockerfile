# tells docker to use the golang version 1.21.0 as the base image of the container.
# the base image of a container provides the basic setup for ur application
# why is this helpful? instead of setting up golang you for example installign go, configuring go, and do evrythign manually, you can use an official pre-configured base image that already has everything set up for you.
FROM golang:1.21.0
# Set the working directory to /usr/src/app
WORKDIR  /usr/src/app

# so our app refreshes 
RUN go install github.com/air-verse/air@latest

# to start the app form the host machine add necessary dependenices 
# Copy the current directory contents into /usr/src/app
COPY . . 
# go mod tidy is a command in the Go programming language’s module system. It ensures that the go.mod and go.sum files are in sync with the source code. It performs two main tasks:
# Removes Unused Dependencies: It removes any dependencies from the go.mod file that are no longer needed by the project.
# Adds Missing Dependencies: It adds any missing dependencies that are required by the source code but are not currently in the go.mod file.
# Purpose in Dockerfile: When you include RUN go mod tidy in your Dockerfile, you're telling Docker to clean up and ensure that the Go module dependencies are up to date and that your project only contains the necessary dependencies. This helps ensure that your application is lighter and cleaner when built inside the Docker container.
RUN go mod tidy 

# golang is a programming language used for backend services. for example to use api's. 
# Docker is a containerization platform that allows you to package applications and their dependencies into containers.
    # a container includes everyhtign eeded to run the application such as code, libraries and system tools, making it portable and consistent across environments
