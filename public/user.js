$(document).ready(function() {
    $('.btn-logout').click(function(e) {
      Cookies.remove('auth-session');
    });
});

$("#editProfile").submit(function( event ) {
  event.preventDefault();
  var $inputs = $('#editProfile :input');
  var values = {};
  $inputs.each(function() {
    values[this.name] = $(this).val();
    if (values[this.name] == "true") {
      values[this.name] = true;
    }  else if (values[this.name] == "false") {
      values[this.name] = false
    }
  });

  $.ajax({
    url: 'http://localhost:3000/profile',
    type: 'POST',
    data: JSON.stringify(values),
    contentType: 'application/json; charset=utf-8',
    dataType: 'json',
    async: false,
    success: function( data ) {
      console.log($inputs);
      alert( "Data Loaded: " + data );
    }
  });
})

