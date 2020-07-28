#!/bin/bash

APP_NAME="printer"
DB_NAME="$APP_NAME"
DOCKER_DB="${APP_NAME}_db"
DOCKER_DBUI="${APP_NAME}_dbui"
DOCKER_NETWORK="$APP_NAME"
SQL_FILES=(uuid.sql location.sql printer.sql user.sql)
SQL_PATH="./schema"
OUTPUT_PATH="./db"
PKG_NAME="db"
POSTGRES_VERSION="postgres:12"
PGADMIN_VERSION="dpage/pgadmin4:latest"

function generate_sqlboiler_config() {
    ip=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' printer_db)
    cat >sqlboiler.toml <<EOF
[psql]
    dbname  = "${DB_NAME}"
    host    = "${ip}"
    port    = 5432
    user    = "postgres"
    sslmode = "disable"
EOF
}

function generate_sqlboiler() {
    mkdir -p $OUTPUT_PATH
    opath=$(readlink -f $OUTPUT_PATH)
    tmp_dir=$(mktemp -d -t sql-XXXXXXXX)
    pushd $tmp_dir > /dev/null

    generate_sqlboiler_config >/dev/null

    sqlboiler -o $opath -p $PKG_NAME \
        --templates "$GOPATH/src/github.com/volatiletech/sqlboiler/templates" \
        --templates "$GOPATH/src/github.com/volatiletech/sqlboiler/templates_test" \
        --templates "$GOPATH/src/github.com/korylprince/httputil/templates" \
        --wipe psql

    goimports -w $opath

    popd > /dev/null
    rm -rf $tmp_dir
    go install $OUTPUT_PATH
}

function test_sqlboiler() {
    opath=$(readlink -f $OUTPUT_PATH)
    tmp_dir=$(mktemp -d -t sql-XXXXXXXX)
    pushd $tmp_dir > /dev/null

    generate_sqlboiler_config

    echo "#!/bin/bash" > pg_dump
    echo "docker run -i --rm --network $DOCKER_NETWORK $POSTGRES_VERSION pg_dump -Upostgres -h $DOCKER_DB \"\$@\"" >> pg_dump
    chmod +x pg_dump

    PATH=$tmp_dir:$PATH go test $opath -test.config=$tmp_dir/sqlboiler.toml

    popd > /dev/null
    rm -rf $tmp_dir
}

function reset_db() {
    docker exec $DOCKER_DB psql -Upostgres -c "REVOKE CONNECT ON DATABASE $DB_NAME FROM public;"
    docker exec $DOCKER_DB psql -Upostgres -c "SELECT pid, pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '$DB_NAME' AND pid <> pg_backend_pid();"
    docker exec $DOCKER_DB psql -Upostgres -c "DROP DATABASE IF EXISTS $DB_NAME;"
    docker exec $DOCKER_DB psql -Upostgres -c "CREATE DATABASE $DB_NAME;"
    pushd $SQL_PATH > /dev/null
    cat ${SQL_FILES[@]} | docker exec -i $DOCKER_DB psql -Upostgres $DB_NAME
    popd > /dev/null
}

function stop_docker() {
    docker kill $DOCKER_DB $DOCKER_DBUI
    docker rm $DOCKER_DB $DOCKER_DBUI
    docker network rm $DOCKER_NETWORK
}

function start_docker() {
    docker network create $DOCKER_NETWORK

    docker run -d --name $DOCKER_DB \
        --network $DOCKER_NETWORK -p 5432:5432 \
        -e "POSTGRES_HOST_AUTH_METHOD=trust" \
        $POSTGRES_VERSION

    docker run -d --name $DOCKER_DBUI \
        --network $DOCKER_NETWORK -p 9090:80 \
        -e "PGADMIN_DEFAULT_EMAIL=test@test.com" \
        -e "PGADMIN_DEFAULT_PASSWORD=test" \
        $PGADMIN_VERSION
}

case "$1" in
    stop)
        stop_docker
        ;;
    restart)
        stop_docker
        sleep 3
        start_docker
        sleep 3
        reset_db
        sleep 1
        generate_sqlboiler
        # skip tests since they fail because of constraints
        # test_sqlboiler
        ;;
    reset)
        reset_db
        sleep 1
        generate_sqlboiler
        # skip tests since they fail because of constraints
        # test_sqlboiler
        ;;
    generate)
        generate_sqlboiler
        ;;
esac
