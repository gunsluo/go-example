# Oracle

### set password (default ORCLCDB)
```
docker exec -it oracle ./setPassword.sh password
```

### connect cdb 

```
docker exec -it oracle /bin/bash

sqlplus sys/password@//127.0.0.1:1521/ORCLCDB as sysdba

# show connection name
show con_name
```

### create cdb user
```
create user c##ac identified by password;
grant dba to c##ac;
grant create session to c##ac;
grant connect, resource to c##ac;
grant all privileges to c##ac;
```

### query all pdb
```
show pdbs

select GRANTEE,PRIVILEGE  from dba_sys_privs where GRANTEE='CONNECT';

SQL> select con_id,dbid,NAME,OPEN_MODE from v$pdbs;

    CON_ID	 DBID NAME															       OPEN_MODE
---------- ---------- -------------------------------------------------------------------------------------------------------------------------------- ----------
	 2 3281539280 PDB$SEED															       READ ONLY
	 3 3274890732 ORCLPDB1															       READ WRITE
```

### create pdb user
```
alter session set container=ORCLPDB1;
show con_name

create user ac identified by password;
grant dba to ac;
grant create session to ac;
grant connect, resource to ac;
grant all privileges to ac;
```

### connect pdb 

```
sqlplus ac/password@//127.0.0.1:1521/ORCLPDB1

# show connection name
show con_name
```
