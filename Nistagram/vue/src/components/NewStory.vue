<template>
    <div class="menu-item">
            <h3>New story</h3>
            <input v-if="!itemss[0].image" type="file" @change="onFileChange(itemss[0], $event)">
            
        <div class="container">
             <div class="row" v-for="item in itemss" :key="item">
                <div class="col-sm-6">
                    <img :src="item.image" v-if="item.image" style="width:25%"/>
                    <input class="btn" type="button" value="X" v-if="item.image" @click="removeImage(itemss[item.iid])"/>
                </div>
            </div>
            <div  v-if="itemss[0].image">
                <small v-if="location==null">Add locattion</small>
                <transition name="fade" appear>
                        <div class="sub-menu" v-if="isOpen">
                            <div v-for="l in locationtag" :key="l"  >
                                <button  @click="location=l, isOpen=false">{{l.place}} {{l.city}} {{l.country}}</button>
                            </div>
                        </div>
                    </transition>
                    <input v-if="location==null" id="locationSearch" type="text" @click="isOpenP=false,isOpen=!isOpen" v-on:input="search()"/>
                 
                <small v-if="location!=null">{{this.location.place}} {{this.location.city}} {{this.location.country}}</small>
                <input v-if="location!=null" class="btn" type="button" value="X"  @click="location=null,isOpen=false,isOpenP=false"/>
                   <br>
                   <br>
                <small>Tag people </small>
                <input type="text" @click="isOpen=false,isOpenP=!isOpenP" id="searchTagId" v-on:input="searchTags()"/>
                   <br>
                 <transition name="fade" appear>
                        <div class="sub-menu" v-if="isOpenP" >
                            <div v-for="p in peoplestag" :key="p">
                                <button v-on:click="Taged(p),isOpen=false,isOpenP=false">{{p.username}}</button>
                            </div>
                        </div>
                    </transition>
                   <br>
                   <div v-for="p in tagedPeoples" :key="p">
                       <small>{{p.username}}</small>
                        <input class="btn" type="button" value="X"  @click="removeTag(p),isOpen=false,isOpenP=false"/>
                   </div>
                <small>Descriotion:</small><br>
            <textarea type="text" v-model="post.description" @click="isOpen=false,isOpenP=false"/><br><br><br>
                <input value="Add post" class="btnP" type="button" v-on:click="CreatePost(),isOpen=false,isOpenP=false"/>
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
            itemss:[
       {
         image: false,
         iid:0,
       }
            ],
            isOpen: false,
            isOpenP: false,
            location:null,
            tagedLocation:null,
            peoples:[],
            tagedPeoples:[],
			peoplestag:[],
            locationtag:[],
            images:[],
            hashTags:[],
            post:{
                description:null,
                tagsLink:[],
                hashTags:[],
                location:null,
                userId:null,
                postType:"STORY"
            }
        }
    },
    methods:{
         searchTags(){
            axios
                .get("http://localhost:8080/api/user/searchProfile/" + document.getElementById("searchTagId").value,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
					this.peoplestag = response.data
              })
        },
        search(){
            axios
                .get("http://localhost:8080/api/post/searchLocation/" + document.getElementById("locationSearch").value,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
					this.locationtag = response.data
                    alert("Ispis id")
                    alert(this.locationtag[0].id)
              })
        },
        CreatePost(){

            let b = this.post.description.split(" ")
            for(let i = 0; i < b.length; i++){
                if(b[i].includes("#")){
                    let c = b[i].split("#")
                    for(let j = 1; j < c.length; j++){
                        let h = {
                            name: c[j]
                        }
                        this.post.hashTags.push(h)
                    }
                }
            }
            for(let i = 0; i < this.tagedPeoples.length; i++){
                let ob = {
                    name : this.tagedPeoples[i].username,
                    linkType : "USER_LINK"
                }
                this.post.tagsLink.push(ob)
            }
            this.post.location = this.location
            this.post.userId = localStorage.getItem('userId')
            alert(localStorage.getItem('token'))
            alert(localStorage.getItem('userId'))
            //alert(this.post.userId)
            const params = new FormData()
            for(let i = 0; i < this.images.length; i++)
                params.append('file', this.images[i])
            let json = JSON.stringify(this.post);
            params.append('data', json);
            //alert(params)
        axios
                .post("http://localhost:8080/api/post/upload", params,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
                                'Content-Type': 'multipart/form-data',
							}
				})
                .then(response => {
                  if (response.status==200){
                    alert('Successful');
					this.userr = this.userDto
                  }
              })
               .catch(error => {
                // print(error.status == 417)
                if(error == "Error: Request failed with status code 400"){
                   alert("Error")
                  }
                })
            },
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
    //alert(e.target.files[0])
    this.images.push(e.target.files[0])
    //alert(this.images)
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