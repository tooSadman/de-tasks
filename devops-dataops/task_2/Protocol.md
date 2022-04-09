# Task 2 - Docker

## Building the image.
To build the image run command: `docker build -t toosadman/task_2:1.0.0 .`  

## Pushing image to DockerHub.
To push the built image to DockerHub the command was run: `docker push toosadman/task_2:1.0.0`  

## Using the image.
To run the container execute: `docker run --name task2 -p 8080:8080 -d toosadman/task_2:1.0.0`  
Run `curl localhost:8080/world` to test if container works properly.
