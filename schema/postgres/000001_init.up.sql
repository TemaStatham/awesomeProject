CREATE TABLE IF NOT EXISTS Projects (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(255) NOT NULL,
                                        created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Goods (
                                    id SERIAL PRIMARY KEY,
                                    project_id SERIAL,
                                    name VARCHAR(255),
                                    description VARCHAR(255),
                                    priority INTEGER,
                                    removed BOOLEAN,
                                    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX ON Goods USING btree(name);

INSERT INTO Projects(name) VALUES ('Первая запись');
