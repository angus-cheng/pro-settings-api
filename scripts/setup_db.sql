CREATE TABLE if NOT EXISTS pro_settings (
    id SERIAL PRIMARY KEY,
    name TEXT,
    sensitivity FLOAT,
    scraped_at TIMESTAMPTZ DEFAULT NOW()
)