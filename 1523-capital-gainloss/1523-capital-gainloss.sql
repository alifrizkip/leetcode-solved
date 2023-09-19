# Write your MySQL query statement below
select
    stock_name,
    sum(gain) capital_gain_loss
from (
    select
        stock_name,
        case
            when operation = "Buy" then (price * -1)
            else price
        end gain
    from Stocks
) d
group by stock_name