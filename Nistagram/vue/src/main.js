import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'bootstrap/dist/css/bootstrap.min.css'
// // const express = require("express")
// // const app = express()
// const cors = require("cors")
// App.use(
//     cors({
//         origin: "localhost:8080",
//         credentials: true
//     })
// )

createApp(App).use(router).mount('#app')
