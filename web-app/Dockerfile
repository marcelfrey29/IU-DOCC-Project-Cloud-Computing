# Builder Image
FROM node:22.9.0-bookworm AS builder
WORKDIR /app
COPY package.json package-lock.json /app/
RUN npm ci
COPY . /app/
RUN npm run build

# Container
FROM nginx:1.27.2-bookworm
COPY --from=builder /app/dist /usr/share/nginx/html
