# go-events-api

A RESTful API for managing events and user registrations, built with Go and Gin. The API allows users to sign up, log in, create and manage events, and register or unregister for events. Authentication is handled via JWT tokens.

## Overview

This API provides endpoints for:

- User registration and authentication
- Creating, updating, retrieving, and deleting events
- Registering and unregistering for events

Some endpoints require authentication via a JWT token.

## Endpoints

### Authentication & Users

- `POST /sign-up`  
  Register a new user.  
  **Body:** `{ "email": string, "password": string }`  
  **Returns:** Success message or error.

- `POST /login`  
  Authenticate a user and receive a JWT token.  
  **Body:** `{ "email": string, "password": string }`  
  **Returns:** JWT token or error.

### Events

- `GET /events`  
  Retrieve a list of all events.  
  **Returns:** Array of event objects.

- `GET /events/:id`  
  Retrieve details for a specific event by ID.  
  **Returns:** Event object.

- `POST /events`  
  Create a new event. **(Requires authentication)**  
  **Body:** `{ "name": string, "description": string, "location": string, "date_time": string }`  
  **Returns:** Created event object.

- `PUT /events/:id`  
  Update an existing event. **(Requires authentication, must be event owner)**  
  **Body:** `{ "name": string, "description": string, "location": string, "date_time": string }`  
  **Returns:** Success message.

- `DELETE /events/:id`  
  Delete an event. **(Requires authentication, must be event owner)**  
  **Returns:** Success message.

### Event Registration

- `POST /events/:id/register`  
  Register the authenticated user for an event. **(Requires authentication)**  
  **Returns:** Success message.

- `DELETE /events/:id/register`  
  Cancel the authenticated user's registration for an event. **(Requires authentication)**  
  **Returns:** Success message.

## Models

### User

- `id`: int64
- `email`: string
- `password`: string (hashed)

### Event

- `id`: int64
- `name`: string
- `description`: string
- `location`: string
- `date_time`: string (ISO8601)
- `user_id`: int64 (owner)

## Getting Started

1. Clone the repository.
2. Install dependencies:  
   `go mod download`
3. Run the server:  
   `go run main.go`
4. Use an API client (like Postman) to interact with the endpoints.
