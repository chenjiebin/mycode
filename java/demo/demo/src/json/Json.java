/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package json;

import com.google.gson.*;
import java.util.ArrayList;
import java.util.List;

//import net.sf.json.JSONObject;
/**
 *
 * @author chenjiebin
 */
public class Json {

    public void decode() {
        String json = "{\"data\":[{\"desc\":\"遇到要我保重的人，就逼他上秤！\",\"id\":\"1\",\"title\":\"每日测试\",\"type\":\"1\"},{\"desc\":\"我不是strong，只是虚胖！\",\"id\":\"2\",\"title\":\"详细分析\",\"type\":\"1\"},{\"desc\":\"完成任务，打倒绿茶婊虏获高富帅，从此走向人生巅峰！\",\"id\":\"3\",\"title\":\"我的任务包\",\"type\":\"1\"},{\"desc\":\"做一次是一夜情，长期来，才有资格说真爱\",\"id\":\"4\",\"title\":\"数据记录\",\"type\":\"1\"},{\"desc\":\"\",\"id\":\"5\",\"title\":\"订购\",\"type\":\"2\"},{\"desc\":\"\",\"id\":\"6\",\"title\":\"讨论\",\"type\":\"2\"},{\"desc\":\"\",\"id\":\"7\",\"title\":\"帮助\",\"type\":\"2\"}],\"errmsg\":\"\",\"errno\":\"0\"}";
        Gson gson = new Gson();
        Result result = gson.fromJson(json, Result.class);
        System.out.println(result.errno);
        System.out.println(result.errmsg);
    }

    public void encode() {
        Gson gson = new Gson();
        List<User> users = new ArrayList<User>();
        for (int i = 0; i < 10; i++) {
            User user = new User();
            user.setName("name" + i);
            user.setId("" + i);
            users.add(user);
        }
        String str = gson.toJson(users);

        System.out.println(str);
    }

}
