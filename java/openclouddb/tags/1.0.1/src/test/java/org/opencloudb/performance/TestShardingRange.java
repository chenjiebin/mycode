package org.opencloudb.performance;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.opencloudb.performance.job.TravelJob;

/**
 * 测试按数据范围分片功能
 * 
 * @author shenzhw
 *
 */
public class TestShardingRange {
	private static int count;
	
	public static void add(int count){
		synchronized(TestShardingHash.class){
			TestShardingRange.count += count;
		}
	}
	
	public static void main(String[] args)  throws Exception {
		int threadCount = 0;//线程数
		int recordCount = 0;//要插入的记录数
		try{
			threadCount = Integer.parseInt(args[0]);
			recordCount = Integer.parseInt(args[1]);
		}catch(NumberFormatException e){
			System.out.println("请输入正确的参数,格式为:[线程池大小] [记录数]");
			BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
			try {
				String params = br.readLine();
				threadCount = Integer.parseInt(params.split(" ")[0]);
				recordCount = Integer.parseInt(params.split(" ")[1]);
			} catch (Exception e2) {
				e2.printStackTrace();
			}
		}
		System.out.println("线程池大小:" + threadCount);
		System.out.println("插入记录数:" + recordCount);
		
		ExecutorService pool = Executors.newFixedThreadPool(threadCount);
		
		long start = System.currentTimeMillis();
		List<Map> list = new ArrayList<Map>();
		int c = 0;//统计个数，当个数达到5000后提交
		for(int i=0;i<recordCount;i++){
			Map m = new HashMap();
			m.put("id", i);
			m.put("user", "user" + i);
			list.add(m);
			c += 1;
			if(c==5000){
				TravelJob job = new TravelJob();
				job.setList(list);
				pool.execute(job);
				c=0;
				list = new ArrayList<Map>();
			}
		}
		System.out.println("已分配完任务，耗时:" + (System.currentTimeMillis() - start)/1000 + "秒");
		while(count < recordCount){
			System.out.println("已插入记录条数:" + count);
			Thread.sleep(1000);
		}
		System.out.println("最终记录路:" + count);
		pool.shutdown();
		System.out.println("共用时" + (System.currentTimeMillis() - start)/1000 + "秒");
	}
}
