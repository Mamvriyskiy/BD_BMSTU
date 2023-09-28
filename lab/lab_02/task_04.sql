select c.first_name, c.last_name, c.sex from client c 
where c.first_name in ('Ivan', 'John') and c.last_name = 'Sato' 