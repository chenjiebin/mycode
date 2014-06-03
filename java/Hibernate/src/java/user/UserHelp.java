/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package user;

import java.util.List;
import org.hibernate.Query;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.cfg.Configuration;
import org.hibernate.service.ServiceRegistry;
import org.hibernate.service.ServiceRegistryBuilder;

/**
 *
 * @author chenjiebin
 */
public class UserHelp {

    Session session = null;

    public UserHelp() {
        Configuration configuration = new Configuration().configure();
        ServiceRegistry serviceRegistry
                = new ServiceRegistryBuilder().applySettings(configuration.getProperties()).buildServiceRegistry();

        SessionFactory sessionFactory = configuration.buildSessionFactory(serviceRegistry);

        this.session = sessionFactory.openSession();
    }

    public List getUser() {
        List<User> users = null;
        try {
            org.hibernate.Transaction tx = session.beginTransaction();
            Query q = session.createQuery("from User");
            users = (List<User>) q.list();
            tx.commit();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return users;
    }

}
