/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package json;

import java.util.ArrayList;
//import java.util.HashMap;
import java.util.List;
//import java.util.Map;
import net.sf.json.JSONArray;
//import net.sf.json.JSONObject;

/**
 *
 * @author chenjiebin
 */
public class Json {

    public void demo() {
        List list = new ArrayList();
        list.add("first");
        list.add("second");
        JSONArray jsonArray2 = JSONArray.fromObject(list);

        System.out.println(jsonArray2);
    }

//    public void demo(){
//        // JSON格式数据解析对象
//        JSONObject jo = new JSONObject();
//
//        // 下面构造两个map、一个list和一个Employee对象
//        Map<String, String> map1 = new HashMap<String, String>();
//        map1.put("name", "Alexia");
//        map1.put("sex", "female");
//        map1.put("age", "23");
//
//        Map<String, String> map2 = new HashMap<String, String>();
//        map2.put("name", "Edward");
//        map2.put("sex", "male");
//        map2.put("age", "24");
//
//        List<Map> list = new ArrayList<Map>();
//        list.add(map1);
//        list.add(map2);
//
////        Employee employee = new Employee();
////        employee.setName("wjl");
////        employee.setSex("female");
////        employee.setAge(24);
//
//        // 将Map转换为JSONArray数据
//        JSONArray ja1 = JSONArray.fromObject(map1);
//        // 将List转换为JSONArray数据
//        JSONArray ja2 = JSONArray.fromObject(list);
//        // 将Bean转换为JSONArray数据
////        JSONArray ja3 = JSONArray.fromObject(employee);
//
//        System.out.println("JSONArray对象数据格式：");
//        System.out.println(ja1.toString());
//        System.out.println(ja2.toString());
////        System.out.println(ja3.toString());
//
//        // 构造Json数据，包括一个map和一个Employee对象
////        jo.put("map", ja1);
////        jo.put("employee", ja2);
////        System.out.println("\n最终构造的JSON数据格式：");
////        System.out.println(jo.toString());
//
////        return jo.toString();
//    }
}
