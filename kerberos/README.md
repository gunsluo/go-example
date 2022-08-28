

brew install krb5

##

```
cp -fr docker/keytabs/krb5.conf /etc/krb5.conf

KRB5_CONFIG=docker/keytabs/krb5.conf kinit luoji@TEST.KRB5.COM

KRB5_CONFIG=docker/keytabs/krb5.conf kinit -kt docker/keytabs/luoji.keytab luoji@TEST.KRB5.COM

kinit -kt docker/keytabs/luoji.keytab luoji@TEST.KRB5.COM

kinit luoji@TEST.KRB5.COM
```

note: please use system's kinit

##

```
cd /Applications/Google\ Chrome.app/Contents/MacOS
./Google\ Chrome --auth-server-whitelist="*.test.krb5.com" --auth-negotiate-delegate-whitelist="*.test.krb5.com"
```
work



not work
```
defaults read com.google.Chrome
```

```
defaults write com.google.Chrome AuthServerWhitelist "*.test.krb5.com"
defaults write com.google.Chrome AuthNegotiateDelegateWhitelist "*.test.krb5.com"

defaults write com.google.Chrome DisableAuthNegotiateCnameLookup true
```

```
defaults delete com.google.Chrome AuthNegotiateDelegateWhitelist
defaults delete com.google.Chrome AuthServerWhitelist
defaults delete com.google.Chrome DisableAuthNegotiateCnameLookup
```
