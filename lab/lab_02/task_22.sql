with BirthClient as (select * from client where date_of_birth BETWEEN '2000-01-01' and '2001-01-01')
select * from BirthClient
