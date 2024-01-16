# URL Shortener

## Description
This is a simple URL shortener API implemented in Go. It allows users to submit any URL to be shortened and receive a valid shortened URL.

## API Endpoints

### Shorten a URL
- **Endpoint:** `/shortcode`
- **Method:** PUT
- **Request Body:**
  ```json
  {
    "destination": "valid url"
  }


<h1 align="center">URL Shortener</h1>

---


## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Built Using](#built_using)
- [Running the tests](#tests)
- [Deployment](#deployment)
- [Authors](#authors)

## 🧐 About <a name = "about"></a>

The URL Shortener is a program which takes a Long URL in input and provides a Short URL in output. 
This Program uses Redis to Store Data Permanently.
The Documentation of APIs will be found at [API Documentation](#usage)

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [Deployment](#deployment) for notes on how to deploy the project on a Local System or on a Kubernetes Server.

### Prerequisites

To run the URL Shortener on Local System, first we need to install following Software Dependencies.

- [Golang](https://go.dev/doc/install)

Once above Dependencies are installed we can move with [further steps](#installing)

### Installing <a name = "installing"></a>

A step by step series of examples that tell you how to get a development env running.

#### Step 1: Setting Up Environmental Variables

Set up the Environmental variables according to your needs. The Application will run with defaults as mentioned in the following table

| Environmental Variable | Usage                              | 
|------------------------|------------------------------------|
| REDIS_HOST             | Host where Redis is Running        |
| REDIS_PORT             | Port where Redis is Running        |

#### Step 2: Run the Code
```
go build
./url-shortner
```

## 🔧 Running the tests <a name = "tests"></a>

To Run the Test Cases, Open a terminal in the Project and run following command
```
go test ./...
```

## 📃 API Documentation <a name="usage"></a>

Here are the API Endpoint used in this Project

```
ENDPOINT: /
REQUEST TYPE: GET
RESPONSES:
  Welcome to the URL Shortener!
```

```
ENDPOINT: /shortcode
REQUEST TYPE: PUT
BODY: 
RESPONSES:
  200: 
    {
      "short_url": "http://localhost/r/100680ad"
    }

  400:
    {
      "message": "Invalid JSON Format"
    }
  OR
    {
      "message": "Missing 'destination' field"
    }
  OR
    {
      "message": "Invalid URL"
    }

  500:
    {
      "message": "Internal Server Error : "
    }
```

```
ENDPOINT: r/{short_url}
REQUEST TYPE: GET
PATH PARAMETERS: short_url
RESPONSES:
  200: 
    Redirect to the Long URL
  
  404:
    {
      "message": "URL not found for localhost:5000/tstng"
    }
```

To check & Test the Supported Endpoints, and it's documentation, Run the Project and kindly use any of the methods mentioned below

## 🚀 Deployment <a name = "deployment"></a>

In order to deploy the Project on a Kubernetes Server kindly follow below steps:
```
kubectl create -f ./deployment/redis-service.yaml
```
```
kubectl create -f ./deployment/redis-deployment.yaml
```
```
kubectl create -f ./deployment/url-shortner-deployment.yaml
```
```
kubectl create -f ./deployment/url-shortner-service.yaml
```

API Host will be exposed on the URL Provided by URL Shortner Service

## ⛏️ Built Using <a name = "built_using"></a>

- [Redis](https://redis.com/) - Primary Database
- [Golang](https://go.dev/) - Backend
- [Docker](https://www.docker.com/) - Containers Solution

## ✍️ Authors <a name = "authors"></a>

- [@akankshakumari393](https://github.com/akankshakumari393) - Idea & Implementation