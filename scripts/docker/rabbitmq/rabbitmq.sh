# default manage port is 15672, guest/guest
docker run -d --publish 15672:15672 --publish 5672:5672 --hostname rabbitmq --name rabbitmq -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management
