/** @type {import('tailwindcss').Config} */
module.exports = {
  // NOTE: Update this to include the paths to all files that contain Nativewind classes.
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  presets: [require("nativewind/preset")],
  theme: {
    extend: {
      colors: {
        dark: {
          bg: '#171717',
          surface: '#262626',
          card: '#404040',
          border: '#525252',
          text: '#f5f5f5',
          muted: '#a3a3a3',
        }
      },
    },
  },
  plugins: [],
}