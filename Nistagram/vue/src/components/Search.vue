<template>
    <div class="menu-item">
        <input style="width:200px" href="#"  @click="isOpen=!isOpen"/>
    <transition name="fade" appear>
    <div class="sub-menu" v-if="isOpen">
        <TabNav :tabs="['Accounts','Tags','Locations']" :selected="selected" @selected="setSelected">
            <Tab :isSelected="selected==='Accounts'">
                <p style="min-height:100px">Proba1</p>
            </Tab>
            <Tab :isSelected="selected==='Tags'">
                <p style="min-height:100px">Proba2</p>
            </Tab>
            <Tab :isSelected="selected==='Locations'">
                <p style="min-height:100px">Proba3</p> 
            </Tab>
        </TabNav>
    </div>
    </transition>
    </div>
</template>

<script>
import TabNav from '../components/TabNav.vue'
import Tab from '../components/Tab.vue'

export default {
    name:'search',
    components:{TabNav,Tab},
    props:['title','items'],
    data(){
        return{
            isOpen: false,
            selected: 'Accounts'
        }
    },
    methods:{
        setSelected(tab){
            this.selected=tab;
        }
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