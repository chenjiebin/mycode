package org.opencloudb.mysql.nio.handler;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.locks.ReentrantLock;

import org.apache.log4j.Logger;
import org.opencloudb.MycatConfig;
import org.opencloudb.MycatServer;
import org.opencloudb.mysql.MySQLDataNode;
import org.opencloudb.mysql.nio.MySQLConnection;
import org.opencloudb.net.mysql.ErrorPacket;
import org.opencloudb.net.mysql.FieldPacket;
import org.opencloudb.net.mysql.RowDataPacket;
import org.opencloudb.server.NonBlockingSession;

/**
 * by wuzh fetch store node of child table ,use sql as following select id from
 * company where id=(select company_id from customer where id=3); the one which
 * return data (id) is the datanode to store child table's records
 * 
 * @author wuzhih
 * 
 */
public class FetchStoreNodeOfChildTableHandler implements ResponseHandler {
	private static final Logger LOGGER = Logger
			.getLogger(NonBlockingSession.class);
	private String sql;
	private volatile String result;
	private volatile String dataNode;
	protected final ReentrantLock lock = new ReentrantLock();

	public String execute(String sql, ArrayList<String> dataNodes) {
		this.sql = sql;
		long startTime = System.currentTimeMillis();
		long endTime = startTime + 5 * 60 * 1000L;
		MycatConfig conf = MycatServer.getInstance().getConfig();
		for (String dn : dataNodes) {
			if (dataNode != null) {
				return dataNode;
			}
			MySQLDataNode mysqlDN = conf.getDataNodes().get(dn);
			try {
				if (LOGGER.isDebugEnabled()) {
					LOGGER.debug("execute in datanode " + dn);
				}
				mysqlDN.getConnection(this, dn);
			} catch (Exception e) {
				LOGGER.warn("get connection err " + e);
			}
			try {
				Thread.sleep(200);
			} catch (InterruptedException e) {

			}
		}
		while (dataNode == null && System.currentTimeMillis() < endTime) {
			try {
				Thread.sleep(50);
			} catch (InterruptedException e) {
				break;
			}
			if (dataNode != null) {
				return dataNode;
			}
		}
		return dataNode;

	}

	@Override
	public void connectionAcquired(MySQLConnection conn) {

		conn.setResponseHandler(this);
		try {
			conn.query(sql);
		} catch (Exception e) {
			executeException(conn, e);
		}
	}

	@Override
	public void connectionError(Throwable e, MySQLConnection conn) {
		LOGGER.warn("connectionError " + e);

	}

	@Override
	public void errorResponse(byte[] data, MySQLConnection conn) {
		ErrorPacket err = new ErrorPacket();
		err.read(data);
		LOGGER.warn("errorResponse " + err.errno + " "
				+ new String(err.message));

	}

	@Override
	public void okResponse(byte[] ok, MySQLConnection conn) {
		LOGGER.info("okResponse ");
		conn.setRunning(false);
		conn.release();

	}

	@Override
	public void fieldEofResponse(byte[] header, List<byte[]> fields,
			byte[] eof, MySQLConnection conn) {
		LOGGER.info("fieldEofResponse ,fields " + fields.size());
		for (byte[] data : fields) {
			FieldPacket fieldPkg = new FieldPacket();
			fieldPkg.read(data);
			String fieldName = new String(fieldPkg.name);
			System.out.println(fieldName + " type " + fieldPkg.type);
		}

	}

	@Override
	public void rowResponse(byte[] row, MySQLConnection conn) {
		if (result == null) {

			RowDataPacket rowDataPkg = new RowDataPacket(1);
			rowDataPkg.read(row);
			byte[] columnData = rowDataPkg.fieldValues.get(0);
			String columnVal = new String(columnData);
			result = columnVal;
			dataNode = (String) conn.getAttachment();
		} else {
			LOGGER.warn("find multi data nodes for child table store, sql is:  "
					+ sql);
		}

	}

	@Override
	public void rowEofResponse(byte[] eof, MySQLConnection conn) {
		LOGGER.info("rowEofResponse   ");
		conn.setRunning(false);
		conn.release();
	}

	private void executeException(MySQLConnection c, Throwable e) {
		LOGGER.warn("executeException   " + e);
		c.setRunning(false);
		c.close();

	}

	@Override
	public void writeQueueAvailable() {
		// TODO Auto-generated method stub
		
	}

}
