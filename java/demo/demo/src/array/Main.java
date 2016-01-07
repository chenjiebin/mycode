/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package array;

import java.util.ArrayList;

/**
 *
 * @author chenjiebin
 */
public class Main {

     public static void main(String[] args) {
        ArrayList List = new ArrayList();
        
        Goods goods = new Goods();
        goods.setGoods_id(570);
        goods.setGoods_num(1);
        List.add(goods);
        
        System.out.println(List);
    }
}
