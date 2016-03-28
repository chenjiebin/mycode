// 大于小于参考
start = new Date("01/03/2016 00:00:00");
end = new Date("01/04/2016 00:00:00");
db.getCollection('shoporders').find({
    "list.commodity":ObjectId("56b0335ac9bad9c91442d783"),
    "created_at": {"$gte" : start, "$lte" : end}
    })
