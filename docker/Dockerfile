# NODE JS Frontend Image
FROM node:22.3-alpine3.20 AS frontend


# Create non root user for scratch
RUN adduser --disabled-password -u 10001 appuser

# Cache frontend node_modules
WORKDIR /tmp/
COPY ./frontend/package.json /tmp/package.json
RUN yarn install
WORKDIR /app/frontend/
RUN cp -a /tmp/node_modules /app/frontend/

# Build frotend
WORKDIR /app/
COPY . .
WORKDIR /app/frontend/
RUN yarn build

# GOLANG Server Image
FROM golang:1.22-alpine3.20 AS server
WORKDIR /app/
COPY . .
RUN go mod tidy
WORKDIR /app/server/
RUN GOOS=linux go build -ldflags="-s" -o "../build/server" -v

# TODO: include svgcleaner as dependency including pull request #210 and implmentation for style="default-color"

FROM ki365-gltfpack AS gltfpack

FROM ki365-git AS git

# SCRATCH Image for Production
FROM scratch AS prod

# All image metadata
LABEL org.opencontainers.image.source https://github.com/Ki365/Ki365

# Copy all program data to final image
COPY --from=frontend /etc/passwd /etc/passwd
USER appuser
COPY --from=git /libexec/git-core/git-http-backend /bin/.
COPY --from=git /libexec/git-core/git-http-backend.dbg /bin/.
COPY --from=gltfpack / /bin/.
COPY --from=frontend /app/build/ .
COPY --from=server /app/build/ .
ENTRYPOINT ["/server", "-p", "8080", "-d", "/frontend"]