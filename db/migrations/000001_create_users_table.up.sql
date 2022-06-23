begin;

create table if not exists public.users (
    id serial primary key,
    username varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    age integer not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

commit;