-- name: check
select 
  ees.plan_item_id , 
  sum(
    case
    	when ees.`type`='Nhập kho' then ees.total_bag
      else 0  
    end    
  ) as import,
    sum(
    case
      when ees.`type` in ('Xuất kho', 'Nhiễm') then ees.total_bag
      else 0  
    end    
  ) as export
from env_environment_storage ees
inner join env_environment_plan_item eepi on eepi.id = ees.plan_item_id 
inner join env_environment_plan eep on eep.id  = eepi.environment_plan_id
inner join env_environment_stock ees2 on ees2.id = eepi.environment_stock_id
where plan_item_id is not null
  and eep.date_work = ?
  and eepi.environment_stock_id = ?
  and eepi.`group`= ? 
group by plan_item_id 
order by plan_item_id desc;

-- name: check1
select 1;

