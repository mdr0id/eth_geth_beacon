#!/bin/bash

ufw deny 22/tcp
ufw allow 30303
ufw allow 13000/tcp
ufw allow 12000/udp
ufw enable
