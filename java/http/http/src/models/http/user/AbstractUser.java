/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package models.http.user;

/**
 *
 * @author chenjiebin
 */
abstract public class AbstractUser extends models.http.AbstractHttp {
    
    public AbstractUser() {
        this.setHost("http://app.pba.cn");
    }
    
    public void request(String uri) {
        super.request(uri);
    }
}
