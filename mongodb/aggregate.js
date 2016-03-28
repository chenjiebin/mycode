

db.shoporders.aggregate(
    [
        { $match: { "list.commodity": ObjectId("56b03352c9bad9c91442d782") } },
        { $group: { _id: "$user", count: { $sum: 1 } } },
        { $sort: { count: -1 } }
    ]                   
)
select user,count(0) from orders where ... group by user 


select count(0) from ( select user from orders where ... group by user )
db.shoporders.aggregate(
    [
        { $match: { "list.commodity": ObjectId("56b03352c9bad9c91442d782") } },
        { $group: { _id: "$user" } },
        { $group: { _id: null, count: { $sum: 1 } } }
    ]                   
)


db.shoporders.aggregate(
    [
        { $group: { _id: {commodity: "$list.commodity", user: "$user"}, count: { $sum: 1 } } }
    ]                   
)