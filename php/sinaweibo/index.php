<?php

//用来抓取微博转发的粉丝
set_time_limit(0);
//getAllUser();
filterUser();

function getAllUser() {
    for ($page = 1; $page <= 50; $page++) {
        $url    = "http://weibo.com/aj/v6/mblog/info/big?ajwvr=6&id=3887097863700268&max_id=3888379865695877&page=$page&__rnd=1442547658610";
        echo $url, "\n";
        $result = shell_exec("curl --compressed '$url' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3' -H 'Cache-Control: no-cache' -H 'Connection: keep-alive' -H 'Cookie: SINAGLOBAL=9381130713587.125.1400259077579; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WhHL7Ieb9S4xqk8hLuU6yTg5JpX5KMt; SUHB=0IIkSr3BwvsYT3; SINAGLOBAL=; SUBP=; SUHB=; ALF=; ULV=; UOR=; UOR=login.sina.com.cn,weibo.com,login.sina.com.cn; YF-Page-G0=f27a36a453e657c2f4af998bd4de9419; SUS=SID-1934495122-1442545657-GZ-a3zrp-48df0fdfe8bbb9478ec748fda54c4900; SUE=es%3Dde354c6371355231c06ccabef29b8139%26ev%3Dv1%26es2%3D582597abbb8c3a97d37529ff2883a730%26rs0%3DF4F5%252BrlzkiMyl4BIpo%252FuQDCTh7pkTqH5Zc0UnogjldMRjPjqgRETMSiSFXKQmzxbiqSrMytI%252FcE5JqcsAuiDezq%252FWo8%252BTopPfKrhNGMn0DOdEGK773p3QXFfbrgWk7H0C0xLYiX4IEllb2P46RbrkHN%252FRqSCwvgmgvjYCKZ%252FneE%253D%26rv%3D0; SUP=cv%3D1%26bt%3D1442545657%26et%3D1442632057%26d%3Dc909%26i%3D4900%26us%3D1%26vf%3D0%26vt%3D0%26ac%3D2%26st%3D0%26uid%3D1934495122%26name%3D303410541%2540qq.com%26nick%3D%25E4%25B8%258A%25E5%25B8%259D%25E6%2589%2593%25E6%2589%258B%26fmp%3D%26lcp%3D2012-07-02%252014%253A13%253A29; SUB=_2A254_w-pDeTxGedH6FYV-SvNyT6IHXVbjWZhrDV8PUNbvtBeLXjdkW9zUGVhiulQwlbf2pyLHG8VKwB0Mg..; ALF=1474081656; SSOLoginState=1442545657; YF-Ugrow-G0=ad06784f6deda07eea88e095402e4243; YF-V5-G0=55f24dd64fe9a2e1eff80675fb41718d; _s_tentry=login.sina.com.cn; Apache=7044914637696.234.1442545678439; ULV=1442545678458:1:1:1:7044914637696.234.1442545678439:' -H 'Host: weibo.com' -H 'Pragma: no-cache' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:39.0) Gecko/20100101 Firefox/39.0'");
        $rs     = json_decode($result, true);
        if ($rs['code'] != "100000") {
            echo "page " . $page . " request error!", "\n";
            continue;
        }

        preg_match_all('/<a.+?node-type="name">(.+?)<\/a>/', $rs['data']['html'], $matches);
        if (!$matches[0]) {
            echo "page " . $page . " match error!", "\n";
            continue;
        }
        $f = fopen("./data/$page.log", "w");
        foreach ($matches[1] as $v) {
            fwrite($f, $v . "\n");
        }
        fclose($f);
    }
}

function filterUser() {
    $users = array();

    $dir    = "./data/";
    $handle = opendir($dir);
    while (($file   = readdir($handle)) !== false) {
        //排除掉当前目录和上一个目录
        if (substr($file, 0, 1) == ".") {
            continue;
        }
        $file = $dir . DIRECTORY_SEPARATOR . $file;

        //如果是文件就打印出来，否则递归调用
        $f    = fopen($file, "r");
        while ($read = fgets($f)) {
            if (!in_array($read, $users)) {
                $users[] = trim($read);
            }
        }
    }

    $f = fopen("./data/result.log", "w");
    fwrite($f, implode("\r\n", $users));
    fclose($f);
}
