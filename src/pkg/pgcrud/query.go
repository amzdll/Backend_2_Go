package pgcrud

const query = `insert into %s (%s) values (%s)`
const get = `select * from %s where %s and %s`
const getall = `select * from %s limit $1 offset $2`
const update = `update %s set address_id = $1 where id = $2`
const delete = `delete from %s where %s`
