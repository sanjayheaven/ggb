/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/pages/*.tsx"],
  corePlugins: { preflight: false },
  theme: {
    extend: {
      colors: {
        theme: "#0890f7",
      },
    },
  },
  plugins: [],
};
