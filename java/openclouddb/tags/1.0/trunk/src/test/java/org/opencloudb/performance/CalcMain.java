package org.opencloudb.performance;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.opencloudb.performance.job.InsertJob;

/**
 * 1）连接MySQL集群中间件的8066端口
 * 2）获取采集数据库的数据将其插入到MySQL集群
 * 3） 取数据进行计算
 * 4）获得计算后的数据条数
 * @author shenzhw
 *
 */
public class CalcMain {
	private static int count;
	
	public static void add(int count){
		synchronized(CalcMain.class){
			CalcMain.count += count;
		}
	}
	
	public static void main(String[] args) throws Exception {
		
		ExecutorService pool = Executors.newFixedThreadPool(50);
		
		long start = System.currentTimeMillis();
		List<Map> list = new ArrayList<Map>();
		int c = 0;
		for(int i=0;i<1000000;i++){
			Map m = new HashMap();
			m.put("id", i);
			m.put("name", "meter" + i);
			if(i%4==0){
				m.put("org_no", "64401");
			}else if(i%4==1){
				m.put("org_no", "64402");
			}else if(i%4==2){
				m.put("org_no", "64403");
			}else if(i%4==3){
				m.put("org_no", "64404");
			}
			list.add(m);
			c += 1;
			if(c==1000){
				InsertJob job = new InsertJob();
				job.setList(list);
				pool.execute(job);
				c=0;
				list = new ArrayList<Map>();
			}
		}
		while(count < 1000000){
			System.out.println("已插入记录条数:" + count);
			Thread.sleep(3000);
		}
		System.out.println("最终记录路:" + count);
		System.out.println("共用时" + (System.currentTimeMillis() - start)/1000 + "秒");
	}
}
