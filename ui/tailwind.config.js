/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './public/index.html',
    './src/**/*.{ts,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        'background': 'var(--background)',
        'foreground': 'var(--foreground)',
        'background-selected': 'var(--background-selected)',
        'background-inverted': 'var(--background-inverted)',
        'foreground-inverted': 'var(--foreground-inverted)',
      }
    },
  },
  plugins: [],
}
