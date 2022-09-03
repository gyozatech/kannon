// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: pool.sql

package sqlc

import (
	"context"
	"time"
)

const cleanPool = `-- name: CleanPool :exec
DELETE FROM sending_pool_emails 
WHERE email = $1 AND message_id = $2
`

type CleanPoolParams struct {
	Email     string
	MessageID string
}

func (q *Queries) CleanPool(ctx context.Context, arg CleanPoolParams) error {
	_, err := q.exec(ctx, q.cleanPoolStmt, cleanPool, arg.Email, arg.MessageID)
	return err
}

const getPool = `-- name: GetPool :one
SELECT id, scheduled_time, original_scheduled_time, send_attempts_cnt, email, message_id, error_msg, error_code, fields, status FROM  sending_pool_emails 
WHERE email = $1 AND message_id = $2
`

type GetPoolParams struct {
	Email     string
	MessageID string
}

func (q *Queries) GetPool(ctx context.Context, arg GetPoolParams) (SendingPoolEmail, error) {
	row := q.queryRow(ctx, q.getPoolStmt, getPool, arg.Email, arg.MessageID)
	var i SendingPoolEmail
	err := row.Scan(
		&i.ID,
		&i.ScheduledTime,
		&i.OriginalScheduledTime,
		&i.SendAttemptsCnt,
		&i.Email,
		&i.MessageID,
		&i.ErrorMsg,
		&i.ErrorCode,
		&i.Fields,
		&i.Status,
	)
	return i, err
}

const getSendingPoolsEmails = `-- name: GetSendingPoolsEmails :many
SELECT id, scheduled_time, original_scheduled_time, send_attempts_cnt, email, message_id, error_msg, error_code, fields, status FROM sending_pool_emails WHERE message_id = $1 LIMIT $2 OFFSET $3
`

type GetSendingPoolsEmailsParams struct {
	MessageID string
	Limit     int32
	Offset    int32
}

func (q *Queries) GetSendingPoolsEmails(ctx context.Context, arg GetSendingPoolsEmailsParams) ([]SendingPoolEmail, error) {
	rows, err := q.query(ctx, q.getSendingPoolsEmailsStmt, getSendingPoolsEmails, arg.MessageID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SendingPoolEmail
	for rows.Next() {
		var i SendingPoolEmail
		if err := rows.Scan(
			&i.ID,
			&i.ScheduledTime,
			&i.OriginalScheduledTime,
			&i.SendAttemptsCnt,
			&i.Email,
			&i.MessageID,
			&i.ErrorMsg,
			&i.ErrorCode,
			&i.Fields,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getToVerify = `-- name: GetToVerify :many
SELECT id, scheduled_time, original_scheduled_time, send_attempts_cnt, email, message_id, error_msg, error_code, fields, status FROM sending_pool_emails
    WHERE status = 'to_verify' ORDER BY scheduled_time asc
    LIMIT $1
`

func (q *Queries) GetToVerify(ctx context.Context, limit int32) ([]SendingPoolEmail, error) {
	rows, err := q.query(ctx, q.getToVerifyStmt, getToVerify, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SendingPoolEmail
	for rows.Next() {
		var i SendingPoolEmail
		if err := rows.Scan(
			&i.ID,
			&i.ScheduledTime,
			&i.OriginalScheduledTime,
			&i.SendAttemptsCnt,
			&i.Email,
			&i.MessageID,
			&i.ErrorMsg,
			&i.ErrorCode,
			&i.Fields,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const prepareForSend = `-- name: PrepareForSend :many
UPDATE sending_pool_emails AS sp
    SET status = 'sending'
    FROM (
            SELECT id FROM sending_pool_emails
            WHERE scheduled_time <= NOW() AND status = 'scheduled'
            LIMIT $1
        ) AS t
    WHERE sp.id = t.id
    RETURNING sp.id, sp.scheduled_time, sp.original_scheduled_time, sp.send_attempts_cnt, sp.email, sp.message_id, sp.error_msg, sp.error_code, sp.fields, sp.status
`

func (q *Queries) PrepareForSend(ctx context.Context, limit int32) ([]SendingPoolEmail, error) {
	rows, err := q.query(ctx, q.prepareForSendStmt, prepareForSend, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SendingPoolEmail
	for rows.Next() {
		var i SendingPoolEmail
		if err := rows.Scan(
			&i.ID,
			&i.ScheduledTime,
			&i.OriginalScheduledTime,
			&i.SendAttemptsCnt,
			&i.Email,
			&i.MessageID,
			&i.ErrorMsg,
			&i.ErrorCode,
			&i.Fields,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const reschedulePool = `-- name: ReschedulePool :exec
UPDATE sending_pool_emails 
SET status='scheduled', scheduled_time =  $1, send_attempts_cnt = send_attempts_cnt + 1 WHERE email = $2 AND message_id = $3
`

type ReschedulePoolParams struct {
	ScheduledTime time.Time
	Email         string
	MessageID     string
}

func (q *Queries) ReschedulePool(ctx context.Context, arg ReschedulePoolParams) error {
	_, err := q.exec(ctx, q.reschedulePoolStmt, reschedulePool, arg.ScheduledTime, arg.Email, arg.MessageID)
	return err
}

const setSendingPoolDelivered = `-- name: SetSendingPoolDelivered :exec
UPDATE sending_pool_emails 
	SET status = 'sent' WHERE email = $1 AND message_id = $2
`

type SetSendingPoolDeliveredParams struct {
	Email     string
	MessageID string
}

func (q *Queries) SetSendingPoolDelivered(ctx context.Context, arg SetSendingPoolDeliveredParams) error {
	_, err := q.exec(ctx, q.setSendingPoolDeliveredStmt, setSendingPoolDelivered, arg.Email, arg.MessageID)
	return err
}

const setSendingPoolScheduled = `-- name: SetSendingPoolScheduled :exec
UPDATE sending_pool_emails 
	SET status = 'scheduled' WHERE email = $1 AND message_id = $2
`

type SetSendingPoolScheduledParams struct {
	Email     string
	MessageID string
}

func (q *Queries) SetSendingPoolScheduled(ctx context.Context, arg SetSendingPoolScheduledParams) error {
	_, err := q.exec(ctx, q.setSendingPoolScheduledStmt, setSendingPoolScheduled, arg.Email, arg.MessageID)
	return err
}
