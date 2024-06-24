-- name: new
INSERT INTO users (first_name,last_name,email,"password" ,phone,"role",created_at,updated_at) 
VALUES ($1,$2,$3,$4,$5,$6,current_timestamp,current_timestamp) RETURNING id;

-- name: authen
SELECT id,password FROM users WHERE email=$1

-- name: findByEmail
SELECT first_name,last_name,email,"password",phone,"role" FROM users WHERE email=$1