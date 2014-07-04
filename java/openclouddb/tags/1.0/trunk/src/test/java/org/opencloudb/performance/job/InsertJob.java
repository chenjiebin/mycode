package org.opencloudb.performance.job;

import java.util.List;
import java.util.Map;

import org.opencloudb.performance.CalcMain;
import org.opencloudb.performance.dao.InsertDAO;

public class InsertJob implements Runnable{
	private List<Map> list;
	@Override
	public void run(){
		InsertDAO dao = new InsertDAO();
		try {
			dao.insert(list);
		} catch (Exception e) {
			e.printStackTrace();
		}
		CalcMain.add(list.size());
	}
	public List<Map> getList() {
		return list;
	}
	public void setList(List<Map> list) {
		this.list = list;
	}
}
