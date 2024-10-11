# Use the official Node.js image as a base
FROM node:18-alpine AS build

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# RUN rm -rf .env

# Build the React application
RUN npm run build

# Use a lightweight web server to serve the build
FROM nginx:stable-alpine

COPY --from=build /app/build /usr/share/nginx/html/todo
#COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]

# Todo: Make it a simple docker file for react, and test it with ingress into k8s
