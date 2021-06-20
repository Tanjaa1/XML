<template>
    <div class="menu-item">
        <input type="file" accept="image/*" @click="isOpen=!isOpen" @change="uploadImage($event)" id="file-input">
    <transition name="fade" appear>
    <div class="sub-menu" v-if="isOpen">
       Proba
    </div>
    </transition><br>
    </div>
</template>

<script >
const axios = require('axios');
export default {
    name:'newPhoto',
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
        },

  uploadImage(event) {

    const URL = 'http://foobar.com/upload'; 

    let data = new FormData();
    data.append('name', 'my-picture');
    data.append('file', event.target.files[0]); 

    let config = {
      header : {
        'Content-Type' : 'image/png'
      }
    }

    axios.put(
      URL, 
      data,
      config
    ).then(
      response => {
        console.log('image upload response > ', response)
      }
    )
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
.btnRegister{
    border: none;
    border-radius: 1.5rem;
    background: #55556a;
    color: #fff;
    cursor: pointer;
}
</style>