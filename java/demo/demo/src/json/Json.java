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

    public void demo() {
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
