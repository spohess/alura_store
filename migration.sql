create table public.products
(
    id          serial primary key,
    name        varchar,
    description varchar,
    price       numeric,
    quantity    integer
);

alter table public.products
    owner to alura;

INSERT INTO public.products (id, name, description, price, quantity)
VALUES  (1, 'Camiseta', 'Azul, bem bonita', 39, 10),
        (2, 'Tênis', 'Confortável', 399, 5),
        (3, 'Fone', 'Muito bom', 2500, 3),
        (4, 'Meia', 'Emborrachada', 10, 1000),
        (5, 'Teclado Gamer', 'Cheio de RGB', 350, 20),
        (6, 'Mouse Gamer', 'Com peso e RGB', 500, 100);