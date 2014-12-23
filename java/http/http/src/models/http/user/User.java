package models.http.user;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
/**
 *
 * @author chenjiebin
 */
public class User extends AbstractUser {

    public User() {
    }

    public void getUserInfo() {
        String uri = "/api/my/info/sso/93d9Wnr6y2CpjNdKgFlUNMQry56U5zxYtn8tObz6wFPB03ev/";
        Long start = System.currentTimeMillis();
        this.request(uri);
        Long end = System.currentTimeMillis();
        System.out.println(end - start);
    }
}
