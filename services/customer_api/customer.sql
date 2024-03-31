DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
  id         INT AUTO_INCREMENT NOT NULL,
  first_name      VARCHAR(128) NOT NULL,
  last_name     VARCHAR(128) NOT NULL,
  date_of_birth      VARCHAR(10) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO customer
  (first_name, last_name, date_of_birth)
VALUES
  ('arun', 'duraisamy', "01-01-1988"),
   ('duraisamy', 'attiannan', "01-01-1958"),
  ('mani', 'duraisamy', "01-01-1960"),
  ('rathika', 'duraisamy', "01-01-1984");