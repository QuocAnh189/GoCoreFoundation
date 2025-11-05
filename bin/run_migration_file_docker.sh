#!/usr/bin/env bash
migration_file=$1
if [ -z "$migration_file" ]; then
    echo "Usage: $0 <migration_file>"
    exit 1
fi

# Extract DB credentials from .env.docker file
USER=$(grep -E '^DB_USER=' .env.docker | cut -d '=' -f 2)
PASSWORD=$(grep -E '^DB_PASSWORD=' .env.docker | cut -d '=' -f 2)
HOST=$(grep -E '^DB_HOST_MIGRATION=' .env.docker | cut -d '=' -f 2)
PORT=$(grep -E '^DB_PORT_MIGRATION=' .env.docker | cut -d '=' -f 2)
NAME=$(grep -E '^DB_NAME=' .env.docker | cut -d '=' -f 2)

# Fallbacks (if migration-specific host/port are not defined)
if [ -z "$HOST" ]; then
    HOST="db"      # container name in docker-compose
fi
if [ -z "$PORT" ]; then
    PORT="3306"    # internal MySQL port
fi

if [ -z "$USER" ] || [ -z "$HOST" ] || [ -z "$PORT" ] || [ -z "$NAME" ]; then
    echo "❌ Error: .env.docker file is missing one or more required variables [DB_USER, DB_HOST, DB_PORT, DB_NAME]"
    exit 1
fi

MYSQL_CONTAINER="db"

echo "🚀 Running migration: $migration_file"
echo "   -> Host: $HOST"
echo "   -> Port: $PORT"
echo "   -> DB:   $NAME"
echo "   -> User: $USER"
echo ""

if [ -z "$PASSWORD" ]; then
    docker exec -i $MYSQL_CONTAINER sh -c \
        "mysql --force -u$USER -h$HOST -P$PORT $NAME" < "$migration_file"
else
    docker exec -i $MYSQL_CONTAINER sh -c \
        "mysql --force -u$USER -p$PASSWORD -h$HOST -P$PORT $NAME" < "$migration_file"
fi

status=$?
if [ $status -eq 0 ]; then
    echo "✅ Migration completed successfully."
else
    echo "❌ Migration failed with status code $status."
fi