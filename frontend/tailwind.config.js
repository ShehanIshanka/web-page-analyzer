/** @type {import('tailwindcss').Config} */
import * as primeui from "tailwindcss-primeui";

export default {
  content: ["./src/**/*.{vue,html,js}"],
  theme: {
    extend: {},
  },
  plugins: [primeui],
};
