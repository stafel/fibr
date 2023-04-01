# FIBR

A distributed language model service for Alpaca / Llama

## Inspiration

[Dalai](https://github.com/cocktailpeanut/dalai) provides a framework
to download, install and use the Alpaca and Llama language models locally using:

* [llama.cpp](https://github.com/ggerganov/llama.cpp)
* [alpaca.cpp](https://github.com/antimatter15/alpaca.cpp)

To run these model in a distributed system in a self hosted enviornment (e.g Kubernetes, Docker),
we can implement something similar using golang. 

## Benefits

The usage of these models are limited by the CPU and memory of the users device when ran locally.
If we can provide a way to run multiple workers (for each model) and a single proxy (with embedded API server)
to request responses, we can utilize distributed ressources and improve performance.

## Challenges

* Providing a frontend REST API (with OpenAPI spec to generate client), allowing clients to interact with the service
* Providing a frontend Websocket API, allowing clients to interact with the service
* Communication and discovery between the proxy and workers
* Sharing the models read only between workers (storage backend agnostic)
* Downloading and installing models at runtime (if not existent)
* Providing a unified configuration
* Authentification and Authorization
