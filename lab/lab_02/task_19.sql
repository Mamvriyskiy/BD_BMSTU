--task_19(Инструкция UPDATE со скалярным подзапросом в предложении SET)
update invoiceforpayment 
set total_price = (SELECT AVG(total_price) from invoiceforpayment)
where booking_id = 4;