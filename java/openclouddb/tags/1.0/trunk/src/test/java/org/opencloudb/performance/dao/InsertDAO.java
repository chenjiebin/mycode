package org.opencloudb.performance.dao;

import java.sql.PreparedStatement;
import java.util.List;
import java.util.Map;

import org.opencloudb.performance.dao.ConnectionManager.ConnLevel;

public class InsertDAO extends BaseDAO{
	public void insert(List<Map> list){
		PreparedStatement ps;
		try {
			this.getConnection(ConnLevel.normal);
			String sql = "insert into meters(id,name,org_no) values(?,?,?)";
			ps = this.prepareStatement(sql);
			int c = 0;
			for (Map map : list) {
				ps.setInt(1, (Integer) map.get("id"));
				ps.setString(2, (String) map.get("name"));
				ps.setString(3, (String) map.get("org_no"));
				ps.addBatch();
				c += 1;
				if (c == 1000) {
					ps.executeBatch();
					c = 0;
				}
			}
			ps.executeBatch();
		} catch (Exception e) {
			e.printStackTrace();
		} finally{
			this.close();
		}
	}
}
