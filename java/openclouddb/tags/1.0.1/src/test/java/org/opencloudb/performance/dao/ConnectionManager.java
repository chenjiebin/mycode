package org.opencloudb.performance.dao;

import java.beans.PropertyVetoException;
import java.io.FileInputStream;
import java.io.IOException;
import java.net.URLDecoder;
import java.sql.Connection;
import java.util.Properties;

import com.mchange.v2.c3p0.ComboPooledDataSource;
import com.mchange.v2.c3p0.DataSources;

/**
 * C3P0数据库连接池
 * 
 * @author zhangjp
 * @version 1.0 2009/08/27
 * 
 */
public final class ConnectionManager {
	/**
	 * 数据库连接级别
	 * 
	 */
	public enum ConnLevel {
		/** 普通连接 */
		normal,
		/** Vip连接 */
		vip
	}
	/**
	 * 数据库模式
	 */
	public static String dbShcema ="task.";
	/** 全局数据库连接池（整个应用） */
	@SuppressWarnings("unused")
	private final static ConnectionManager instance = new ConnectionManager();

	/** 数据源 */
	private static ComboPooledDataSource ds;

	/** vip数据源 */
	private static ComboPooledDataSource dsVip;

	/** 连接池的初始大小 */
	private int initialConnections;

	/** 连接池中保留的最小连接数 */
	private int minPoolSize;

	/** 连接池自动增加的大小 */
	private int incrementalConnections;

	/** 连接池最大的大小 */
	private int maxConnections;

	/** 数据库驱动 */
	private String jdbcDriver = "";

	/** 数据 URL */
	private String dbUrl = "";

	/** 数据库用户名 */
	private String dbUsername = "";

	/** 数据库用户密码 */
	private String dbPassword = "";

	/** 间隔多长时间检查所有数据库连接池中的空闲连接 */
	private int idleConnectionTestPeriod;

	/** 空闲连接的最长保存时间 */
	private int freeMaxTime;

	/** vip连接数 */
	private int vipConnSize;

	/** 连接关闭时默认将所有未提交的操作回滚。Default: false */
	private boolean autoCommitOnClose;

	/** 在从数据库获取新连接失败后重复尝试的次数。 */
	private int acquireRetryAttempts;

	/** 两次连接的间隔时间 单位毫秒 */
	private int acquireRetryDelay;

	/** 获取连接失败线程抛出异常时，设定数据源是否依旧有效，false：有效，true：关闭数据源 */
	private boolean breakAfterAcquireFailure;

	/** 用来测试连接是否可用的表名，其可以由用户指定，c3p0自动创建。 */
//	private String automaticTestTable;

	/**
	 * 获取获取每一个连接时，是否测试连接是否可用，default：false 不测试 true：测试。
	 */
	private boolean testConnectionOnCheckin;

	/**
	 * 当连接池的连接用完时客户端调用getConnection()后等待获取新连接的时间，
	 * 超时后将抛出SQLException,如设为0则无限期等待。单位毫秒。Default: 0
	 */
	private int checkoutTimeout;

	/** 用户修改系统配置参数执行前最多等待时间 单位：秒 */
	private int propertyCycle;

	/** 普通连接调用次数 */
	private static int callNormalConnCount;

	/** 普通连接调用百万计数 */
	private static int callNormalConnMillionCount;

	/** VIP连接调用次数 */
	private static int callVipConnCount;

	/** VIP连接调用百万计数 */
	private static int callVipConnMillionCount;

