package repository

const (
	queryCreate = `insert into client (client_name, client_surname, birthday, gender, address_id) values ($1, $2, $3, $4, $5)`
	queryGet    = `select * from client where client_name = $1 and client_surname = $2`
	queryGetAll = `select * from client limit $1 offset $2`
	queryUpdate = `update client set address_id = $1 where id = $2`
	queryDelete = `delete from client where id = $1`
)
