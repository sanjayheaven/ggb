/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/pages/*.tsx"],
  corePlugins: { preflight: false },
  theme: {
    extend: {
      colors: {
        theme: "#6AD7E5",
      },
    },
  },
  plugins: [],
};
