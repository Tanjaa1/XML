<template>
    <div class="menu-item">
            <h3>New post</h3>
        <div v-if="!items[0].image"><br>
            <input type="file" @change="onFileChange(items[0], $event)">
            </div>
            <div v-else>

            <div v-if="!items[1].image">
                <input type="file" @change="onFileChange(items[1], $event)">
            </div>
            <div v-else>

                <div v-if="!items[2].image">
                    <input type="file" @change="onFileChange(items[2], $event)">
                </div>
                <div v-else>

                    <div v-if="!items[3].image">
                        <input type="file" @change="onFileChange(items[3], $event)">
                    </div>
                    <div v-else>

                        <div v-if="!items[4].image">
                            <input type="file" @change="onFileChange(items[4], $event)">
                        </div>

                    </div>

                </div>
            
            </div>
            <br>
        </div>
        <div class="container">
            <div class="row" v-for="item in items" :key="item">
                <div class="col-sm-6">
                    <video v-if="item.image.includes('video')" style="width:45%" controls>
					<source v-bind:src="item.image" type="video/mp4">
					Your browser does not support the video tag.
					</video>
                    <img :src="item.image" v-if="item.image.includes('image')" style="width:25%"/>
                    <input class="btn" type="button" value="X" v-if="item.image != ''" @click="removeImage(items[item.iid])"/>
                </div>
            </div>
            <div  v-if="items[0].image || items[1].image || items[2].image || items[3].image || items[4].image">
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
                <input type="text" @click="isOpen=false,isOpenP=!isOpenP" id="ss" v-on:input="searchTags()"/>
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
    name:'newPhoto',
    data(){
        return{
            items: [
       {
         image: '',
         iid:0,
       },
       {
         image: '',
         iid:1,
       },
       {
         image: '',
         iid:2,
       },
       {
         image: '',
         iid:3,
       },
       {
         image: '',
         iid:4,
       },
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
                postType:"POST"
            }
        }
    },
    beforeMount(){
    },
    methods:{
        searchTags(){
              axios
                .get("http://localhost:8080/api/user/searchProfile/" + document.getElementById("ss").value,
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
            const params = new FormData()
            for(let i = 0; i < this.images.length; i++)
                params.append('file', this.images[i])
            let json = JSON.stringify(this.post);
            params.append('data', json);
        axios
                .post("http://localhost:8080/api/post/upload", params,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
                                'Content-Type': 'multipart/form-data',
							}
				})
                .then(response => {
                  if (response.status==201){
                    alert('Successful');
					location.reload()
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
      item.image = ''; 
    },
        Taged(p){
            if(this.tagedPeoples.indexOf(p, 0)==-1)
                this.tagedPeoples.push(p)
        },
        removeTag(p){
            this.tagedPeoples.splice(this.tagedPeoples.indexOf(p, 0),1)
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