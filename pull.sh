#!/bin/bash

mkdir -p var/tarball
mkdir -p var/data
mkdir -p var/sql
./wget_safe.sh -s "$@"/pub/ipfile.samp.tar.bz2 -t var/tarball/ipfile.samp.tar.bz2 -O var/data/ipfile.samp

cat > var/sql/updateNqmTargetAvailable.sql <<-EOM
SET NAMES 'utf8';
UPDATE nqm_target SET tg_available=true
WHERE tg_host IN (
EOM
awk '{print "\""$2"\","} END{print "\""$2"\""}' var/data/ipfile.samp >> var/sql/updateNqmTargetAvailable.sql
cat >> var/sql/updateNqmTargetAvailable.sql <<-EOM
);
EOM

cat > var/sql/updateNqmTargetStatus.sql <<-EOM
SET NAMES 'utf8';
UPDATE nqm_target SET tg_status=true
WHERE tg_host IN (
EOM
awk '{print "\""$2"\","} END{print "\""$2"\""}' var/data/ipfile.samp >> var/sql/updateNqmTargetStatus.sql
cat >> var/sql/updateNqmTargetStatus.sql <<-EOM
);
EOM
