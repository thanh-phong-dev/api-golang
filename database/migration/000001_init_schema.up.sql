CREATE TABLE public.accounts (
	id serial NOT NULL,
	username text NOT NULL,
	"password" text NOT NULL,
	email text NULL,
	fullname text NOT NULL,
	phonenumber text NOT NULL,
	"role" text NOT NULL,
	sex text NOT NULL,
	address text NULL,
	date_of_birth text NULL,
	is_deleted bool NOT NULL DEFAULT false,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT accounts_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX accounts_username_idx ON public.accounts (username);
CREATE UNIQUE INDEX accounts_email_idx ON public.accounts (email);
