import { createRouter, createWebHashHistory } from 'vue-router'
import Registration from '../views/Registration.vue'
import Proba from '../views/Proba.vue'
import Profil from '../views/Profil.vue'
import EditProfile from '../views/EditProfile.vue'
import Login from '../views/Login.vue'
import VerificationRequest from '../views/VerificationRequest.vue'
import Requests from '../views/Requests.vue'
import LikesAndDislikes from '../views/LikesAndDislikes.vue'
import Collections from '../views/Collections.vue'
import AllStories from '../views/AllStories.vue'
import ProfilUser from '../views/ProfilUser.vue'
import PostsByLocation from '../views/PostsByLocation.vue'
import PostsByHashtag from '../views/PostsByHashtag.vue'
import PostsByUserPublic from '../views/PostsByUserPublic.vue'
import PostsByLocationPublic from '../views/PostsByLocationPublic.vue'
import PostsByHashtagPublic from '../views/PostsByHashtagPublic.vue'
import RequestList from '../views/RequestList.vue'
import PostsLiked from '../views/PostsLiked.vue'

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
    path: '/VerificationRequest',
    name: 'VerificationRequest',
    component: VerificationRequest
  },
  {
    path: '/Requests',
    name: 'Requestins',
    component: Requests
  },
  {
    path: '/LD',
    name: 'LikesAndDislikes',
    component: LikesAndDislikes
  },
  {
    path: '/Collections',
    name: 'Collections',
    component: Collections
  },
  {
    path: '/AllStories',
    name: 'AllStories',
    component: AllStories
  },
  {
    path: '/ProfilUser',
    name: 'ProfilUser',
    component: ProfilUser
  },
  {
    path: '/PostsByLocation',
    name: 'PostsByLocation',
    component: PostsByLocation
  },
  {
    path: '/PostsByHashtag',
    name: 'PostsByHashtag',
    component: PostsByHashtag
  },
  {
    path: '/PostsByUserPublic',
    name: 'PostsByUserPublic',
    component: PostsByUserPublic
  },
  {
    path: '/PostsByLocationPublic',
    name: 'PostsByLocationPublic',
    component: PostsByLocationPublic
  },
  {
    path: '/PostsByHashtagPublic',
    name: 'PostsByHashtagPublic',
    component: PostsByHashtagPublic
  },
  {
    path: '/RequestList',
    name: 'RequestList',
    component: RequestList
  },
  {
    path: '/PostsLiked',
    name: 'PostsLiked',
    component: PostsLiked
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
