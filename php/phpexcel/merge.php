<?php

/**
 * 合并单元格操作，目前只解决纵向合并的问题
 */
/** Error reporting */
error_reporting(E_ALL);
ini_set('display_errors', TRUE);
ini_set('display_startup_errors', TRUE);
date_default_timezone_set('Europe/London');

if (PHP_SAPI == 'cli')
    die('This example should only be run from a Web Browser');

/** Include PHPExcel */
require_once dirname(__FILE__) . '/PHPExcel.php';

$content = array(
    array("id", "name", "fields", "age"),
    array("1", "cjb", array("1", "2"), "20"),
    array("1", "cjb", "", "20"),
);

$filename = "temp.xls";

// Create new PHPExcel object
$objPHPExcel = new PHPExcel();

// Set document properties
$objPHPExcel->getProperties()->setCreator("Maarten Balliauw")
        ->setLastModifiedBy("Maarten Balliauw")
        ->setTitle("Office 2007 XLSX Test Document")
        ->setSubject("Office 2007 XLSX Test Document")
        ->setDescription("Test document for Office 2007 XLSX, generated using PHP classes.")
        ->setKeywords("office 2007 openxml php")
        ->setCategory("Test result file");

$objPHPExcel->setActiveSheetIndex(0);

//var_dump($content);
//exit();
$objPHPExcel->getActiveSheet()->setCellValue("D3", 20);

$i = 0;
foreach ($content as $row) {
//    echo "i: " . $i;
//    echo "<br />";
    $j = 0;
    foreach ($row as $v) {
        if (is_array($v)) {
            $tmpI = $i;
            foreach ($v as $t) {
                $pCoordinate = PHPExcel_Cell::stringFromColumnIndex($j) . '' . ($tmpI + 1);
                $pValue      = $t;
                $objPHPExcel->getActiveSheet()->setCellValue($pCoordinate, $pValue);

                $tmpI++;
            }
        } else {
            $pCoordinate = PHPExcel_Cell::stringFromColumnIndex($j) . '' . ($i + 1);
            $pValue      = $v;
            $objPHPExcel->getActiveSheet()->setCellValue($pCoordinate, $pValue);
        }

        $j++;
    }

    //计算该行占用几行
    $lines = 1;
    foreach ($row as $v) {
        if (is_array($v)) {
            $c = count($v);
            if ($c > $lines) {
                $lines = $c;
            }
        }
    }
    //合并单元格
    if ($lines > 1) {
        $j = 0;
        foreach ($row as $v) {
            if (!is_array($v)) {
                $pCoordinateStart = PHPExcel_Cell::stringFromColumnIndex($j) . '' . ($i + 1);
                $pCoordinateEnd   = PHPExcel_Cell::stringFromColumnIndex($j) . '' . ($i + $lines);
                $objPHPExcel->getActiveSheet()->mergeCells("$pCoordinateStart:$pCoordinateEnd");
            }
            $j++;
        }
    }

    $i = $i + $lines;
}
//exit();
// Redirect output to a client’s web browser (Excel2007)
header('Content-Type: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet');
header('Content-Disposition: attachment;filename="' . $filename . '"');
header('Cache-Control: max-age=0');
// If you're serving to IE 9, then the following may be needed
header('Cache-Control: max-age=1');

// If you're serving to IE over SSL, then the following may be needed
header('Expires: Mon, 26 Jul 1997 05:00:00 GMT'); // Date in the past
header('Last-Modified: ' . gmdate('D, d M Y H:i:s') . ' GMT'); // always modified
header('Cache-Control: cache, must-revalidate'); // HTTP/1.1
header('Pragma: public'); // HTTP/1.0

$objWriter = PHPExcel_IOFactory::createWriter($objPHPExcel, 'Excel5');
$objWriter->save('php://output');
exit;
