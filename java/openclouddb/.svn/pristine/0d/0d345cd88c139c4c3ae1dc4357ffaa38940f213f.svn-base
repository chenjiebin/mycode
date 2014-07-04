package org.opencloudb.performance.job;

import java.util.List;
import java.util.Map;

import org.opencloudb.performance.TestShardingHash;
import org.opencloudb.performance.dao.CustomerDAO;

public class CustomerJob implements Runnable{
	private List<Map> list;
	@Override
	public void run(){
		CustomerDAO dao = new CustomerDAO();
		try {
			dao.insert(list);
			TestShardingHash.add(list.size());
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
