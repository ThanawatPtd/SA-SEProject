-- Create user table first
CREATE TABLE "user" (
    id UUID,
    email VARCHAR(100) NOT NULL,
    fname VARCHAR(100) NOT NULL,
    lname VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,

    PRIMARY KEY (id)
);

-- Then create employee table that inherits from user
CREATE TABLE employee (
    salary REAL NOT NULL,
    position VARCHAR(10) NOT NULL
) INHERITS ("user");