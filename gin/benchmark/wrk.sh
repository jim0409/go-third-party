#!/bin/bash

# formula benchmark endpoint would be /benchmark
wrk -t10 -c10 -d5 http://127.0.0.1:8000/benchmark
