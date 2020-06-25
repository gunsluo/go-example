# Trace Demo

### sql
```
CREATE TABLE account (
  id varchar(256) PRIMARY KEY,
  name text NOT NULL DEFAULT '',
  email text NOT NULL DEFAULT ''
);

CREATE TABLE identity (
  id varchar(256) PRIMARY KEY,
  name text NOT NULL DEFAULT '',
  cert_id text NOT NULL DEFAULT ''
);

INSERT INTO account(id, name, email) VALUES('123','Rachel Floral Designs','rachel@test.com');
INSERT INTO account(id, name, email) VALUES('567','Amazing Coffee Roasters','amazing@test.com');
INSERT INTO account(id, name, email) VALUES('392','Trom Chocolatier','trom@test.com');
INSERT INTO account(id, name, email) VALUES('731','Japanese Desserts','dess@test.com');

INSERT INTO identity(id, name, cert_id) VALUES('123','Rachel Floral Designs','xxxx001');
INSERT INTO identity(id, name, cert_id) VALUES('567','Amazing Coffee Roasters','xxxx002');
INSERT INTO identity(id, name, cert_id) VALUES('392','Trom Chocolatier','xxxx003');
INSERT INTO identity(id, name, cert_id) VALUES('731','Japanese Desserts','xxxx004');
```
