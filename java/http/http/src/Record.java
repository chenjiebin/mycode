package models.http.scale;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
/**
 *
 * @author chenjiebin
 */
public class Record extends AbstractScale {

    public Record() {
    }

    public void list() {
        String uri = "/api/record/list/?people_id=1&start_time=1&sso/93d9Wnr6y2CpjNdKgFlUNMQry56U5zxYtn8tObz6wFPB03ev/";
        Long start = System.currentTimeMillis();
        this.request(uri);
        Long end = System.currentTimeMillis();
        System.out.println(end - start);
    }
}
