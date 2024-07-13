-- name: findByIdAndEffAndExp
SELECT id,room_id,"type",effective,expire
FROM restrictions
WHERE room_id = $1 AND effective >= $2 AND effective <= $3