	/**
	 * 
	 * 读取配置文件中的参数
	 * 
	 * 初始化连接池
	 * 
	 */
	private ConnectionManager() {

		// 读取配置文件
		try {
			init();
		} catch (IOException e1) {
			// TODO 日志
			e1.printStackTrace();
		}

		// 取得数据源
		ds = new ComboPooledDataSource();
		// 设置驱动字符串
		try {
			ds.setDriverClass(jdbcDriver);
		} catch (PropertyVetoException e) {
			// TODO 日志
			e.printStackTrace();
		}
		// 设置数据库路径
		ds.setJdbcUrl(dbUrl);
		// 设置数据库用户名
		ds.setUser(dbUsername);
		// 设置数据库密码
		ds.setPassword(dbPassword);

		// 初始化时获取三个连接，取值应在minPoolSize与maxPoolSize之间。
		ds.setInitialPoolSize(initialConnections);

		// 连接池中保留的最大连接数。
		ds.setMaxPoolSize(maxConnections);

		// 连接池中保留的最小连接数。
		ds.setMinPoolSize(minPoolSize);

		// 当连接池中的连接耗尽的时候c3p0一次同时获取的连接数。
		ds.setAcquireIncrement(incrementalConnections);

		// 间隔多少秒检查所有连接池中的空闲连接。单位（ 秒）Default: 0 idleConnectionTestPeriod
		ds.setIdleConnectionTestPeriod(idleConnectionTestPeriod);

		// 最大空闲时间,单位（秒）最大空闲时间内未使用则连接被丢弃。
		// 若为0则永不丢弃。Default: 0 maxIdleTime
		ds.setMaxIdleTime(freeMaxTime);

		// 连接关闭时默认将所有未提交的操作回滚。Default: false autoCommitOnClose
		ds.setAutoCommitOnClose(autoCommitOnClose);

		// c3p0将建一张名为Test的空表，并使用其自带的查询语句进行测试。如果定义了这个参数那么
		// 属性preferredTestQuery将被忽略。你不能在这张Test表上进行任何操作，它将只供c3p0测试使用
//		ds.setAutomaticTestTable(automaticTestTable);

		// 如果设为true那么在取得连接的同时将校验连接的有效性。Default: false testConnectionOnCheckin
		ds.setTestConnectionOnCheckin(testConnectionOnCheckin);

		// 定义在从数据库获取新连接失败后重复尝试的次数。Default: 30 acquireRetryAttempts
		ds.setAcquireRetryAttempts(acquireRetryAttempts);

		// 两次连接中间隔时间，单位毫秒。Default: 1000 acquireRetryDelay
		ds.setAcquireRetryDelay(acquireRetryDelay);

		// 获取连接失败将会引起所有等待连接池来获取连接的线程抛出异常。但是数据源仍有效
		// 保留，并在下次调用getConnection()的时候继续尝试获取连接。如果设为true，那么在尝试
		// 获取连接失败后该数据源将申明已断开并永久关闭。Default: false breakAfterAcquireFailure
		ds.setBreakAfterAcquireFailure(breakAfterAcquireFailure);

		// 当连接池用完时客户端调用getConnection()后等待获取新连接的时间，超时后将抛出
		// SQLException,如设为0则无限期等待。单位毫秒。Default: 0
		ds.setCheckoutTimeout(checkoutTimeout);

		// 用户修改系统配置参数执行前最多等待300秒
		ds.setPropertyCycle(propertyCycle);

		// vip 连接池
		dsVip = new ComboPooledDataSource();
		// 数据库驱动字符串
		try {
			dsVip.setDriverClass(jdbcDriver);
		} catch (PropertyVetoException e) {
			// TODO 日志
			e.printStackTrace();
		}
		// 数据库路径
		dsVip.setJdbcUrl(dbUrl);
		// 数据库用户名
		dsVip.setUser(dbUsername);
		// 数据库密码
		dsVip.setPassword(dbPassword);

		// 初始化时获取三个连接，取值应在minPoolSize与maxPoolSize之间。
		dsVip.setInitialPoolSize(vipConnSize);

		// 设置最小连接数
		dsVip.setMinPoolSize(0);

		// 连接池中保留的最大连接数。
		dsVip.setMaxPoolSize(vipConnSize);

		// 间隔多少秒检查所有连接池中的空闲连接。Default: 0 idleConnectionTestPeriod
		dsVip.setIdleConnectionTestPeriod(idleConnectionTestPeriod);

		// 最大空闲时间,单位（秒）最大空闲时间内未使用则连接被丢弃。
		// 若为0则永不丢弃。Default: 0 maxIdleTime
		dsVip.setMaxIdleTime(freeMaxTime);

		// 连接关闭时默认将所有未提交的操作回滚。Default: false autoCommitOnClose
		dsVip.setAutoCommitOnClose(autoCommitOnClose);

		// c3p0将建一张名为Test的空表，并使用其自带的查询语句进行测试。如果定义了这个参数那么
		// 属性preferredTestQuery将被忽略。你不能在这张Test表上进行任何操作，它将只供c3p0测试使用
//		dsVip.setAutomaticTestTable(automaticTestTable);

		// 如果设为true那么在取得连接的同时将校验连接的有效性。Default: false testConnectionOnCheckin
		dsVip.setTestConnectionOnCheckin(testConnectionOnCheckin);

		// 定义在从数据库获取新连接失败后重复尝试的次数。Default: 30 acquireRetryAttempts
		dsVip.setAcquireRetryAttempts(acquireRetryAttempts);

		// 两次连接中间隔时间，单位毫秒。Default: 1000 acquireRetryDelay
		dsVip.setAcquireRetryDelay(acquireRetryDelay);

		// 获取连接失败将会引起所有等待连接池来获取连接的线程抛出异常。但是数据源仍有效
		// 保留，并在下次调用getConnection()的时候继续尝试获取连接。如果设为true，那么在尝试
		// 获取连接失败后该数据源将申明已断开并永久关闭。Default: false breakAfterAcquireFailure
		dsVip.setBreakAfterAcquireFailure(breakAfterAcquireFailure);

		// 当连接池用完时客户端调用getConnection()后等待获取新连接的时间，超时后将抛出
		// SQLException,如设为0则无限期等待。单位毫秒。Default: 0
		dsVip.setCheckoutTimeout(checkoutTimeout);

		// 用户修改系统配置参数执行前最多等待300秒
		dsVip.setPropertyCycle(propertyCycle);

	}

