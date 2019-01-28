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

