<template>
    <div class="menu-item">
            <h3>New story</h3>
            <input v-if="!items.image" type="file" @change="onFileChange(items, $event)">
            
        <div class="container">
            <div class="row" v-for="item in items" :key="item">
                <div class="col-sm-6">
                    <img :src="items.image" v-if="items.image" style="width:25%"/>
                    <input class="btn" type="button" value="X" v-if="item.image" @click="removeImage(items)"/>
                </div>
            </div>
            <div  v-if="items.image">
                <small v-if="location==null">Add locattion</small>
                <transition name="fade" appear>
                        <div class="sub-menu" v-if="isOpen">
                            <div v-for="l in locationtag" :key="l"  >
                                <button  @click="location=l, isOpen=false">{{l}}</button>
                            </div>
                        </div>
                    </transition>
                    <input v-if="location==null" type="text" @click="isOpenP=false,isOpen=!isOpen"/>
                 
                <small v-if="location!=null">{{this.location}}</small>
                <input v-if="location!=null" class="btn" type="button" value="X"  @click="location=null,isOpen=false,isOpenP=false"/>
                   <br>
                   <br>
                <small>Tag people </small>
                <input type="text" @click="isOpen=false,isOpenP=!isOpenP"/>
                   <br>
                 <transition name="fade" appear>
                        <div class="sub-menu" v-if="isOpenP" >
                            <div v-for="p in peoplestag" :key="p">
                                <button v-on:click="Taged(p),isOpen=false,isOpenP=false">{{p}}</button>
                            </div>
                        </div>
                    </transition>
                   <br>
                   <div v-for="p in tagedPeoples" :key="p">
                       <small>{{p}}</small>
                        <input class="btn" type="button" value="X"  @click="removeTag(p),isOpen=false,isOpenP=false"/>
                   </div>
                <small>Descriotion:</small><br>
            <textarea type="text" @click="isOpen=false,isOpenP=false"/><br><br><br>
                <input value="Add post" class="btnP" type="button" v-on:click="NewPhoto(),isOpen=false,isOpenP=false"/>
            </div>
        </div><br>
    </div>
</template>

<script >
const axios = require('axios');
export default {
    name:'newStory',
    data(){
        return{
            items:
       {
         image: false,
         iid:0,
       },
            isOpen: false,
            isOpenP: false,
            location:null,
            tagedLocation:null,
            peoples:[],
            tagedPeoples:[],
			peoplestag:['pera','mika','janko'],
            locationtag:['Novi Sad','Beograd','FTN']
        }
    },
    methods:{
        setSelected(tab){
            this.selected=tab;
        },
        Taged(p){
            if(this.tagedPeoples.indexOf(p, 0)==-1)
                this.tagedPeoples.push(p)
        },
        removeTag(p){
            this.tagedPeoples.splice(this.tagedPeoples.indexOf(p, 0),1)
        },
   onFileChange(item, e) {

       const URL = 'http://foobar.com/upload'; 

    let data = new FormData();
    data.append('name', 'my-picture');
    data.append('file', e.target.files[0]); 

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

      var files = e.target.files || e.dataTransfer.files;
      if (!files.length)
        return;
      this.createImage(item, files[0]);
    },
    createImage(item, file) {
      var reader = new FileReader();

      reader.onload = (e) => {
        item.image = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    removeImage: function (item) {
      item.image = false; 
    },
    removeLocation(){
        this.location=null
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
.btn{
    border: none;
    border-radius: 1.5rem;
    color:rgb(159, 112, 112);
    cursor: pointer;
}
.choose{
    padding-left: 20%;
}

.sub-menu{
    position: absolute;
    background-color: #9191c4;
    width:150px;
    height: 100px;
    transform: translateX(6.5vw);
}

.btnP{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    background: #4608d6; 
    color: #fff;
    width: 20%;
    cursor: pointer;
}
</style>