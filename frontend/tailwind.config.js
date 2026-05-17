// tailwind.config.js
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      colors: {
        brand: {
          50: "#fff7ed",
          100: "#ffedd5",
          200: "#fed7aa",
          300: "#fdba74",
          400: "#fb923c",
          500: "#f97316",
          600: "#ea580c",
          700: "#c2410c",
          800: "#9a3412",
          900: "#7c2d12",
        },
      },
      backgroundImage: {
        "brand-gradient": "linear-gradient(160deg, #9a3412 0%, #c2410c 45%, #ea580c 100%)",
        "brand-gradient-soft": "linear-gradient(135deg, #fff7ed 0%, #ffedd5 100%)",
      },
      boxShadow: {
        brand: "0 12px 30px rgba(249, 115, 22, 0.18)",
      },
    },
  },
  plugins: [],
};