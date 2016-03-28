// group by查询
// 参考地址：https://docs.mongodb.org/manual/reference/method/db.collection.group/

db.shoporders.group(
   {
     key: { user: 1 },
     cond: { created_at: { $gt: new Date( '01/01/2012' ) } },
     reduce: function ( curr, result ) {
     		result.total += 1;
     },
     initial: { total: 0 }
   }
)

