CREATE TABLE IF NOT EXISTS forms (
	source VARCHAR NOT NULL CHECK (length(source < 33)),
	content VARCHAR NOT NULL CHECK (length(content < 4097))
);


