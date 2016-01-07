/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package fastjson;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;

import org.json.JSONException;

import java.nio.*;

/**
 *
 * @author chenjiebin
 */
public class Main {

    public static void main(String[] args) {
        System.out.println(ByteOrder.nativeOrder());
        
        ArrayList List = new ArrayList();
        Goods goods = new Goods();
        goods.setA("1");
        goods.setB("570");
        goods.setC("570");
        goods.setD("570");
        goods.setE("570");
        goods.setF("570");
        List.add(goods);

        JSONObject json = new JSONObject();
        json.put("type", "1");
        json.put("goods", List);
        json.put("app_in_source", "");
        System.out.println(json.toString());

        Map<String, Object> map1 = (Map<String, Object>) JSON.parse(json.toString());
        System.out.println(map1);

        Map map2 = parserToMap(json.toString());
        System.out.println(map2);

        org.json.JSONObject j = new org.json.JSONObject(json);
        try {
            System.out.println(j.get("goods").toString());
        } catch (JSONException ex) {
            Logger.getLogger(Main.class.getName()).log(Level.SEVERE, null, ex);
        }
    }

    public static Map parserToMap(String s) {
        Map map = new HashMap();
        org.json.JSONObject json = null;
        try {
            json = new org.json.JSONObject(s);
        } catch (JSONException e) {
            e.printStackTrace();
        }
        System.out.println(json);
        Iterator keys = json.keys();
        while (keys.hasNext()) {
            String key = (String) keys.next();
            String value = null;
            try {
                value = json.get(key).toString();
            } catch (JSONException e) {
                e.printStackTrace();
            }
            if (value.startsWith("{") && value.endsWith("}")) {
                map.put(key, parserToMap(value));
            } else {
                map.put(key, value);
            }

        }
        return map;
    }
}
