import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import AnalyzerView from "../views/AnalyzerView.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "analyzer",
    component: AnalyzerView,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
