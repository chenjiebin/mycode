package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	j := `[
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466006400)
            },
            "totalView" : 15
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466092800)
            },
            "totalView" : 2
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466179200)
            },
            "totalView" : 3
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466352000)
            },
            "totalView" : 2
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466438400)
            },
            "totalView" : 46
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466524800)
            },
            "totalView" : 147
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466611200)
            },
            "totalView" : 200
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466697600)
            },
            "totalView" : 89
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1466784000)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae0a9b91e6ee711fa3268",
                "day" : NumberLong(1467129600)
            },
            "totalView" : 6
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1466006400)
            },
            "totalView" : 8
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1466352000)
            },
            "totalView" : 3
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1466438400)
            },
            "totalView" : 17
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1466524800)
            },
            "totalView" : 69
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1466611200)
            },
            "totalView" : 97
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1466697600)
            },
            "totalView" : 39
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1466784000)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae165b91e6ee711fa3281",
                "day" : NumberLong(1467043200)
            },
            "totalView" : 5
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466006400)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466179200)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466352000)
            },
            "totalView" : 4
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466438400)
            },
            "totalView" : 19
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466524800)
            },
            "totalView" : 100
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466611200)
            },
            "totalView" : 126
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466697600)
            },
            "totalView" : 47
        }, 
        {
            "_id" : {
                "url" : "app://MallProduct/id=572ae190b91e6ee711fa328c",
                "day" : NumberLong(1466784000)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1434988800)
            },
            "totalView" : 3
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1464192000)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466006400)
            },
            "totalView" : 43
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466092800)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466179200)
            },
            "totalView" : 39
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466265600)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466352000)
            },
            "totalView" : 39
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466438400)
            },
            "totalView" : 336
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466524800)
            },
            "totalView" : 1063
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466611200)
            },
            "totalView" : 1319
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466697600)
            },
            "totalView" : 522
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466784000)
            },
            "totalView" : 5
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466870400)
            },
            "totalView" : 5
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1466956800)
            },
            "totalView" : 1
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1467043200)
            },
            "totalView" : 14
        }, 
        {
            "_id" : {
                "url" : "app://TabShop/",
                "day" : NumberLong(1467129600)
            },
            "totalView" : 2
        }
    ]`
	j = strings.Replace(j, "NumberLong(", "", -1)
	j = strings.Replace(j, ")", "", -1)

	data := []struct {
		Id struct {
			url string `json:"url"`
			day string `json:"day"`
		}
		TotalView int `json:"totalView"`
	}{}

	err := json.Unmarshal([]byte(j), &data)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(data)
}
