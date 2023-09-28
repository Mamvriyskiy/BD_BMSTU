--task_13(Инструкция SELECT, использующая вложенные подзапросы с уровнем вложенности 3)
select * from booking 
where booking_id in (Select booking_id from booking b
                    where b.count_room = (select MAX(count_room) from booking))
