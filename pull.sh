#!/bin/bash

mkdir -p var/tarball
mkdir -p var/data
mkdir -p var/sql
./wget_safe.sh -s "$@"/pub/ipfile.samp.tar.bz2 -t var/tarball/ipfile.samp.tar.bz2 -O var/data/ipfile.samp

cat > var/sql/updateNqmTargetAvailable.sql <<-EOM
UPDATE nqm_target SET tg_available=true
WHERE tg_host IN (
EOM
awk '{print "\""$2"\","}' var/data/ipfile.samp | sed '$s/,$//' >> var/sql/updateNqmTargetAvailable.sql
cat >> var/sql/updateNqmTargetAvailable.sql <<-EOM
);
EOM

cat > var/sql/updateNqmTargetStatus.sql <<-EOM
UPDATE nqm_target SET tg_status=true
WHERE tg_host IN (
EOM
awk '{print "\""$2"\","}' var/data/ipfile.samp | sed '$s/,$//' >> var/sql/updateNqmTargetStatus.sql
cat >> var/sql/updateNqmTargetStatus.sql <<-EOM
);
EOM
