package org.opencloudb.parser;

import java.sql.SQLSyntaxErrorException;

import junit.framework.Assert;

import org.junit.Test;
import org.opencloudb.mpp.InsertParsInf;
import org.opencloudb.mpp.InsertSQLAnalyser;
import org.opencloudb.parser.SQLParserDelegate;

import com.akiban.sql.parser.QueryTreeNode;

public class TestInsertSQLAnalyser {

	@Test
	public void testInsertSQL() throws SQLSyntaxErrorException {
		String sql = null;
		QueryTreeNode ast = null;
		InsertParsInf parsInf = null;
		sql = "insert into table1  select * FROM table2 WHERE id not in ( select id from  table1) ";
		ast = SQLParserDelegate.parse(sql, SQLParserDelegate.DEFAULT_CHARSET);
		parsInf = InsertSQLAnalyser.analyse( ast);
		Assert.assertEquals("table1".toUpperCase(), parsInf.tableName);
		Assert.assertEquals(0, parsInf.columnPairMap.size());
		Assert.assertNotNull(parsInf.fromQryNode);

		sql = "insert into table1(column1,column2,column3,colum4,column5,column6,column7)values('aaa',5,'1999-2-2',true,\"test\",111,55.66) ";
		ast = SQLParserDelegate.parse(sql, SQLParserDelegate.DEFAULT_CHARSET);
		parsInf = InsertSQLAnalyser.analyse( ast);
		Assert.assertEquals("table1".toUpperCase(), parsInf.tableName);
		Assert.assertEquals(7, parsInf.columnPairMap.size());
		Assert.assertNull(parsInf.fromQryNode);

		sql = "inSErt into offer_detail (`offer_id`, gmt) values (123,now())";
		ast = SQLParserDelegate.parse(sql, SQLParserDelegate.DEFAULT_CHARSET);
		parsInf = InsertSQLAnalyser.analyse( ast);
		Assert.assertEquals("offer_detail".toUpperCase(), parsInf.tableName);
		Assert.assertEquals(2, parsInf.columnPairMap.size());
		Assert.assertNull(parsInf.fromQryNode);
		
		sql = "insert into offer_detail (offer_id, gmt) values (0, now()), (1, now()), (2, now())";
		ast = SQLParserDelegate.parse(sql, SQLParserDelegate.DEFAULT_CHARSET);
		try{
		parsInf = InsertSQLAnalyser.analyse( ast);
		}catch(Exception e)
		{
			Assert.assertEquals("insert multi rows not supported", e.getMessage());
		}

//	sql = "delete from offer.*,wp_image.* using offer a,wp_image b where a.member_id=b.member_id and a.member_id='abc' ";
		

	}
}
