ARG PORT=9000

########################################
# Stage 1: builder
########################################
FROM golang:1.19.9-alpine3.17 as builder

# set destination for COPY
WORKDIR /src

# download go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy the source code
COPY . ./

# build the API
RUN CGO_ENABLED=0 go build -v ./cmd/gaexporter/gaexporterd.go


#########################################
## Stage 2: runner
#########################################
FROM scratch as runner

## copy just the compiled binary
COPY --from=builder /src/gaexporterd /opt/

# configure host and port for the API
ARG PORT
ENV HOST=0.0.0.0
ENV PORT=$PORT

CMD ["/opt/gaexporterd"]
