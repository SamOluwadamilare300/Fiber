# File Upload Service

## Overview

This is a Go and the Fiber web framework tutorial. It provides a basic undedrstanding to fiber. Reading the documentations helps in developing thi tutorial. https://docs.gofiber.io

## Features

- Accepts multiple file uploads in a single request
- Generates unique filenames using timestamps
- Provides detailed error responses
- Uses the high-performance Fiber web framework

## Requirements

- Go 1.16 or higher
- [Fiber v2](https://github.com/gofiber/fiber)

## Installation

1. Clone this repository
   ```bash
   git clone https://github.com/username/file-upload-service.git
   cd file-upload-service
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Build the application
   ```bash
   go build -o upload-service
   ```

4. Run the service
   ```bash
   ./upload-service
   ```

The server will start on port 3000.
