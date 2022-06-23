# Start from golang base image we can use "FROM scratch"
FROM golang:alpine

RUN mkdir /app
WORKDIR /app

# Cory from the source provided in the command to the WD in the Container
COPY . .

# Download all the dependencies(no installation)
RUN go get -d -v ./...

# So we need to install the packages
RUN go install -v ./...

# Move to the cmd file then build
RUN cd cmd && go build -o /server

# Expose port 8082 to the outside
EXPOSE 8082

# Run the executable in the container(normaly it is on port 8082 "hardCoded")
CMD [ "/server" ]