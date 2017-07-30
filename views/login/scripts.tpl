<script type="text/javascript">
   $(function() {

    $('#login-form-link').click(function(e) {
		$("#login-form").delay(100).fadeIn(100);
 		$("#register-form").fadeOut(100);
		$('#register-form-link').removeClass('active');
		$(this).addClass('active');
		e.preventDefault();
	});
	$('#register-form-link').click(function(e) {
		$("#register-form").delay(100).fadeIn(100);
 		$("#login-form").fadeOut(100);
		$('#login-form-link').removeClass('active');
		$(this).addClass('active');
		e.preventDefault();
	});

	$('#login-submit').click(function(e){
		var username = $('#username').val();
		var password = $('#password').val();
		$.post('/signin',{username:username,password,password},function(result){

			if(result.Success)
			{
				alert('登录成功！')
				//windown.href="/"
			}
			else{
				alert(result.Message);
			}

		})
	})

});
</script>