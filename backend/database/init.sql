CREATE TABLE public.person_events
(
    id uuid NOT NULL,
    person_id uuid NOT NULL,
    event_type text COLLATE pg_catalog."default" NOT NULL,
    "timestamp" timestamp without time zone NOT NULL DEFAULT now(),
    data jsonb,
    CONSTRAINT person_events_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.person_events
    OWNER to postgres;