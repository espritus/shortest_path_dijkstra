# Use a lightweight NGINX image
FROM nginx:alpine

# Copy the static files to the NGINX document root
COPY ./ /usr/share/nginx/html

# Expose the desired port (if needed)
EXPOSE 80

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]
