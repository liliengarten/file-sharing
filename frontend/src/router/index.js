import { createRouter, createWebHistory } from 'vue-router'
import HomeView from "@/views/HomeView.vue";
import LoginView from "@/views/LoginView.vue";
import RegisterView from "@/views/RegisterView.vue";
import ProfileView from "@/views/ProfileView.vue";
import PinView from "@/views/PinView.vue";
import CreatePinView from "@/views/CreatePinView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      component: HomeView
    },
    {
      path: "/login",
      component: LoginView
    },
    {
      path: "/register",
      component: RegisterView
    },
    {
      path: "/profile",
      component: ProfileView
    },
    {
      path: "/pin/:id",
      component: PinView
    },
    {
      path: "/create_pin",
      component: CreatePinView
    }
  ]
})

export default router
