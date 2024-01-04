package constants

const CreateLinkQuery = `INSERT INTO link (long_url, short_key, short_url)
						VALUES($1, $2, $3)
						RETURNING long_url, short_key, short_url`
