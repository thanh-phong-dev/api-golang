CREATE TABLE public.accounts (
	id serial NOT NULL,
	username text NOT NULL,
	"password" text NOT NULL,
	email text NULL,
	fullname text NOT NULL,
	phonenumber text NOT NULL,
	"role" text NOT NULL,
	CONSTRAINT accounts_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX accounts_username_idx ON public.accounts (username);
CREATE UNIQUE INDEX accounts_email_idx ON public.accounts (email);
