package queries

const FindLinkQuery = `SELECT long_url FROM link
						WHERE short_key = $1`
