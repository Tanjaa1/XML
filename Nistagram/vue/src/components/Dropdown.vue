<template>
    <div class="menu-item" @click="isOpen=!isOpen">
        <button href="#">
            {{title}}
        </button>
        <svg viewBox="0 0 1030 638">
            <path d="M1017 68L541 626q-11 12-26 12t-26-12L13
            68Q-3 49 6 24.5T39 0h952q24 0 33 24.5t-7 43.5z"
            fill="#FFF"></path>
        </svg>
    <transition name="fade" appear style="height:250px">
    <div class="sub-menu" v-if="isOpen">
        <div v-for="(item,i) in items" :key="i" class="menu-item">
            <a v-on:click="Find(item)">{{item.title}}</a>
        </div>
    </div>
    </transition>
    </div>
</template>

<script>
export default {
    name:'dropdown',
    props:['title','items'],
    data(){
        return{
            isOpen: false
        }
    },
    methods:{        
        Find(item) {
            if(item.title == 'Log Out'){
                localStorage.setItem('userId', null)
                localStorage.setItem('username', '')
                localStorage.setItem('token', null)
            }
            this.$router.push(item.link)
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