create table public.products
(
    id          serial    primary key,
    name        varchar                 not null,
    description varchar,
    price       numeric                 not null,
    quantity    integer   default 1,
    created_at  timestamp default now() not null,
    updated_at  timestamp               not null
);

alter table public.products
    owner to alura;
