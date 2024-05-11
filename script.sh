docker network create db-network

cd ./db/
docker-compose up -d

cd ..

cd ./account-manager/
docker-compose up -d

cd ..

cd ./payment-manager/
docker-compose up -d

cd ..

# Wait for the database container to be ready
echo "Waiting for the database to start..."
while ! docker exec postgres_dev pg_isready -U superuser -d db_identity > /dev/null 2>&1; do
    sleep 1
done

# Execute SQL queries from file
echo "Executing SQL queries..."
docker exec -i postgres_dev psql -U superuser -d db_identity < ./query.sql

echo "SQL queries executed successfully."