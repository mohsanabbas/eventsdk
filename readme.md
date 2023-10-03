# Event Tracking SDK

An SDK for sending and batching event data to an endpoint. Easily configurable and extendable.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
    - [Initialization](#initialization)
    - [Sending Individual Events](#sending-individual-events)
    - [Batching Events](#batching-events)
    - [HTTP Client Configuration](#http-client-configuration)
    - [Error Handling](#error-handling)
- [Contribute](#contribute)
- [License](#license)

## Features

- Send individual events to an endpoint.
- Batch events and send them together.
- Customizable HTTP client with built-in retry logic.
- Easily extendable and integrable with other systems.

## Installation

```bash
go get github.com/mohsanabbas/event-tracking-sdk
```


## Usage

### Initialization

Create a new SDK instance:

```go
sdk := eventsdk.New(endpoint, transport, marshaller, config)
```
## Sending Individual Events
### To send an individual event:

```go
err := sdk.SendEvent("track", trackEvent)
```

## Batching Events
### To enable batching:
```go
config := eventsdk.SDKConfig{
    EnableBatching: true,
    BatchSize:      10,
    MaxWaitTime:    1 * time.Minute, // max time to accumulate batch requests before flush
}
sdk := eventsdk.New("https://httpbin.org/post", httpTransport, eventCodec, config)

```

## HTTP Client Configuration
You can configure the HTTP client to adjust various parameters, like timeouts and retry mechanisms:

```go
clientConfig := transport.HTTPClientConfig{
    SemaphoreBufferSize: 50,
    Timeout:             10 * time.Second,
    // ... other configurations
}
customClient := transport.NewCustomHTTPClient(clientConfig)

```

## Error Handling
Errors can occur during the event sending or batching processes. Always check for errors and handle them appropriately:

```go
err := sdk.SendEvent("track", trackEvent)
if err != nil {
    // Handle error
}

```