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


var files;

$("#imageForm").submit(function(e){
  e.stopPropagation(); // Stop stuff happening
  e.preventDefault(); // Totally stop stuff happening

  var data = new FormData();
  $.each(files, function(key, value) {
    data.append("file", value);
  }); 

  $.ajax({
    url: '/profile/image',
    type: 'POST',
    data: data,
    cache: false,
    dataType: 'json',
    processData: false, // Don't process the files
    contentType: false, // Set content type to false as jQuery will tell the server its a query string request
    success: function(data, textStatus, jqXHR)
    {
      if(typeof data.error === 'undefined')
        {
          // Success so call function to process the form
          submitForm(event, data);
        }
        else
          {
            // Handle errors here
            console.log('ERRORS: ' + data.error);
          }
    },
    error: function(jqXHR, textStatus, errorThrown)
    {
      // Handle errors here
      console.log('ERRORS: ' + textStatus);
      // STOP LOADING SPINNER
    }
  });
});

$("#imageUpload").change(function(e){
  e.stopPropagation(); // Stop stuff happening
  e.preventDefault(); // Totally stop stuff happening
  files = event.target.files;
  console.log(e.target.files);

});

$('#imageUpload').on('click touchstart' , function(){
  $(this).val('');
});
