docker run --publish 5050:5050 \
  -e PGADMIN_SETUP_EMAIL=postgres \
  -e PGADMIN_SETUP_PASSWORD=password \
  -e SERVER_PORT=5050 \
  --name="pgadmin4" \
  --hostname="pgadmin4" \
  --network="n1" \
  --detach \
registry.developers.crunchydata.com/crunchydata/crunchy-pgadmin4:centos7-12.4-4.4.1