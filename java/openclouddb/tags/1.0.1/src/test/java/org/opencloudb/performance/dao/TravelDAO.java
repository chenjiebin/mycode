package org.opencloudb.performance.dao;

import java.sql.PreparedStatement;
import java.util.List;
import java.util.Map;

import org.opencloudb.performance.dao.ConnectionManager.ConnLevel;

public class TravelDAO extends BaseDAO {
	public void insert(List<Map> list){
		PreparedStatement ps;
		try {
			this.getConnection(ConnLevel.normal);
			this.beginTransAction();
			String sql = "insert into travelrecord(id,user,traveldate,fee,days) values(?,?,curdate(),rand()*10000,floor(rand()*30))";
			ps = this.prepareStatement(sql);
			for (Map map : list) {
				ps.setInt(1, (Integer) map.get("id"));
				ps.setString(2, (String) map.get("user"));
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
