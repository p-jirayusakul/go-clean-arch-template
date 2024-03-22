CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS public.accounts
(
    id text NOT NULL DEFAULT uuid_generate_v4(),
    email text NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    CONSTRAINT accounts_pkey PRIMARY KEY (id),
    CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.addresses
(
    id text NOT NULL DEFAULT uuid_generate_v4(),
    street_address character varying(255),
    city character varying(100) NOT NULL,
    state_province character varying(100) NOT NULL,
    postal_code character varying(20) NOT NULL,
    country character varying(100) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    accounts_id text,
    CONSTRAINT addresses_pkey PRIMARY KEY (id),
    CONSTRAINT fk_accounts_address FOREIGN KEY (accounts_id)
        REFERENCES public.accounts (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE SET NULL
);