CREATE TABLE IF NOT EXISTS tasks (
    id bigserial PRIMARY KEY,
    activation_time timestamp(0) with time zone NOT NULL,
    title text NOT NULL,
    done bool NOT NULL,
    version integer NOT NULL DEFAULT 1,
    UNIQUE(activation_time)
);

 
