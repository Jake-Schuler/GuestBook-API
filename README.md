# Guestbook API

This is a simple API for a guestbook application. It uses SQLite as a database and Gorm as an ORM. The API is written in Go and uses the Echo web framework.

## Endpoints

### GET /messages

Returns all messages in the database.

### GET /totalMessages

Returns the total number of messages in the database.

### POST /newMessage

Creates a new message in the database.

### GET /uptime

Returns the uptime of the server since it started.

### Static Assets

Static files are served from the `assets` directory.