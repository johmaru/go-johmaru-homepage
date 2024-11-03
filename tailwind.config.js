/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js}"],
  theme: {
    screens: {
      'sm': '640px',
      'md': '768px',
      'lg': '1024px',
    },
    colors: {
      'red': '#4c0519',
      'green': '#052e16',
      'eme': '#4ade80'
    }
  },
  plugins: [],
}

