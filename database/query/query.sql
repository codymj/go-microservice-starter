-- name: GetUsers :many
select
    id_user,
    username,
    email,
    created_on,
    last_login
from
    users;

-- name: GetUserById :one
select
    id_user,
    username,
    email,
    created_on,
    last_login
from
    users
where
    id_user = $1 limit 1;
