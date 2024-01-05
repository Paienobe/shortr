package queries

const CreateLinkQuery = `INSERT INTO link (long_url, short_key, short_url)
						VALUES($1, $2, $3)
						ON CONFLICT (short_key) DO UPDATE
						SET short_key = EXCLUDED.short_key
						RETURNING long_url, short_key, short_url`
