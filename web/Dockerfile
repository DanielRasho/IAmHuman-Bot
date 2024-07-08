# Stage 1: Build Image
FROM node:22.3.0 as build

WORKDIR /web
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Use the official Nginx image from the Docker Hub
FROM nginx:latest

# Copy your Nginx configuration file
COPY nginx.conf /etc/nginx/conf.d/nginx.conf

# Copy your static website files
COPY --from=build /web/dist/ /usr/share/nginx/html

# Expose port 80
EXPOSE 80