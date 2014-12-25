/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package demo;

//import json.Json;
import volley.VolleyDemo;

/**
 *
 * @author chenjiebin
 */
public class Demo {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
//        Json json = new Json();
//        json.encode();
//        json.decode();

        VolleyDemo volley = new VolleyDemo();
        volley.httpGet();

    }

}
