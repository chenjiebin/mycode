// 大于小于参考
start = new Date("03/01/2016 00:00:00");
end = new Date("05/01/2016 00:00:00");
db.getCollection('shoporders').find({
    "list.commodity":ObjectId("56b0335ac9bad9c91442d783"),
    "created_at": {"$gte" : start, "$lte" : end}
    })
