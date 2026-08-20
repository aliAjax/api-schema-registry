fetch('/healthz').then(function (r) { return r.json(); }).then(function (data) {
  document.querySelector('.status').textContent = data.status || 'healthy';
}).catch(function () {
  document.querySelector('.status').textContent = 'Service unavailable';
});
