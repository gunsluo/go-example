#!/bin/bash

export REALM=${KRB5_REALM:-TEST.GOKRB5}
export DOMAIN=${KRB5_DOMAIN:-test.gokrb5}
KDC_HOST=kdc.${DOMAIN}
ADMIN_USERNAME=adminuser
HOST_PRINCIPALS=${HOST_PRINCIPALS:-${KDC_HOST}}
SPNS=${SPNS}
USERS=${USERS}

USER_PASSWORD=${USER_PASSWORD:-password}
HOST_PASSWORD=${HOST_PASSWORD:-password}
SRV_PASSWORD=${SRV_PASSWORD:-password}

echo "-> ${REALM}"
echo "-> ${DOMAIN}"
echo "-> ${HOST_PRINCIPALS}"
echo "-> ${SPNS}"

# render config template
cat /cfg-tmpl/krb5.conf | envsubst > /etc/krb5.conf
cat /cfg-tmpl/kdc.conf | envsubst > /var/kerberos/krb5kdc/kdc.conf
cat /cfg-tmpl/kadm5.acl | envsubst > /var/kerberos/krb5kdc/kadm5.acl
cp -fr /etc/krb5.conf /keytabs


create_entropy() {
   while true
   do
     sleep $(( ( RANDOM % 10 )  + 1 ))
     echo "Generating Entropy... $RANDOM"
   done
}

create_entropy &
ENTROPY_PID=$!
echo "-> ${ENTROPY_PID}"


echo "Kerberos initialisation required. Creating database for ${REALM} ..."
echo "This can take a long time if there is little entropy. A process has been started to create some."

MASTER_PASSWORD=$(echo $RANDOM$RANDOM$RANDOM | md5sum | awk '{print $1}')
echo "MASTER PASSWORD: ${MASTER_PASSWORD}"

/usr/sbin/kdb5_util create -r ${REALM} -s -P ${MASTER_PASSWORD}

kill -9 ${ENTROPY_PID}

echo "Kerberos database created."

/usr/sbin/kadmin.local -q "add_principal -randkey ${ADMIN_USERNAME}/admin"
echo "Kerberos admin user created: ${ADMIN_USERNAME} To update password: sudo /usr/sbin/kadmin.local -q \"change_password ${ADMIN_USERNAME}/admin\""


KEYTAB_DIR="/keytabs"
mkdir -p $KEYTAB_DIR


if [ ! -z "${HOST_PRINCIPALS}" ]; then
  for host in ${HOST_PRINCIPALS}
  do
    /usr/sbin/kadmin.local -q "add_principal -pw ${HOST_PASSWORD} -kvno 1 host/$host"
    echo "Created host principal host/$host"
  done
fi


if [ ! -z "${SPNS}" ]; then
  for spn in ${SPNS}
  do
    /usr/sbin/kadmin.local -q "add_principal -pw ${SRV_PASSWORD} -kvno 1 HTTP/$spn"
    echo "Created server principal HTTP/$spn"
    /usr/sbin/kadmin.local -q "ktadd -k /keytabs/$spn.svc.keytab HTTP/$spn@${REALM}"
    echo "Created server keytab /keytabs/$spn.svc.keytab"
  done
fi


/usr/sbin/kadmin.local -q "add_principal -pw password -kvno 1 DNS/ns.${DOMAIN}"

# default users
/usr/sbin/kadmin.local -q "add_principal -pw ${USER_PASSWORD} -kvno 1 testuser1"
/usr/sbin/kadmin.local -q "add_principal +requires_preauth -pw ${USER_PASSWORD} -kvno 1 testuser2"
/usr/sbin/kadmin.local -q "add_principal -pw ${USER_PASSWORD} -kvno 1 testuser3"

# Set up trust
/usr/sbin/kadmin.local -q "add_principal -requires_preauth -pw trustpasswd -kvno 1 krbtgt/${REALM}@RESDOM.GOKRB5"
/usr/sbin/kadmin.local -q "add_principal -requires_preauth -pw trustpasswd -kvno 1 krbtgt/RESDOM.GOKRB5@${REALM}"

if [ ! -z "${USERS}" ]; then
  for user in ${USERS}
  do
    /usr/sbin/kadmin.local -q "add_principal -pw ${USER_PASSWORD} -kvno 1 $user"
    echo "Created server principal $user"
  done
fi


echo "Kerberos initialisation complete"

