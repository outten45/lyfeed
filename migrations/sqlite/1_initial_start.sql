-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
create table items (
  Id integer primary key,
  Feed text,
  Channel text,
  Guid text,
  Title text,
  Description text,
  Author text,
  PubDate date
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE items;
