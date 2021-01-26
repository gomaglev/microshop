docker run -it \
    --network n1 \
    --name mongo-express \
    -p 8081:8081 \
    -e ME_CONFIG_OPTIONS_EDITORTHEME="ambiance" \
    -e ME_CONFIG_MONGODB_SERVER="mongodb" \
    -e ME_CONFIG_MONGODB_ADMINUSERNAME="admin" \
    -e ME_CONFIG_MONGODB_ADMINPASSWORD="password" \
    -e ME_CONFIG_MONGODB_AUTH_DATABASE="admin" \
    -e ME_CONFIG_MONGODB_AUTH_USERNAME="admin" \
    -e ME_CONFIG_MONGODB_AUTH_PASSWORD="password" \
    -e ME_CONFIG_BASICAUTH_USERNAME="admin" \
    -e ME_CONFIG_BASICAUTH_PASSWORD="password" \
    mongo-express