--task_11(Создание новой временной локальной таблицы из результирующего набора данных инструкции SELECT.)
select distinct c.first_name, c.last_name, c.date_of_birth, now(), now() - interval '15 days'
INTO NowMonthBirthday
from client c 
where EXTRACT(MONTH FROM c.date_of_birth) = EXTRACT(MONTH FROM now())