CREATE DATABASE device_service;

\c device_service;


CREATE TABLE devices (
    id UUID PRIMARY KEY,
    device_name VARCHAR(255) NOT NULL,
    device_brand VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);
