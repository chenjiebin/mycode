package org.opencloudb.performance.dao;

import java.sql.PreparedStatement;
import java.util.List;
import java.util.Map;

import org.opencloudb.performance.dao.ConnectionManager.ConnLevel;

public class CustomerDAO extends BaseDAO{
	public void insert(List<Map> list){
		PreparedStatement ps;
		try {
			this.getConnection(ConnLevel.normal);
			this.beginTransAction();
			String sql = "insert into customer(id,name,company_id,sharding_id) values(?,?,?,?)";
			ps = this.prepareStatement(sql);
			for (Map map : list) {
				ps.setInt(1, (Integer) map.get("id"));
				ps.setString(2, (String) map.get("name"));
				ps.setString(3, (String) map.get("company_id"));
				ps.setString(4, (String)map.get("sharding_id"));
				ps.addBatch();
				ps.executeBatch();
				this.commitTransAction();
			}
		} catch (Exception e) {
			e.printStackTrace();
		} finally{
			this.close();
		}
	}
}
