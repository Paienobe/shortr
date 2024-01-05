package queries

const DeleteLinkQuery = `DELETE FROM link
						WHERE short_key = $1`
