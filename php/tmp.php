<html>
    <head>
        <script src="http://cdn.bootcss.com/jquery/2.2.0/jquery.min.js"></script>
    </head>

    <body>
        <script>
            $.ajax({
                url: "http://192.168.1.119:10001/api/course/list/",
                type: 'POST',
                data: JSON.stringify({"start_time": 1456761600}),
                cache: false,
                dataType: "json",
                success: function(result) {
                    alert(result.status);
                    alert(result.message);
                },
                error: function(err) {
                    alert(err)
                }
            });
        </script>
    </body>
</html>