/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package models.http.scale;

/**
 *
 * @author chenjiebin
 */
abstract public class AbstractScale extends models.http.AbstractHttp {

    public AbstractScale() {
        this.setHost("http://youxing.pba.cn:8002");
    }

    public void request(String uri) {
        super.request(uri);
    }
}
