'use strict';

var order = [
    [
        {
            "supplier_id": 2,
            "addr_info": {
                "address": "观日路52号楼",
                "city": "厦门市",
                "district": "思明区",
                "id_card_num": "3303XXXXXXXXXX53",
                "id_card_photo_back": "",
                "id_card_photo_type": "",
                "province": "福建省",
                "receiver": "陈杰斌",
                "receiver_phone": "13860171761",
                "zip": "364000"
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
                    "product_id": 10001,
                    "goods_id": 100002,
                    "goods_price": "27496",
                    "goods_ver": "1",
                    "quantity": "1",
                    "ext_data": "扩展数据"
                }
            ]
        }
    ]
];
console.log(JSON.stringify(order));


module.exports = function (router) {
    // 测试json转成string
    router.get('/test', function (req, res) {
        var order = [
            {
                "order_id": "20160712164848",
                "goods": [
                    {
                        "goods_id": "1"
                    }
                ]
            }
        ];
        return res.end(JSON.stringify(order));
    });
};
