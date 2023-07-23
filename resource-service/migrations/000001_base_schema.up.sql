CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS workgroups (
    id bigserial primary key,
    group text not null unique
);

CREATE TABLE IF NOT EXISTS job_titles (
    id bigserial primary key,
    title text not null unique
);

CREATE TABLE IF NOT EXISTS employment_types (
    id bigserial primary key,
    type text not null unique
);

CREATE TABLE IF NOT EXISTS resources (
    id integer primary key,
    name text not null,
    email text not null unique,
    job_title_id integer not null,
    workgroup_id integer not null,
    employment_type_id integer not null,
    manager_id integer,
    active boolean DEfAULT 1
);