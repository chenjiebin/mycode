<?php
$sso = "e589hR6VnO8K1CNQZ4PSP/LWGBhRKE5VckawQwl1TdE8d4Q5E7tW";
setcookie("sso", $sso);
?>
login success...
<script type="text/javascript">
    function jumpTo() {
        location.href = "http://localhost/~chenjiebin/mycode/php/cookie/a.com/";
    }
</script>
<iframe src="http://localhost/~chenjiebin/mycode/php/cookie/b.com/synclogin.php?sso=<?php echo $sso; ?>"></iframe>

