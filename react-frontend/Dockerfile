# Use the official Node image
FROM node:20

# Set working directory inside the container
WORKDIR /app

# Copy only package.json and package-lock.json for dependency installation
COPY package.json package-lock.json ./

# Install dependencies using npm
RUN npm install

# Copy the rest of the application files
COPY . .

# Expose the React development server's default port
EXPOSE 3000

# Run the React development server
CMD ["npm", "start"]
