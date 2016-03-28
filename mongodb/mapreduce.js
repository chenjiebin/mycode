db.runCommand(
   {
     mapReduce: "shoporders",
     map: function(){
        emit(this.user, {user: this.user})
     },
     reduce: function(key, value){
        var count = 0;
        value.forEach(function(){
            count+=1;
        });
        return count
     },
     // finalize: <function>,
     out: "order_people_num"
     // query: <document>,
     // sort: <document>,
     // limit: <number>,
     // scope: <document>,
     // jsMode: <boolean>,
     // verbose: <boolean>,
     // bypassDocumentValidation: <boolean>
   }
 )
db.order_people_num.find()