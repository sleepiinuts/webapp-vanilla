-- name: findAll
select id,name,description,price,img_path from rooms

-- name: findById
select id,name,description,price,img_path from rooms where id=$1