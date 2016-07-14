'use strict';




module.exports = function (router) {
    // 测试json转成string
    router.get('/test', function (req, res) {
        var order = [
            {
                "supplier_id": 2,
                "addr_info": {
                    "province": "福建省",
                    "city": "厦门市",
                    "district": "思明区",
                    "address": "观日路52号楼",
                    "zip": "364000",
                    "receiver": "陈杰斌",
                    "receiver_phone": "13860171761",
                    "id_card_num": "3303XXXXXXXXXX53",
                    "id_card_photo_back": "",
                    "id_card_photo_type": ""
                },
                "total_fee": 27496,
                "official_fee": 0,
                "transshipment_fee": 0,
                "goods_fee": 27496,
                "uid": "1",
                "note": "备注信息",
                "out_trade_no": "1",
                "apply_time": "2016-07-12 14:00:00",
                "attach_data": "这个数据回回传",
                "notify_url": "http://ts-pay859.kuiyinapp.com/duoshoubang/notify",
                "items": [
                    {
                        "goods_id": 100002,
                        "product_id": 10001,
                        "goods_price": 27496,
                        "quantity": "1",
                        "goods_ver": "0.01",
                        "ext_data": "扩展数据"
                    }
                ]
            }
        ];
        var orderString = JSON.stringify(order);
        console.log(orderString);
        console.log(JSON.parse(orderString));

        return res.end(orderString);
    });
};
