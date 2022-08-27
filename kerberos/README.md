

brew install krb5

##

```
KRB5_CONFIG=/Users/luoji/gopath/src/github.com/jcmturner/gokrb5-test/testenv/docker/krb5kdc/krb5.conf kinit luoji@TEST.GOKRB5

KRB5_CONFIG=/Users/luoji/gopath/src/github.com/jcmturner/gokrb5-test/testenv/docker/krb5kdc/krb5.conf kinit -kt /Users/luoji/gopath/src/github.com/jcmturner/gokrb5-test/testenv/docker/krb5kdc/keytabs/luoji.keytab luoji@TEST.GOKRB5

kinit -kt /Users/luoji/gopath/src/github.com/jcmturner/gokrb5-test/testenv/docker/krb5kdc/keytabs/luoji.keytab luoji@TEST.GOKRB5

kinit luoji@TEST.GOKRB5
```

##

```
./Google\ Chrome --auth-server-whitelist="*.test.gokrb5" --auth-negotiate-delegate-whitelist="*.test.gokrb5"
```
work



not work
```
defaults read com.google.Chrome
```

```
defaults write com.google.Chrome AuthServerWhitelist "*.test.gokrb5"
defaults write com.google.Chrome AuthNegotiateDelegateWhitelist "*.test.gokrb5"

defaults write com.google.Chrome DisableAuthNegotiateCnameLookup true
```

```
defaults delete com.google.Chrome AuthNegotiateDelegateWhitelist
defaults delete com.google.Chrome AuthServerWhitelist
defaults delete com.google.Chrome DisableAuthNegotiateCnameLookup
```
