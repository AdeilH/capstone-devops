FROM golang:1.14

## Step 1:
# Create a working directory
WORKDIR /app


## Step 2:
# Copy source code to working directory
COPY . main.go /app/
COPY . template.html /app/
COPY Godeps /app/Godeps

## Step 3:
RUN go get github.com/tools/godep
RUN go build main.go

## Step 4:
# Expose port 80
EXPOSE 8000

## Step 5:
# Run webapp at container launch
CMD ["/app/main"]

