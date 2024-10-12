# Use an official Node.js image
FROM node:14

# Set the working directory
WORKDIR /app

# Copy package.json and install dependencies

RUN npm install

# Copy the rest of the application code
COPY . .

# Expose the port
EXPOSE 3000

# Start the application
CMD ["node", "app.js"]