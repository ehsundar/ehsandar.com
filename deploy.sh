#! /bin/bash

hugo
rm -rf /var/www/ehsandar.top/public
mv public /var/www/ehsandar.top/public
