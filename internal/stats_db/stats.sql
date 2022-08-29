-- name: InsertPrepared :exec
INSERT INTO prepared (email, message_id, timestamp, first_timestamp, domain) VALUES (@email, @message_id, @timestamp, @timestamp, @domain)
	ON CONFLICT (email, message_id, domain) DO UPDATE
	SET timestamp = @timestamp;

-- name: InsertStat :exec
INSERT INTO stats (email, message_id, type, timestamp, domain, data) VALUES  (@email, @message_id, @type, @timestamp, @domain, @data);

-- name: QueryStats :many
SELECT * FROM stats 
WHERE domain = $1 
AND timestamp BETWEEN @start AND @stop
LIMIT @take OFFSET @skip;

-- name: CountQueryStats :one
SELECT COUNT(*) FROM stats 
WHERE domain = $1 
AND timestamp BETWEEN @start AND @stop;