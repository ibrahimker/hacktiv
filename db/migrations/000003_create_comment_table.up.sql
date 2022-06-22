begin;

create table if not exists public.comment (
    id serial primary key,
    user_id int not null references public.users(id),
    photo_id int not null references public.photos(id),
    message text not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

commit;