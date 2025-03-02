import { createApp } from "vue";
import App from "./App.vue";
import PrimeVue from "primevue/config";
import router from "./router";
import "primeicons/primeicons.css";
import Aura from "@primevue/themes/aura";
import { definePreset } from "@primevue/themes";
import Toast from "primevue/toast";
import ToastService from "primevue/toastservice";

const Noir = definePreset(Aura, {
  semantic: {
    primary: {
      50: "{zinc.50}",
      100: "{zinc.100}",
      200: "{zinc.200}",
      300: "{zinc.300}",
      400: "{zinc.400}",
      500: "{zinc.500}",
      600: "{zinc.600}",
      700: "{zinc.700}",
      800: "{zinc.800}",
      900: "{zinc.900}",
      950: "{zinc.950}",
    },
    colorScheme: {
      light: {
        primary: {
          color: "{zinc.950}",
          inverseColor: "#ffffff",
          hoverColor: "{zinc.900}",
          activeColor: "{zinc.800}",
        },
        highlight: {
          background: "{zinc.950}",
          focusBackground: "{zinc.700}",
          color: "#ffffff",
          focusColor: "#ffffff",
        },
      },
      dark: {
        primary: {
          color: "{zinc.50}",
          inverseColor: "{zinc.950}",
          hoverColor: "{zinc.100}",
          activeColor: "{zinc.200}",
        },
        highlight: {
          background: "rgba(250, 250, 250, .16)",
          focusBackground: "rgba(250, 250, 250, .24)",
          color: "rgba(255,255,255,.87)",
          focusColor: "rgba(255,255,255,.87)",
        },
      },
    },
  },
});
const app = createApp(App);

app.use(router);
app.use(PrimeVue, {
  // Default theme configuration
  ripple: true,
  theme: {
    preset: Noir,
    options: {
      prefix: "p",
      darkModeSelector: ".dark-mode",
      cssLayer: false, // Todo see https://v4.primevue.org/theming/styled/#csslayer
    },
  },
});
app.use(ToastService);
app.component("Toast", Toast);
app.mount("#app");
