/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package user;

import org.hibernate.SessionFactory;
import org.hibernate.cfg.Configuration;

/**
 * Hibernate Utility class with a convenient method to get Session Factory
 * object.
 *
 * @author chenjiebin
 */
public class HibernateUtil {

    public static SessionFactory getSessionFactory() {
        try {
            // Create the SessionFactory from standard (hibernate.cfg.xml) 
            // config file.
            SessionFactory sessionFactory = new Configuration().configure().buildSessionFactory();

            return sessionFactory;

        } catch (Throwable ex) {
            // Log the exception. 
            System.err.println("Initial SessionFactory creation failed." + ex);
            throw new ExceptionInInitializerError(ex);
        }
    }
}
