$(function() {
    $.post('/counter', function(data) {
  $('.counter').click(function() {
      $('.counter').text(data);
    });
  });
});
