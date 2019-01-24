$(document).ready(function() {
    $('.btn-logout').click(function(e) {
      Cookies.remove('auth-session');
    });
});


$("#editProfile").submit(function( event ) {
  event.preventDefault();
  console.log(event)
  $.post( "/profile", { city: "blahd"})
  .done(function( data ) {
    var $inputs = $('#editProfile :input');
    console.log($inputs);

    // not sure if you wanted this, but I thought I'd add it.
    // get an associative array of just the values.
    var values = {};
    $inputs.each(function() {
      values[this.name] = $(this).val();
    });
    alert( "Data Loaded: " + data );
  });
});
