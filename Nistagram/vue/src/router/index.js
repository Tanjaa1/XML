import { createRouter, createWebHashHistory } from 'vue-router'
import Registration from '../views/Registration.vue'
import Proba from '../views/Proba.vue'
import Profil from '../views/Profil.vue'
import EditProfile from '../views/EditProfile.vue'
import Login from '../views/Login.vue'
const routes = [
  {
    path: '/',
    name: 'Registartion',
    component: Registration
  },
  {
    path: '/Proba',
    name: 'Proba',
    component: Proba
  },
  {
    path: '/Profil',
    name: 'Profil',
    component: Profil
  },
  {
    path: '/EditProfile',
    name: 'EditProfile',
    component: EditProfile
  },
  {
    path: '/Login',
    name: 'Login',
    component: Login
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
