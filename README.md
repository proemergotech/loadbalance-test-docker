# loadbalance-test-docker
A mini application for testing where the loadbalancer directs tcp connections

## Usage:
1. run the server (see [docker-compose](#docker-compose-example))
2. `nc localhost 30001` 
3. You will see the ip of the container coming on the tcp connection every 2 sec.

If you replace localhost and port with your loadbalancer's host and port, this docker image can be used to see
which container the tcp request was routed to.

## docker-compose example:
```yaml
version: "3.6"

services:
  app:
    image: proemergotech/loadbalance-test
    ports:
      - 30001:12345
    environment:
      - PORT=12345

```