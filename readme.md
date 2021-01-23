Background Job System
=====================
> https://www.codementor.io/projects/background-job-system-atx32exogo

## How to use

Start the server:

```golang
go run *.go
```

This will start both the webserver and the jobs handler.

## Web page

Open your preferred browser, reach the page `localhost:8081`.

In this page, you can insert a job.

## What do you mean with 'job'?

A job is composed by some parameters.

- a payload
- a payload content-type
- an endpoint
- a HTTP status code
- a due date
- and a friendly name for your eyes

### Payload

A payload is what you need to elaborate. Ideally a JSON but you could send
whatever you need.

### Payload Content-Type

The receiving server will be more happy if you tell it what kind of payload
you are going to send to it. Defaults to `application/json`, change if needed.

### Endpoint

The server who will receive this job and will digest. Jep, because this is a
forwarding server only. The dirty work is made by `endpoint`s.

What, don't you have an endpoint to use for testing this?
I'll tell you a secret. The folder `/listener` contains a dummy web server ;)

### HTTP Status Code

The job handler will flag as `done` a job if the `endpoint` returns a HTTP
Status Code like the one you want.

### Due date

Set a date in the past and the job is run ASAP.

### Name

This is for you. In the background, an UUID will be generated and used.

## Jobs page

Here, with a paginator made both on server and client side, you will see
everything you need to see. Select a job if you want to delete or run ASAP.

## ASAP?

What do I mean with ASAP? As soon as possible, in this server means as far as
the tick clocks, and this happens every minute.

## Tests

Todo ASAP.
