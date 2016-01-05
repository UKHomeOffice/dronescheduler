#!/bin/bash
UPDATE=0
CRONTAB=/etc/crontabs/root
CROND=/usr/sbin/crond
CROND_FLAGS="-f"

echo '' > $CRONTAB

for item in `env`
do
   case "$item" in
       CRONTASK_*)
            ENVVAR=`echo $item | cut -d \= -f 1`
            printenv $ENVVAR >> $CRONTAB
            UPDATE=1
            ;;
   esac
done

if [ $UPDATE == 1 ]; then
    echo "root" > /etc/crontabs/cron.update
fi

exec $CROND $CROND_FLAGS
