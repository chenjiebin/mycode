/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package http;

import models.http.scale.Record;
import models.http.user.*;
import models.http.scale.*;

/**
 *
 * @author chenjiebin
 */
public class Http {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        System.out.println("hehe");
        User user = new User();
        user.getUserInfo();
        
        Record record = new Record();
        record.list();
    }

}
