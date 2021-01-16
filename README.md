# Byte Scraper API

The Byte Scraper API is an API with endpoints that return JSON data about users and bytes from [Byte](https://byte.co/) using web scraping. This API provides an alternative method of obtaining information from Byte since Byte's APIs are private. This project is unofficial and not affiliated with Byte Inc.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
   - [Object Models](#object-models)
   - [Response HTTP Status Codes](#response-http-status-codes)
   - [Endpoints](#endpoints)

## Installation

```shell
go get github.com/jjlock/byte-scraper-api
```

## Usage

To start the server:

```shell
go run main.go
```

You can also specify a port using the port flag (default port 8000):

```shell
go run main.go -port 8080
```

To gracefully shutdown the server, type Ctrl-C on the command line.

### Object Models

**User Object**

| Field             | Type          | Description                                                   |
| ----------------- | ------------- | ------------------------------------------------------------- |
| username          | string        | The Byte username of the user.                                |
| profile_image_url | string        | The source URL of the user's profile image.                   |
| bio               | string        | The about section of a user's profile.                        |
| recent_byte_ids   | array[string] | The byte IDs of the user's most recent bytes (max 10 IDs).    |
| recent_byte_urls  | array[string] | The byte URLs of the user's most recent bytes (max 10 bytes). |
| url               | string        | The URL of the user's profile.                                |

**Byte Object**

| Field         | Type          | Description                                                                                                                          |
| ------------- | ------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| id            | string        | The ID of the byte. This can be found at the end of a byte url.<br><br>Example: In `https://byte.co/@zachking/Lqs1Y8INqlt` the ID is `Lqs1Y8INqlt` |
| user          | string        | The username of the user who posted the byte.                                                                                        |
| user_url      | string        | The profile URL of the user who posted the byte.                                                                                     |
| thumbnail_url | string        | The source URL of the byte thumbnail.                                                                                                |
| caption       | string        | The caption of the byte.                                                                                                             |
| created_at    | string        | How long ago the byte was posted, or if more than a year old, the date of when the byte was posted.                                  |
| loops         | int           | The number of views of the byte.                                                                                                     |
| urls          | string[array] | The URLs for the byte.                                                                                                               |

**Error Object**

| Field   | Type   | Description                                                 |
| ------- | ------ | ----------------------------------------------------------- |
| status  | int    | The HTTP status code (also located in the response header). |
| message | string | The reason the error occured.                               |

### Response HTTP Status Codes

| HTTP Status Code | Description                                                                                                                             |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| 200              | OK - The request was successful. The client can read the response header and body.                                                      |
| 404              | Not Found - The resource cannot be found. See note below.                                                                               |
| 500              | Internal Server Error - The server encountered an error internally and could not handle the request.                                    |
| 503              | Service Unavailable - The Byte website responded with a 5xx error during scraping and the server temporarily cannot handle the request. |

*Note:* On a 404 response if a user was requested, the user either does not exist or does exist but has not posted any bytes. If a user has not posted any bytes, they cannot be found on the Byte website meaning it is not possible to scrape information on that user.

**Error Object Example**

`GET http://localhost:8000/api/users/username404`

```
{
  status: 404,
  message: "User not found. User either does not exist or does exist but has no bytes."
}
```

### Endpoints

Only GET requests are allowed and all responses will be returned as a JSON object.

#### Get a User

`GET http://localhost:8000/api/users/{username}`

**Request**

| Path Parameter | Required | Type   | Description                    |
| -------------- | -------- | ------ | ------------------------------ |
| `username`     | required | string | The Byte username of the user. |

**Response**

On success, the HTTP status code in the response header will be `200` OK and the response body will contain an user object. On error, the response HTTP status code will be in the response header and the response body will contain an error object.

Example:

`GET http://localhost:8000/api/users/zachking`

```
{
  username: "zachking",
  profile_image_url: "https://e6k9t9a9.stackpathcdn.com/avatars/PYH6U42FJFG3TICMHTZHTYYA44.jpg",
  bio: "Magic",
  recent_byte_ids: ["Lqs1Y8INqlt", "EF5TfBkZMzW", "IqVBw16uTdR", "2OJDBFJY4H7", "1Ubxgq6ZQ0t", "Cm35Fle2Za0", "FbnnYMceIRy", "9rj2orF5gQ", "LkkbbYE1DyU", "LsVmWhe51bT"],
  recent_byte_urls: ["https://byte.co/@zachking/Lqs1Y8INqlt", "https://byte.co/@zachking/EF5TfBkZMzW", "https://byte.co/@zachking/IqVBw16uTdR", "https://byte.co/@zachking/2OJDBFJY4H7", "https://byte.co/@zachking/1Ubxgq6ZQ0t", "https://byte.co/@zachking/Cm35Fle2Za0", "https://byte.co/@zachking/FbnnYMceIRy", "https://byte.co/@zachking/9rj2orF5gQ", "https://byte.co/@zachking/LkkbbYE1DyU", "https://byte.co/@zachking/LsVmWhe51bT"],
  url: "https://byte.co/@zachking"
}
```

#### Get a Byte

`GET http://localhost:8000/api/bytes/{id}`

**Request**

| Path Parameter | Required | Type   | Description       |
| -------------- | -------- | ------ | ----------------- |
| `id`           | required | string | The ID of a byte. |

**Response**

On success, the HTTP status code in the response header will be `200` OK and the response body will contain a byte object. On error, the response HTTP status code will be in the response header and the response body will contain an error object.

Example:

`GET http://localhost:8000/api/bytes/Lqs1Y8INqlt`

```
{
  id: "Lqs1Y8INqlt",
  user: "zachking",
  user_url: "https://byte.co/zachking",
  thumbnail_url: "https://e6k9t9a9.stackpathcdn.com/videos/URL33EDD7RHETMB7JBG5ZEBAWQ.jpg",
  caption: "Fetch with my best friend",
  created_at: "2d",
  loops: 16566,
  urls: ["https://byte.co/zachking/Lqs1Y8INqlt","https://byte.co/b/Lqs1Y8INqlt"]
}
```
