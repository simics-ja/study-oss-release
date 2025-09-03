# @study-oss/string-formatter

A simple string formatting utility library for learning OSS release processes.

## Installation

```bash
npm install @study-oss/string-formatter
```

## Usage

```javascript
const StringFormatter = require('@study-oss/string-formatter');

// Capitalize first letter
StringFormatter.capitalize('hello world'); // 'Hello world'

// Convert to camelCase
StringFormatter.camelCase('hello world'); // 'helloWorld'

// Convert to kebab-case
StringFormatter.kebabCase('helloWorld'); // 'hello-world'

// Convert to snake_case
StringFormatter.snakeCase('helloWorld'); // 'hello_world'

// Reverse string
StringFormatter.reverse('hello'); // 'olleh'

// Truncate string
StringFormatter.truncate('This is a very long string', 20); // 'This is a very lo...'
```

## API

### `capitalize(str)`
Capitalizes the first letter of a string.

### `camelCase(str)`
Converts a string to camelCase.

### `kebabCase(str)`
Converts a string to kebab-case.

### `snakeCase(str)`
Converts a string to snake_case.

### `reverse(str)`
Reverses a string.

### `truncate(str, length = 30, suffix = '...')`
Truncates a string to the specified length.

## Testing

```bash
npm test
```

## License

MIT