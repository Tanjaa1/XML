<template>
  <div id="Profile">
            <div class="col-md-12">
                    <div class="bg-white shadow rounded overflow-hidden divs">
                        <div class="px-4 pt-0 pb-4 cover">
                                <div class=""><img  src="https://i.pinimg.com/originals/11/5f/4f/115f4f233582670e085966ee8250e75f.png" alt="..." width="130"  class="rounded-circle">
                                </div>
                                <div class="media-body mb-5">     
                                    <div class="p-4 d-flex justify-content-end text-center">
                                        <ul class="list-inline mb-0">
                                            <li class="list-inline-item">
                                                <h5 class="font-weight-bold mb-0 d-block">{{pict.length}}</h5><small class="text-muted"><i class="fa fa-image mr-1"></i>Photos</small>
                                            </li>
                                            <li class="list-inline-item">
                                                <h5 class="font-weight-bold mb-0 d-block">745</h5><small class="text-muted"> <i class="fa fa-user mr-1"></i>Followers</small>
                                            </li>
                                            <li class="list-inline-item">
                                                <h5 class="font-weight-bold mb-0 d-block">340</h5><small class="text-muted"> <i class="fa fa-user mr-1"></i>Following</small>
                                            </li>
                                        </ul>
                                    </div>           
                                </div>
                                <NewStory/>
                        </div><br>
                        <h2 class="mt-0 mb-0">Homer Simpson</h2>
                        <p> Clumsy, fat and very lazy, also an alcoholic, and  not very intelligent, but i love donuts</p>
                        <NewPhoto/>           
                        <div>
                            <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
                            <ul class="nav nav-tabs">
                                <li class="active"><a data-toggle="tab" href="#home">Pictures</a></li>
                                <li><a data-toggle="tab" href="#menu1">Videos</a></li>
                                <li><a data-toggle="tab" href="#menu2">Saved</a></li>
                            </ul>
                            <div class="tab-content">
                                <div id="home" class="tab-pane fade in active">
                                    <div class="container">
                                        <div class="row">
                                            <div class="col-sm-6" v-for="p in pict" :key="p">
                                                <div  v-for="im in p.images" :key="im">
                                                    <video v-if="im.img.includes('video')" controls>
                                                    <source v-bind:src="im.img" type="video/mp4">
                                                    Your browser does not support the video tag.
                                                    </video>
                                                    <img :src="im.img" v-if="im.img.includes('image')"/>
                                                </div>
                                                <h4 id="cardId"><b>{{p.description}}</b></h4>
                                                <div class="d-flex text-center">
                                                    <ul class="list-inline mb-0">
                                                        <li class="list-inline-item">
                                                            <h5 v-if="p.likes != null" class="font-weight-bold mb-0 d-block">{{p.likes.length}}</h5>
                                                            <h5 v-if="p.likes == null" class="font-weight-bold mb-0 d-block">0</h5>
                                                            <small class="text-muted"><i class="fa fa-thumbs-up"></i>Likes</small>
                                                        </li>
                                                        <li class="list-inline-item">
                                                            <h5 v-if="p.dislikes != null" class="font-weight-bold mb-0 d-block">{{p.dislikes.length}}</h5>
                                                            <h5 v-if="p.dislikes == null" class="font-weight-bold mb-0 d-block">0</h5>
                                                            <small class="text-muted"><i class="fa fa-thumbs-down"></i>Dislikes</small>
                                                        </li>
                                                        <li class="list-inline-item">
                                                            <h5 v-if="p.comments != null" class="font-weight-bold mb-0 d-block">{{p.comments.length}}</h5>
                                                            <h5 v-if="p.comments == null" class="font-weight-bold mb-0 d-block">0</h5>
                                                            <small class="text-muted"><i class="fa fa-comments"></i>Comments</small>
                                                        </li>
                                                    </ul>
                                                </div><br>Comments:
                                                <div v-for="c in p.comments" :key="c" style="word-wrap: break-word;width: 300px;">
                                                    {{c.username}}: {{c.content}}<br>
                                                </div><br>
                                                 <input id="addComment" type="text" v-model="comment.content"/>
                                                 <button  v-on:click="AddComment(p.id)">Add comment</button>
                                            </div>
                                        </div>
                                    </div>       
                                </div>                                    
                                <div id="menu1" class="tab-pane fade">
                                    <h3>Menu 1</h3>
                                </div>
                                <div id="menu2" class="tab-pane fade">
                                    <h3>Menu 2</h3>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
        </div>			
</template>

<script>
import NewPhoto  from "../components/NewPhoto.vue"
import NewStory  from "../components/NewStory.vue"

const axios=require('axios')

export default {
  name: 'Profil',
  components: {
      NewPhoto,NewStory
  },
  data: function () {
		return {
            usernameText: null,
			passwordText: null,
			loginResponse: null,
			token: null,
			role:null,
			jwtAuthenticationRequest:{},
            pict:[],
            likes:null,
            im:null,
            comment:{
                username:null,
                content:null
            }
		}
	},beforeMount() {
        alert(localStorage.getItem('userId'))
        axios
                .get("http://localhost:8080/api/post/getPByUserId/" + localStorage.getItem('userId'),
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
                  if (response.status==200){
					this.pict = response.data
                    for(let j = 0; j < this.pict.length; j++){
                        for (let i = 0; i < this.pict[j].images.length; i++) {
                           // pom1 = this.pict[j].images[i].filepath.split('\\')
                            //if (pom1.length == 4) {
                                //this.pict[j].images[i].filepath = pom1[1] + '/' + pom1[2] + '/' + pom1[3]
                                alert(this.pict[j].images[i].filename)
                                if(this.pict[j].images[i].filename.includes('mp4')){
                                    this.pict[j].images[i].img = 'data:video/mp4;base64,' + this.pict[j].images[i].img
                                    alert(this.pict[j].images[i].img)
                                }else{
                                    this.pict[j].images[i].img = 'data:image/png;base64,' + this.pict[j].images[i].img
                                    alert(this.pict[j].images[i].img)
                                }
                            //}
                        }
                    }
                }
              })
	},
  methods: {
      Edit(){
		this.$router.push('/');
    },
    AddComment(id){
        this.comment.username = localStorage.getItem('username')
        axios
                .post("http://localhost:8080/api/post/addComment/" + id, this.comment,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
							}
				})
                .then(response => {
                  if (response.status==200){
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
  }
}
</script>

<style scoped>
.container{
    background: -webkit-linear-gradient(left, #eeeef5, #fcfeff);
}
.btnRegister{
    border: none;
    border-radius: 1.5rem;
    background: #55556a;
    color: #fff;
    cursor: pointer;
}
</style>