	/**
	 * 
	 * 初始化数据库连接使用的参数
	 * <ul>
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 * 
	 * @throws IOException
	 */
	private void init() throws IOException {

		// // 读取配置文件
		// InputStream is = getClass().getResourceAsStream("/db.properties");
		// // 实例一个属性文件类
		// Properties dbProps = new Properties();
		// // 加载数据流
		// dbProps.load(is);

		// 实例一个属性文件类
		Properties property = new Properties();

		// 根据路径读取配置文件
//		InputStream is = this.getClass().getClassLoader().getResourceAsStream("db.properties");
		String loc = URLDecoder.decode(ConnectionManager.class.getProtectionDomain().getCodeSource().getLocation().getFile(), "UTF-8");
		loc = loc.substring(0, loc.lastIndexOf("/"));
		FileInputStream f = new FileInputStream(loc + "/db.properties");
		property.load(f);

		// 数据库驱动类
		jdbcDriver = property.getProperty("jdbc.driverClassName");

		// 路径
		dbUrl = property.getProperty("jdbc.url");

		// 用户名
		dbUsername = property.getProperty("jdbc.username");

		// 密码
		dbPassword = property.getProperty("jdbc.password");

		// 最大连接数
		maxConnections = Integer.parseInt(property.getProperty("defaultpoolname.maxconn").trim());

		// 最小连接数
		initialConnections = Integer.parseInt(property.getProperty("initial.connections").trim());

		// 自动增长数
		incrementalConnections = Integer.parseInt(property.getProperty("incremental.connections").trim());

		// vip连接数
		vipConnSize = Integer.parseInt(property.getProperty("vip.conn.size").trim());

		// 连接关闭时默认将所有未提交的操作回滚。Default: false
		autoCommitOnClose = Boolean.parseBoolean(property.getProperty("autocommit.onclose").trim());

		// 数据库获取新连接失败后重复尝试的次数
		acquireRetryAttempts = Integer.parseInt(property.getProperty("acquire.retry.attempts").trim());

		// 两次连接中间隔时间，单位毫秒
		acquireRetryDelay = Integer.parseInt(property.getProperty("acquire.retry.delay").trim());

		// 获取连接失败线程抛出异常时，设定数据源是否依旧有效，false：有效，true：关闭数据源
		breakAfterAcquireFailure = Boolean.parseBoolean(property.getProperty("break.after.acquire.failure").trim());

		// 用来测试连接是否可用的表名，其可以由用户指定，c3p0自动创建。
//		automaticTestTable = property.getProperty("automatic.test.table").trim();

		// 获取获取每一个连接时，是否测试连接是否可用
		testConnectionOnCheckin = Boolean.parseBoolean(property.getProperty("testConnection.oncheckin").trim());

		// 当连接池的连接用完时客户端调用getConnection()后等待获取新连接的时间
		checkoutTimeout = Integer.parseInt(property.getProperty("checkout.timeout").trim());

		// 用户修改系统配置参数执行前最多等待时间 单位：秒
		propertyCycle = Integer.parseInt(property.getProperty("propertycycle").trim());

		// 连接池中保留的最小连接数。
		minPoolSize = Integer.parseInt(property.getProperty("defaultpoolname.minconn").trim());

	}

