begin;

create table if not exists public.photos (
    id serial primary key,
    title varchar(100) not null,
    caption varchar(100) not null,
    photo_url text not null,
    user_id int not null references public.users(id),
    created_at timestamp not null,
    updated_at timestamp not null
);

commit;