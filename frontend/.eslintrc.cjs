module.exports = {
  extends: [
    'next/core-web-vitals',
  ],
  root: true,
  ignorePatterns: ['.eslintrc.cjs'],
  rules: {
    '@typescript-eslint/interface-name-prefix': 'off',
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    '@typescript-eslint/no-explicit-any': 'off',
    'import/no-anonymous-default-export': 'off',
    // 'max-len': [1, { code: 200 }],
  },
};