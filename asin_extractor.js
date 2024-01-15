var htmlContent = document.documentElement.outerHTML;

var regex = /(?<=<div data-asin=")[^"]*/g;

var matches = htmlContent.match(regex);

var blob = new Blob([matches.join('\n')], {type: 'text/plain'});

var link = document.createElement('a');
link.href = URL.createObjectURL(blob);
link.download = 'asin_codes.txt';

document.body.appendChild(link);

link.click();

document.body.removeChild(link);
