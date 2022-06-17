# Backend Code Challenge

Welcome Candidate!

The purpose of this challenge is primarily to demonstrate your ability to understand requirements and build a slim, simple to use microservice while following best practices and architectural patterns.

In this code challenge, we would like you to build a golang microservice that serves REST endpoints for the `events` resource. Events are going to be used to display calendar events on a Scheduler. The UI **could** (bonus requirement) be limited between two time ranges: `from` and `to` (think about Google Calendar events). 

The user should be able to perform CRUD operations on calendar events, using the service's REST endpoints. 

`events` should reside in an SQL database, in a table with the same name. Changes to `events` performed over the endpoints should obviously be persisted to the database as well.

We have prepared some boilerplate code that connects to a Postgres database and runs `gin` for routing. It's the candidate's choice to either use the boilerpate code provided, or just build it from scratch. 

The Candidate **MUST** submit their solution into a git repository (github, gitlab, etc), and invite the reviewers as contributors. It would be best if you could open a Pull/Merge Request with your changes, so that we can do a proper code review on it.

We will be asking about the technical decisions you have made in your solution, and about how you structured your code and files. Make sure to use best coding principles!

You will have 3 hours to do the challenge, there's no need to complete all the bonus requirements. Pick the ones you find most important, and try to have a complete solution for us that works.

## Minimum Requirements
- Create a golang service
- Connect it to an SQL database of choice
- Create the `events` table in the database
    - `title`: text
    - `start_time`: timestamp
    - `end_time`: timestamp
    - `address`: text
    - `status`: text / enum (pending, in progress, done)
    - Bonus: populate with dummy data
- events CRUD REST endpoints (Create, List, Update, Delete)
- Add request validations on create and update
    - `title` required
    - `start_time` required
    - `end_time` required
    - `start_time` < end_time
    - `status` required, valid values are: `pending, in progress, done`

## Bonus Requirements
- Create an addresses table in the database and make the relation between it and the events table. Use address_id in events table to reference it
    - `full_address`: text
- Add date filtering to the List endpoint
`from` and `to` query parameters (timestamps)
respond with events between the timespan 
- Unit tests (as much as you see fit)
- Basic authentication in the endpoints, use a middleware
- Log errors (`zap`, `logrus`, etc) (implement for the List endpoint)


Best of luck!