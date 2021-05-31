## We specify the base image we need for our
## go application
FROM golang:1.12.0-alpine3.9
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /gameapp
## We copy everything in the root directory
## into or /app directory
ADD . /gameapp
## We specify that we now wish to execute 
## any further commands inside our /app
## director
RUN chmod a+x /gameapp
WORKDIR /gameapp
## we run go build to compile the binary
## executable of our Go program
RUN go build -o app .
## Our start command which kicks off
## our newly created binary executable
EXPOSE 9000
CMD ["/gameapp/app"]
RUN chmod a+x /gameapp/app
