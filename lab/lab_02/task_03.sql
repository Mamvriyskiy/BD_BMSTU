select distinct c.first_name, c.last_name, c.passport from client c 
where c.passport like '%123%'
