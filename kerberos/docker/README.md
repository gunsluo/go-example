# KDC Intergation Test Instance for TEST.KRB5.COM

DO NOT USE THIS CONTAINER FOR ANY PRODUCTION USE!!!

To run:
```bash
docker run -v /etc/localtime:/etc/localtime:ro -p 88:88 -p 88:88/udp -p 464:464 -p 464:464/udp --rm --name gokrb5-kdc-centos-default gunsluo/gokrb5:kdc-centos-default &
```

To build:
```bash
docker build -t gunsluo/gokrb5:kdc-centos-default --force-rm=true --rm=true .
docker push gunsluo/gokrb5:kdc-centos-default
```

Command


1. For User

Add User
(luoji@TEST.KRB5.COM)

kadmin.local -q "add_principal -pw luoji123 -kvno 1 luoji"


1.1 Creating keytab

ktutil
ktutil: addent -password -p luoji@TEST.KRB5.COM -k 1 -e aes256-cts-hmac-sha1-96
ktutil: wkt /keytabs/luoji.keytab 
ktutil: l

-- addent -password -p luoji1@TEST.KRB5.COM -k 1 -f


1.2 Query User

kadmin.local -q "listprincs"


1.3 Cache file

kinit -kt /keytabs/luoji.keytab luoji@TEST.KRB5.COM

klist -Aef

kdestroy


2. For Server
(SPN: HTTP/sso.test.gokrb5)


2.1 Add Server User(service principal name)

kadmin.local -q "addprinc -pw luoji123 -kvno 1  HTTP/sso.test.gokrb5"

2.2 Creating keytab

kadmin.local -q "ktadd -k /keytabs/sso.test.gokrb5.srv.keytab HTTP/sso.test.gokrb5@TEST.KRB5.COM"


kadmin.local -q "ktadd -k /keytabs/host.test.gokrb5.srv.keytab HTTP/host.test.gokrb5@TEST.KRB5.COM"


note: this way can't login by password


3. keytab file

klist -kt /keytabs/luoji.keytab

klist -kt /keytabs/sso.test.gokrb5.srv.keytab

