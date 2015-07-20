<?php

//cvs demo

include_once 'PHPExcel.php';

$file = "./csv.csv";

$excelReader = PHPExcel_IOFactory::createReader("CSV");
$excelReader->setReadDataOnly(true);

$phpexcel           = $excelReader->load($file);
$activeSheet        = $phpexcel->getActiveSheet();
$highestRow         = $activeSheet->getHighestRow(); //总行数
$highestColumn      = $activeSheet->getHighestColumn(); //最后列数所对应的字母，例如第2行就是B
$highestColumnIndex = \PHPExcel_Cell::columnIndexFromString($highestColumn); //总列数

$excelData = array();
for ($row = 0; $row <= $highestRow; $row++) {
    //注意highestColumnIndex的列数索引从0开始
    for ($col = 0; $col < $highestColumnIndex; $col++) {
        $excelData[$row][] = (string) $activeSheet->getCellByColumnAndRow($col, $row)->getValue();
    }
}

var_dump($excelData);

