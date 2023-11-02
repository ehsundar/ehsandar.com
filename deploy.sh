#! /bin/bash

rsync -avz . htz:/root/ehsandar.com
ssh htz bash -c 'cd /root/ehsandar.com && /root/go/bin/hugo'
ssh htz bash -c '/root/go/bin/caddy reload --config /root/ehsandar.com/Caddyfile'
