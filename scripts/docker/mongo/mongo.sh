docker build  -f /Users/jasondu/microshop/debug/mongo/DockerFile -t mongoauth .

docker run -p27017:27017 --name mongodb --network n1 --env-file /Users/jasondu/microshop/debug/mongo/.env -d mongoauth
