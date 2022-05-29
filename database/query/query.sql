-- name: GetUsers :many
select
    `id_user`,
    `username`,
    `email`,
    `created_on`,
    `last_login`
from
    users;
