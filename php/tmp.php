<script>



    var temp = function(doc) {
        console.log("transformer: " + JSON.stringify(doc));

        var config = [
            {
                "ns": "foo.user",
                "fields": ["_id", "age"]
            },
            {
                "ns": "foo.order",
                "fields": ["_id", "order"]
            }
        ];

        var contains = function(array, item) {
            for (var i = 0; i < array.length; i++) {
                if (array[i] === item) {
                    return true;
                }
            }
            return false;
        }

        for (var i in config) {
            if (doc.ns === config[i].ns) {
                for (var field in doc.data) {
                    if (!contains(config[i].fields, field)) {
                        delete doc.data[field];
                    }
                }
                break;
            }
        }

        console.log("transformer: " + JSON.stringify(doc));
        return doc;
    }

    temp({"data": {"_id": {"$oid": "56fb983f5c0d0011f971e7f5"}, "order": "1234", "user": "cjb"}, "ns": "foo.order", "op": "insert", "ts": 1461121116});
</script>