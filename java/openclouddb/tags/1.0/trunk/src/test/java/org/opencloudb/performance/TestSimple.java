package org.opencloudb.performance;

import java.sql.Connection;
import java.sql.ResultSet;

import org.opencloudb.performance.dao.ConnectionManager;

public class TestSimple {

	public static void main(String[] args) throws Exception {
		Connection con = ConnectionManager.getNormalConnection();
		ResultSet rs = con.createStatement().executeQuery(
				"select * from company");
		while (rs.next()) {
			System.out.println(" " + rs.getString(1));

		}
		con.close();
	}

}
