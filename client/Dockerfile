# Initialise container to install packages & build application
FROM node:lts
WORKDIR /app
COPY package*.json ./
RUN npm install

# Copy source code and build application
COPY . .
RUN npm run build

# Set command to start the application
CMD [ "npm", "start" ]

# Metadata
LABEL author="Jim Kelly"