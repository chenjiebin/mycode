/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package sockettester;

import java.io.*;
import java.net.*;
import java.util.Locale;

/**
 *
 * @author chenjiebin
 */
public class DailyAdviceServer {

    String[] adviceList = {"1", "2", "3"};

    public void go() {
        try {
            ServerSocket serverSock = new ServerSocket(4242);

            while (true) {
                Socket sock = serverSock.accept();

                InputStreamReader streamReader = new InputStreamReader(sock.getInputStream());
                BufferedReader reader = new BufferedReader(streamReader);

                String info = reader.readLine();
                System.out.println("accept from client:" + info);

                PrintWriter writer = new PrintWriter(sock.getOutputStream());
                String advice = getAdvice();
                writer.println(advice);
                writer.close();
                System.out.println(advice);
            }
        } catch (IOException ex) {

        }

    }

    public String getAdvice() {
        int random = (int) (Math.random() * adviceList.length);
        return adviceList[random];
    }

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        DailyAdviceServer server = new DailyAdviceServer();
        server.go();
    }
}
