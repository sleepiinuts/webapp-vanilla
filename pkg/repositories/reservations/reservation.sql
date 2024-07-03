-- name: findByArrivalAndDeparture
SELECT id,user_id,room_id,arrival,departure 
FROM reservations 
WHERE arrival <= $2 AND departure >= $1

-- name: findByIdAndArrAndDep
SELECT id,user_id,room_id,arrival,departure 
FROM reservations 
WHERE id = $1 AND arrival <= $2 AND departure >= $3

