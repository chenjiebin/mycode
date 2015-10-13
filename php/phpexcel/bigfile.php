<?php

/*
  phpexcel读取大文件demo
  Author:陈杰斌
  Author URI: http://www.01happy.com
 */

echo memory_get_usage();

include_once __DIR__ . '/PHPExcel.php';

/**
 * 读取excel过滤器
 */
class PHPExcelReadFilter implements PHPExcel_Reader_IReadFilter {

    public $startRow = 1;
    public $endRow;

    public function readCell($column, $row, $worksheetName = '') {
        //如果endRow没有设置表示读取全部
        if (!$this->endRow) {
            return true;
        }
        //只读取指定的行
        if ($row >= $this->startRow && $row <= $this->endRow) {
            return true;
        }

        return false;
    }

}

$startRow  = 1;
$endRow    = null;
$excelFile = __DIR__ . '/data/payment.xlsx';

$result = readFromExcel($excelFile, null, $startRow, $endRow);

echo "<br />";
echo memory_get_usage();

/**
 * 读取excel转换成数组
 * 
 * @param string $excelFile 文件路径
 * @param string $excelType excel后缀格式
 * @param int $startRow 开始读取的行数
 * @param int $endRow 结束读取的行数
 * @retunr array
 */
function readFromExcel($excelFile, $excelType = null, $startRow = 1, $endRow = null) {
    include_once __DIR__ . '/PHPExcel.php';

    $excelReader = \PHPExcel_IOFactory::createReader("Excel2007");
    $excelReader->setReadDataOnly(true);

    //如果有指定行数，则设置过滤器
    if ($startRow && $endRow) {
        $perf           = new PHPExcelReadFilter();
        $perf->startRow = $startRow;
        $perf->endRow   = $endRow;
        $excelReader->setReadFilter($perf);
    }

    $phpexcel    = $excelReader->load($excelFile);
    $activeSheet = $phpexcel->getActiveSheet();
    if (!$endRow) {
        $endRow = $activeSheet->getHighestRow(); //总行数
    }

    $highestColumn      = $activeSheet->getHighestColumn(); //最后列数所对应的字母，例如第2行就是B
    $highestColumnIndex = \PHPExcel_Cell::columnIndexFromString($highestColumn); //总列数

    $data = array();
    for ($row = $startRow; $row <= $endRow; $row++) {
        for ($col = 0; $col < $highestColumnIndex; $col++) {
            $data[$row][] = (string) $activeSheet->getCellByColumnAndRow($col, $row)->getValue();
        }
    }
    return $data;
}
