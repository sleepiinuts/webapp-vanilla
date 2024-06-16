-- name: findByArrivalAndDeparture
SELECT id,user_id,room_id,arrival,departure 
FROM reservations 
WHERE arrival <= $2 AND departure >= $1
