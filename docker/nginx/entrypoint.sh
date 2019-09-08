#!/bin/ash

if [ -z "$MF_MQTT_CLUSTER" ]
then
      envsubst < /etc/nginx/snippets/mqtt-upstream-single.conf > /etc/nginx/snippets/mqtt-upstream.conf
      envsubst < /etc/nginx/snippets/mqtt-ws-upstream-single.conf > /etc/nginx/snippets/mqtt-ws-upstream.conf
else
      envsubst < /etc/nginx/snippets/mqtt-upstream-cluster.conf > /etc/nginx/snippets/mqtt-upstream.conf
      envsubst < /etc/nginx/snippets/mqtt-ws-upstream-cluster.conf > /etc/nginx/snippets/mqtt-ws-upstream.conf
fi

envsubst < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

exec nginx -g "daemon off;"
