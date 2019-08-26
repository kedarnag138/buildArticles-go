-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
  id SERIAL,
  title varchar(100),
  content text,
  author varchar(320),
  PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE articles;
-- +goose StatementEnd
