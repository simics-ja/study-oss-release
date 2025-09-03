const StringFormatter = require('./index');

console.log('Testing StringFormatter...\n');

console.log('capitalize("hello world"):', StringFormatter.capitalize('hello world'));
console.log('camelCase("hello world"):', StringFormatter.camelCase('hello world'));
console.log('kebabCase("helloWorld"):', StringFormatter.kebabCase('helloWorld'));
console.log('snakeCase("helloWorld"):', StringFormatter.snakeCase('helloWorld'));
console.log('reverse("hello"):', StringFormatter.reverse('hello'));
console.log('truncate("This is a very long string that needs to be truncated", 20):', 
  StringFormatter.truncate('This is a very long string that needs to be truncated', 20));

console.log('\nAll tests completed!');