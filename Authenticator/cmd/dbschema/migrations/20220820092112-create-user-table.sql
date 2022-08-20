-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id            BIGINT PRIMARY KEY,
    name          TEXT                NOT NULL,
    created_at    TIMESTAMP           NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP           NOT NULL DEFAULT now()
);

-- +migrate StatementBegin
create
    or replace function trigger_set_updated_at() returns trigger as
$$
begin
    NEW.updated_at = now();

    return NEW;

end
$$ language plpgsql;

create trigger
    users_set_updated_at
    before
        update
    on users
    for each row
execute
    procedure trigger_set_updated_at();

-- +migrate StatementEnd

-- +migrate Down
drop function
    if exists users_set_updated_at cascade;
drop function
    if exists trigger_set_updated_at cascade;

DROP TABLE IF EXISTS users;
