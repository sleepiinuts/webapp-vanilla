-- name: newUser
INSERT INTO users (first_name,last_name,email,"password" ,phone,"role",created_at,updated_at) 
VALUES ($firstname,$lastName,$email,$password,$phone,$role,current_timestamp,current_timestamp) RETURNING id;