	/**
	 * 
	 * 返回一个普通连接
	 * 
	 * <ul>
	 * <li>@return Connection
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static synchronized final Connection getNormalConnection() throws Exception {
		try {
			ConnectionManager.callNormalConnCount++;
			// 如果计数器达到1000000，则计数器归零，百万计数器加1
			if (ConnectionManager.callNormalConnCount == 1000000) {
				ConnectionManager.callNormalConnCount = 0;
				ConnectionManager.callNormalConnMillionCount++;
			}
		} catch (Exception e) {
			callNormalConnCount = 0;
			callNormalConnMillionCount = 0;
		}

		// 返回一个数据库连接
		return ds.getConnection();
	}

	/**
	 * 
	 * 返回一个vip连接
	 * 
	 * <ul>
	 * <li>@return Connection
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static synchronized final Connection getVipConnection() throws Exception {
		try {
			ConnectionManager.callVipConnCount++;
			// 如果计数器达到1000000，则计数器归零，百万计数器加1
			if (ConnectionManager.callVipConnCount == 1000000) {
				ConnectionManager.callVipConnCount = 0;
				ConnectionManager.callVipConnMillionCount++;
			}
		} catch (Exception e) {
			callVipConnCount = 0;
			callVipConnMillionCount = 0;
		}

		// 返回一个数据库连接
		return dsVip.getConnection();
	}

	/**
	 * 
	 * 返回初始的普通权限连接数
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getInitialPoolSize() throws Exception {

		// 取得初始连接数
		int connSize = ds.getInitialPoolSize();
		// 返回连接数
		return connSize;
	}

	/**
	 * 
	 * 返回活动状态的普通权限连接
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getNumBusyConnections() throws Exception {

		// 取得普通连接池中 活动状态的连接数
		int connSize = ds.getNumBusyConnections();
		// 返回活动状态连接数
		return connSize;
	}

	/**
	 * 
	 * 返回最大普通权限连接数
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getMaxPoolSize() throws Exception {

		// 取得普通连接池中 最大连接数
		int connSize = ds.getMaxPoolSize();
		// 返回最大连接数
		return connSize;
	}

	/**
	 * 
	 * 返回最小普通权限连接数
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getMinPoolSize() throws Exception {

		// 取得普通连接池最小连接数
		int connSize = ds.getMinPoolSize();
		// 返回最小连接数
		return connSize;
	}

	/**
	 * 
	 * 返回初始的Vip权限连接数
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getInitialVipPoolSize() throws Exception {

		// 取得初始的最小vip连接数
		int connSize = ds.getInitialPoolSize();
		// 返回最小连接数
		return connSize;
	}

	/**
	 * 
	 * 返回活动状态的Vip权限连接
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getNumBusyVipConnections() throws Exception {

		// 取得活动状态的vip连接数
		int connSize = dsVip.getNumBusyConnections();
		// 返回活动状态的vip连接数
		return connSize;
	}

	/**
	 * 
	 * 返回最大Vip权限连接数
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getMaxVipPoolSize() throws Exception {

		// 取得最大vip连接数
		int connSize = dsVip.getMaxPoolSize();
		// 返回最大vip连接数
		return connSize;
	}

	/**
	 * 
	 * 返回最小Vip权限连接数
	 * 
	 * <ul>
	 * <li>@return int connSize
	 * <li>@throws Exception 抛出Exception
	 * </ul>
	 */
	public static int getMinVipPoolSize() throws Exception {

		// 取得最小vip连接数
		int connSize = dsVip.getMinPoolSize();
		// 返回最小vip连接数
		return connSize;
	}

	/**
	 * 
	 * 析构函数
	 * 
	 * <ul>
	 * <li>@throws Throwable 抛出Throwable
	 * </ul>
	 */
	protected void finalize() throws Throwable {

		// 关闭数据源
		DataSources.destroy(ds);
		// 清理连接池所占用资源
		super.finalize();
	}

	/**
	 * 获取普通连接调用次数
	 * 
	 * @return int
	 */
	public static int getCallNormalConnCount() {
		return ConnectionManager.callNormalConnCount;
	}

	/**
	 * 获取普通连接调用百万次数
	 * 
	 * @return int
	 */
	public static int getCallNormalConnMillionCount() {
		return ConnectionManager.callNormalConnMillionCount;
	}

	/**
	 * 获取VIP连接调用次数
	 * 
	 * @return int
	 */
	public static int getCallVipConnCount() {
		return ConnectionManager.callVipConnCount;
	}

	/**
	 * 获取VIP连接调用百万次数
	 * 
	 * @return int
	 */
	public static int getCallVipConnMillionCount() {
		return ConnectionManager.callVipConnMillionCount;
	}
	
	

}
