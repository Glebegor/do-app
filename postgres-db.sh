set -e

host="$1"
shift
cmd="$@"

until POSTGRES_PASSWORD=$DB_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
    >&2 echo "POSTGRES IS UNVAILABLE - sleeping"
    sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd