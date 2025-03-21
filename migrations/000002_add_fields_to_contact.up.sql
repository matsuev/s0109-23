ALTER TABLE IF EXISTS public.contact
    ADD COLUMN given_name character varying(50) NOT NULL DEFAULT '';

ALTER TABLE IF EXISTS public.contact
    ADD COLUMN family_name character varying(50) NOT NULL DEFAULT '';

ALTER TABLE IF EXISTS public.contact
    ADD COLUMN full_name character varying(200) NOT NULL DEFAULT '';
