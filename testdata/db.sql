
DROP TABLE IF EXISTS feed;
DROP TABLE IF EXISTS category;

CREATE TABLE category
(
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(240),
    category_image_url VARCHAR(240),
    category_description VARCHAR(240),
    feeds_count INTEGER
);

CREATE TABLE feed
(
    id SERIAL PRIMARY KEY,
    feed_name VARCHAR(240),
    feed_url TEXT,
    number_feed_entries INTEGER,
    feed_icon_url VARCHAR(240),
    last_fetched TIMESTAMP WITH TIME ZONE,
    feed_entries_count INTEGER,
    feed_description TEXT,
    category_id INTEGER NOT NULL,
    FOREIGN KEY (category_id) REFERENCES category (id) ON DELETE CASCADE
);