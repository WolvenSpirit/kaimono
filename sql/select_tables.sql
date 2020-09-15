SELECT table_name
  FROM information_schema.tables
 WHERE table_schema='public'
 AND table_name != $1 
 AND table_type='BASE TABLE';