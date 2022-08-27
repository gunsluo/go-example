#!/bin/bash

/opt/krb5/bin/krb5kdc-init.sh
/usr/sbin/kadmind &
/usr/sbin/krb5kdc -n
