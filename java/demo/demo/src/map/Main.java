/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package map;

import java.util.HashMap;
import java.util.Map;
import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import java.util.LinkedHashMap;
/**
 *
 * @author chenjiebin
 */
public class Main {

    public static void main(String[] args) {
        Map map = new HashMap();
        map.put("a", "570");
        map.put("c", "1");
        map.put("b", "1");
        
        map.put("d", "1");
        map.put("e", "1");
        map.put("f", "1");
        map.put("g", "1");
        map.put("h", "1");
        
        System.out.println(map);
    }
}
