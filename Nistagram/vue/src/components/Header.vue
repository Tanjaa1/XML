<template>
    <nav>
        <Search/>
       <div class="menu-item"></div>
       <div class="menu-item"></div>
       <div class="menu-item"></div>
       <Dropdown title="More" :items="services"/>
    </nav>
</template>

<script>
import Dropdown  from "../components/Dropdown"
import Search  from "../components/Search.vue"

export default {
  components: { Dropdown,Search },
    name:'navbar',
    data(){
        return{
            services:[
                {
                    title:'Profile',
                    link:'Profil'
                },
                {
                    title:'AllStories',
                    link:'AllStories'
                },
                {
                    title:'Collections',
                    link:'Collections'
                },
                {
                    title:'EditProfile',
                    link:'EditProfile'
                },
                {
                    title:'Log Out',
                    link:'Login'
                }
            ],
            searches:[
                {
                    title:'Account',
                    link:''
                },
                {
                    title:'Tags',
                    link:''
                },
                {
                    title:'Location',
                    link:''
                }
            ]

        }
    }
}
</script>

<style>
nav{
    display: flex;
    align-items: center;
    justify-content: center;
}

nav .menu-item{
    color: #FFF;
    padding: 10px 20px;
    position: relative;
    text-align: center;
    border-bottom: 3px solid transparent;
    display: flex;
    transition: 0.4s;
}

nav .menu-item.active,
nav .menu-item:hover{
    background-color: #4c4c5a;
    border-bottom-color: beige;
}

nav .menu-item a{
    color: inherit;
    text-decoration: none;
}
</style>