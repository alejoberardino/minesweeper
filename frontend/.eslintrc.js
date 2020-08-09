module.exports = {
  root: true,
  env: {
    node: true
  },
  plugins: ['es-beautifier'],
  extends: ['plugin:vue/essential', 'plugin:es-beautifier/standard', '@vue/typescript', 'plugin:prettier/recommended'],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'linebreak-style': 'off',
    'prettier/prettier': ['error', {
      'singleQuote': true,
      'trailingComma': 'es5',
      'semi': false,
      'printWidth': 120,
      'tabWidth': 2
    }]
  },
  parserOptions: {
    parser: '@typescript-eslint/parser'
  }
 };