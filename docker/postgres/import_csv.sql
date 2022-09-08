COPY markets FROM '/tmp/DEINFO_AB_FEIRASLIVRES_2014.csv' CSV HEADER;

SELECT setval(pg_get_serial_sequence('markets', 'id'), (SELECT MAX(id) FROM markets)+1);