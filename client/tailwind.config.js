/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'body': '#181818',
        'secondary-body': '#272727',
        'selected-text': '#545FE8',
        'theme': '#A0A7FA',
        'nav': '#BABABA',
        'primary': '#545FE8',
        'secondary': '#23254D',
        'tertiary': '#2F337D',
        'badge': '#3F3F51',
        'input-border': '#565666',
        'input': '#2A2A35'
      },
      fontFamily: {
        'roboto': ["'Roboto'", 'sans-serif']
      },
    },
  },
}