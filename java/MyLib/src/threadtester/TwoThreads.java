/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package threadtester;

/**
 *
 * @author chenjiebin
 */
public class TwoThreads {

    public void go() {
        MyRunnable runner = new MyRunnable();
        Thread alpha = new Thread(runner);
        Thread beta = new Thread(runner);
        alpha.setName("Alpha thread");
        beta.setName("Beta thread");
        alpha.start();
        beta.start();
    }

    public static void main(String[] args) {
        TwoThreads twoThreads = new TwoThreads();
        twoThreads.go();
    }

    public class MyRunnable implements Runnable {

        @Override
        public void run() {
            for (int i = 0; i < 25; i++) {
                String threadName = Thread.currentThread().getName();
                System.out.println(threadName + " is running");
            }
        }

    }
}
