# Go-Microservice-Proxy

A (WIP/PoC) proxy server written in GO

## Motivation

I've recently been building larger front-end applications as a set of microservices, where each is responsible for serving the entire payload for a subsection of the larger app. This works out well for experimenting with different front-end frameworks/libraries/languages/architectures/etc. but it also introduces a number of challenges, one of those being "how can I run my microservices under different paths of the same URL instead of subdomains?".

One possible solution is to host a reverse proxy at `https://www.myapp.com` that will then route requests to the microservices based on the request path, which is what this project aims to do. Eventually I'd like to inject the front-end microservice's response into a template hosted here in the orchestration layer where I could centralize some logic and UI in a single app rather than in a shared library.
