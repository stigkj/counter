$(function() {
  $('.counter').click(function() {
    $.get('/counter?type=inc', function(data) {
      $('.counter').text(data);
    });
  });
});
