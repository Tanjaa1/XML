import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'bootstrap/dist/css/bootstrap.min.css'
//import axios from 'axios'
//import VueAxios from 'vue-axios'
// // const express = require("express")
// // const app = express()
// const cors = require("cors")
// App.use(
//     cors({
//         origin: "localhost:8080",
//         credentials: true
//     })
// )

//VueAxios.use(VueAxios, axios)

createApp(App).use(router).mount('#app')
