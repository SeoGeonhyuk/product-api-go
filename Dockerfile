FROM golang:latest



# Create application directory
RUN mkdir /app



# Copy the product-api binary
COPY . .
RUN go install .


# Copy the configuration file


# Set entrypoint
ENTRYPOINT [ "go", "run", "main.go" ]