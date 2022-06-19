# Backend Code Challenge

REST endpoints for this service are

* GET /event with `id` as parameter
* GET /events with `from` and `to` as boundaries parameters
* POST /event with event as parameter in body with content type JSON header  - for creation
* PUT /event with event as parameter in body with content type JSON header - for update
* DELETE /event with `id` as parameter

Sample for PUT and POST is:

``` 
curl -X PUT 'localhost:8080/event'    -H "Content-Type: application/json"    -d '{"id":"03d66078-8da3-4d8a-9169-a15d69731fde","Title":"first event","start_time":"2022-05-19 18:30","end_time":"2022-05-19 11:30","address":"none","status":"done"}'
```

Use `make setup` for prepare environmemt, `make migrate`  to init database with data and `make run` for service start

## Minimum Requirements

x Create a golang service
x Connect it to an SQL database of choice
x Create the `events` table in the database
    - `title`: text
    - `start_time`: timestamp
    - `end_time`: timestamp
    - `address`: text
    - `status`: text / enum (pending, in progress, done)
    - Bonus: populate with dummy data
x events CRUD REST endpoints (Create, List, Update, Delete), also GET route added
x Add request validations on create and update
    - `title` required
    - `start_time` required
    - `end_time` required
    - `start_time` < end_time
    - `status` required, valid values are: `pending, in progress, done`

## Bonus Requirements

x Add date filtering to the List endpoint
`from` and `to` query parameters (timestamps)
respond with events between the timespan
x Unit tests (as much as you see fit)
x Log errors (`zap`, `logrus`, etc) (implement for the List endpoint)
x initial DB data
