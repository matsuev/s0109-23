-- Create table public.contact

CREATE TABLE IF NOT EXISTS public.contact
(
   id bigserial NOT NULL, PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.contact
   OWNER to admin;