<template>
    <div class="menu-item">
        <input style="width:200px" href="#" id="searchTagId" @click="isOpen=!isOpen"  v-on:input="search()"/>
    <transition name="fade" appear>
    <div class="sub-menu" v-if="isOpen">
        <TabNav :tabs="['Accounts','Tags','Locations']" :selected="selected" @selected="setSelected">
            <Tab :isSelected="selected==='Accounts'">
                <p  v-for="p in profiles" :key="p">
                        <button class="btn" v-on:click="SaveIdUser(p)">{{p.username}} </button>
                    </p>
            </Tab>
            <Tab :isSelected="selected==='Tags'">
                 <p  v-for="p in tags" :key="p">
                        <button class="btn" >{{p.name}} </button>
                    </p>
            </Tab>
            <Tab :isSelected="selected==='Locations'">
                <p  v-for="p in locations" :key="p">
                        <button class="btn" v-on:click="SaveLocation(p)">{{p.place}},{{p.city}},{{p.country}} </button>
                    </p>
            </Tab>
        </TabNav>
    </div>
    </transition>
    </div>
</template>

<script>
export var myVar
export var myLocation

const axios=require('axios')
import TabNav from './TabNav.vue'
import Tab from '../components/Tab.vue'

export default {
    name:'search',
    components:{TabNav,Tab},
    props:['title','items'],
    data(){
        return{
            isOpen: false,
            selected: 'Accounts',
            tags:[],
            profiles:[],
            locations:[]
        }
    },
    methods:{
        SaveIdUser(p){
            myVar = p
            if(p.id == localStorage.getItem('userId'))
                this.$router.push('/Profil')
            else
                this.$router.push('/ProfilUser');
        },
         SaveLocation(p){
            myLocation = p
            this.$router.push('/PostsByLocation');
        },
        setSelected(tab){
            this.selected=tab;
        },
        search(){
             axios
                .get("http://localhost:8080/api/user/searchProfile/" + document.getElementById("searchTagId").value,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
					this.profiles = response.data
              })
            axios
                .get("http://localhost:8080/api/post/searchLocation/" + document.getElementById("searchTagId").value,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
					this.locations = response.data
              })

              axios
                .get("http://localhost:8080/api/post/searchHashtag/" + document.getElementById("searchTagId").value,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
					this.tags = response.data
              })
        },
    }
}
</script>

<style>
nav .menu-item svg{
    width: 10px;
    margin-left: 10px;
}

nav .menu-item .sub-menu{
    position: absolute;
    background-color: #4c4c5a;
    top:calc(100% +18px);
    left:0%;
    transform: translateY(20%);
    width:max-content;
    border-radius: 0px 0px 16px 16px;
}
.btn{
    background-color: #4c4c5a;
}
.fade-enter-active,
.fade-leave-active{
    transition: all .5s ease-out;
}

.fade-enter
.fade-leave-to{
    opacity: 0;
    transform: translateX(100vw);
}
</style>