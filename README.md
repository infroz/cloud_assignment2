# Assignment 2 for Cloud Technologies at NTNU Gjøvik
## Task Description
For this assignment, you will develop an API that aggregates information from our Gitlab deployment i.e., https://git.gvk.idi.ntnu.no. Please ensure you use our deployment, not gitlab.com, since you will most certainly trigger their rate limit.
This includes two parts,

- the development of an API for direct invocation, as well as an
- interface for the registration of Webhooks for invocation upon certain events.

The developed services will be dockerised and deployed on OpenStack (we will talk about both those in upcoming sessions). For persistence, use a NoSQL storage option of your choice.

## How to run
In /cmd/ go run main.go

## Endpoints
http://url:port/repocheck/v1/commits{?limit=[0-9]+{&auth=<access-token>}}
  returns repositories with most commits
  
http://url:port//repocheck/v1/languages{?limit=[0-9]+{&auth=<access-token>}}
   Returns most popular languages in order descending
http://url:port//repocheck/v1/issues{?type=(users|labels){&auth=<access-token>}}
  Not implemented
http://url:port//repocheck/v1/status
  Service status
