--task_02 (список клиентов, у которых дата рождения между ...)
select distinct c.first_name, c.last_name, c.date_of_birth from client c 
where c.date_of_birth between '1998-01-01' and '1998-05-01'
