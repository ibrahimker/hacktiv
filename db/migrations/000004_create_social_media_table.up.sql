begin;

create table if not exists public.social_media (
    id serial primary key,
    user_id int not null references public.users(id),
    name varchar(100) not null,
    url text not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

commit;