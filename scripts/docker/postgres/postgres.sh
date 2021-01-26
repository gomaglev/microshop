docker run --publish 5432:5432 \
  -e PG_MODE=primary \
  -e PG_PRIMARY_USER=postgres \
  -e PG_PRIMARY_PASSWORD=password \
  -e PG_DATABASE=postgres \
  -e PG_USER=postgres \
  -e PG_PASSWORD=password \
  -e PG_ROOT_PASSWORD=password \
  -e PG_PRIMARY_PORT=5432 \
  --name="postgres" \
  --hostname="postgres" \
  --network="n1" \
  --detach \
registry.developers.crunchydata.com/crunchydata/crunchy-postgres:centos7-12.2-4.2.2