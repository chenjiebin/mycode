package org.opencloudb.performance.job;

import java.util.List;
import java.util.Map;

import org.opencloudb.performance.TestShardingRange;
import org.opencloudb.performance.dao.TravelDAO;

public class TravelJob implements Runnable {

	private List<Map> list;
	@Override
	public void run(){
		TravelDAO dao = new TravelDAO();
		try {
			dao.insert(list);
			TestShardingRange.add(list.size());
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	public List<Map> getList() {
		return list;
	}
	public void setList(List<Map> list) {
		this.list = list;
	}

}
