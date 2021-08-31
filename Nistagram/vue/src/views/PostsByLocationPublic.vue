<template>

  <div id="Profile">
            <div class="col-md-12">
                    <div class="bg-white shadow rounded overflow-hidden divs">
                        <div class="px-4 pt-0 pb-4 cover">
                                <div class=""><img v-on:click="ShowStory()" style="cursor: pointer;" src="https://i.pinimg.com/originals/11/5f/4f/115f4f233582670e085966ee8250e75f.png" alt="..." width="130"  class="rounded-circle">
                                </div>
                               
                        </div><br>
  
                        <h2 class="mt-0 mb-0">{{locationName}}</h2>           
                        <div>
                            <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
                            <ul class="nav nav-tabs">
                                <li class="active"><a data-toggle="tab" href="#home">Pictures</a></li>
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
                                                            <h5 v-if="p.likes != null" class="font-weight-bold mb-0 d-block" @click="likesModal = p.likes,showModal = true" style="cursor: pointer;">{{p.likes.length}}</h5>
                                                            <h5 v-if="p.likes == null" class="font-weight-bold mb-0 d-block">0</h5>
                                                            <small class="text-muted"><i class="fa fa-thumbs-up"></i>Likes</small>
                                                        </li>
                                                        <li class="list-inline-item">
                                                            <h5 v-if="p.dislikes != null" class="font-weight-bold mb-0 d-block"  @click="likesModal = p.dislikes,showModal = true" style="cursor: pointer;">{{p.dislikes.length}}</h5>
                                                            <h5 v-if="p.dislikes == null" class="font-weight-bold mb-0 d-block">0</h5>
                                                            <small class="text-muted"><i class="fa fa-thumbs-down"></i>Dislikes</small>
                                                        </li>
                                                        <li class="list-inline-item">
                                                            <h5 v-if="p.comments != null" class="font-weight-bold mb-0 d-block">{{p.comments.length}}</h5>
                                                            <h5 v-if="p.comments == null" class="font-weight-bold mb-0 d-block">0</h5>
                                                            <small class="text-muted"><i class="fa fa-comments"></i>Comments</small>
                                                        </li><br>
                                    
                                                    </ul>
                                                </div><br>
                                                <h5 id="cardId"><b>{{p.location.place}} {{p.location.city}} {{p.location.country}}</b></h5><br>
                                                Taged people:
                                                <div v-for="t in p.tagsLink" :key="t">
                                                    @{{t.name}}
                                                </div>
                                                <br>Comments:
                                                <div v-for="c in p.comments" :key="c" style="word-wrap: break-word;width: 300px;">
                                                    {{c.username}}: {{c.content}}<br>
                                                </div><br>
                                                
                                            </div>
                                        </div>
                                    </div>       
                                </div>                                    
                                
                            </div>
                        </div>
                    </div>
                </div>
              
        <div class="modal-mask" v-if="showModal">
          <div class="modal-wrapper">
            <div class="modal-container">

              <div class="modal-header">
                <slot name="header">
                  Users
                </slot>
              </div>

              <div class="modal-body">
                <slot name="body">
                  <div v-for="l in likesModal" :key="l">
                      {{l}}<br>
                  </div>
                </slot>
              </div>

              <div class="modal-footer">
                <slot name="footer">
               
                  <button class="modal-default-button" @click="showModal=false">
                    Close
                  </button>
                </slot>
              </div>
            </div>
          </div>
        </div>

         <div class="modal-mask" v-if="showModalCollections">
          <div class="modal-wrapper">
            <div class="modal-container">

              <div class="modal-header">
                <slot name="header">
                  Collections
                </slot>
              </div>

              <div class="modal-body">
                <slot name="body">
                  <div v-for="l in collections" :key="l">
                   <a v-on:click="AddInCollection(l.name)" style="cursor: pointer;">{{l.name}}</a><br>
                  </div><hr>
                  <input type="text" v-model="newCollection"/>
                  <button class="modal-default-button" v-on:click="CreateNewCollection()">
                    New
                  </button>
                </slot>
              </div>

              <div class="modal-footer">
                <slot name="footer">
                  <button class="modal-default-button" @click="showModalCollections=false">
                    Close
                  </button>
                </slot>
              </div>
            </div>
          </div>
        </div>

        <!--Modal story-->
     <div class="modal-mask" v-if="storyModal">
          <div class="modal-wrapper">
            <div class="modal-container">


              <div class="modal-body">
                <slot name="body">
                    <input type="button" value="<" v-on:click="Previous()"/>
                    <img :src="stories[i].images[0].img" style="width:160px; height:150px"/>     
                    <input type="button" value=">" v-on:click="Next()"/><br>
                      <h4 id="cardId"><b>{{stories[i].description}}</b></h4><br>
                      <h5 id="cardId"><b>{{stories[i].location.place}} {{stories[i].location.city}} {{stories[i].location.country}}</b></h5><br>
                    Taged people:
                   <div v-for="t in stories[i].tagsLink" :key="t">
                   @{{t.name}}
                   </div>
                </slot>
              </div>

              <div class="modal-footer">
                <slot name="footer">
                  <button class="modal-default-button" @click="storyModal=false">
                    Close
                  </button>
                  </slot>
              </div>
            </div>
          </div>
        </div>


        <!--Modal Highlights-->
        <div class="modal-mask" v-if="highlightsModal">
          <div class="modal-wrapper">
            <div class="modal-container">


              <div class="modal-body">
                <slot name="body">
                    <input type="button" value="<" v-on:click="PreviousH()"/>
                    <img :src="storiesHighlight[j].images[0].img" style="width:160px; height:150px"/>     
                    <input type="button" value=">" v-on:click="NextH()"/><br>
                         <h4 id="cardId"><b>{{storiesHighlight[j].description}}</b></h4><br>
                         <h5 id="cardId"><b>{{storiesHighlight[j].location.place}} {{storiesHighlight[j].location.city}} {{storiesHighlight[j].location.country}}</b></h5><br>
                    Taged people:
                   <div v-for="t in storiesHighlight[j].tagsLink" :key="t">
                   @{{t.name}}
                   </div>
                </slot>
              </div>

              <div class="modal-footer">
                <slot name="footer">
                  <button class="modal-default-button" @click="highlightsModal=false">
                    Close
                  </button>
                  </slot>
              </div>
            </div>
          </div>
        </div>
        <!--Modal Highlights end-->
</div>			
</template>

<script>

import { myLocation } from '../components/Search.vue'

const axios=require('axios')

export default {
  name: 'PostsByLocationPublic',
  components: {
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
            },
             showModal: false,
             like:{
               postId:null,
               userId:null,
               username:null,
               linkType:null,
             },
             likesModal:[],
             dislikesModal:[],
             post:null,
             collections:[],
             showModalCollections:false,
             newCollection:null,
             stories:[],
             storyModal:false,
             replace: true,
             highlightPICTURE:[1,2,3],
             i:0,
             highlights:[],
             highlightsModal:false,
             j:0,
             storiesHighlight:[],
             locationId:null,
             locationName:''
		}
	},async beforeMount() {
        this.locationId = myLocation.id
        this.locationName = myLocation.place + ' ' + myLocation.city + ' ' + myLocation.country
       await  axios
                .get("http://localhost:8080/api/post/getPostsByLocation/" + this.locationId + '/' + '1',
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
                  if (response.status==200){
                    for(let a = 0; a < response.data.length; a++){
                         axios
                              .get("http://localhost:8080/api/user/checkPrivate/" + response.data[a].userId,
                      {
                            headers: {
                              'Authorization': 'Bearer' + " " + localStorage.getItem('token')
                            }
                      })
                              .then(response1 => {
                                if (response1.status==200){
                                  if(!response1.data){
                                     for (let i = 0; i < response.data[a].images.length; i++) {
                                        if(response.data[a].images[i].filename.includes('mp4')){
                                            response.data[a].images[i].img = 'data:video/mp4;base64,' + response.data[a].images[i].img
                                            alert(response.data[a].images[i].img)
                                        }else{
                                            response.data[a].images[i].img = 'data:image/png;base64,' + response.data[a].images[i].img
                                            
                                        }
                          }
                                    this.pict.push(response.data[a])
                                }
                                
                              }
                            })
                    }
                  }
              })



                await  axios
                .get("http://localhost:8080/api/post/getStoriesByLocation/" + this.locationId + '/' + localStorage.getItem('userId'),
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
                  if (response.status==200){
                    for(let a = 0; a < response.data.length; a++){
                         axios
                              .get("http://localhost:8080/api/user/checkPrivate/" +response.data[a].userId,
                      {
                            headers: {
                              'Authorization': 'Bearer' + " " + localStorage.getItem('token')
                            }
                      })
                              .then(response1 => {
                                if (response1.status==200){
                                  if(!response1.data){
                                     for (let i = 0; i < response.data[a].images.length; i++) {
                            
                                        if(response.data[a].images[i].filename.includes('mp4')){
                                            response.data[a].images[i].img = 'data:video/mp4;base64,' + response.data[a].images[i].img
                                            alert(response.data[a].images[i].img)
                                        }else{
                                            response.data[a].images[i].img = 'data:image/png;base64,' + response.data[a].images[i].img
                                            
                                        }
                          }
                                    this.stories.push(response.data[a])
                                }
                                
                              }
                            })
                    }
                  }
              })


  },
  methods: {
      Edit(){
		this.$router.push('/');
    },
    ShowHighlightStories(l){
      this.storiesHighlight = l
    },
    ShowStory(){
      this.storyModal = true
    },
    Previous(){
      if((this.i-1)<0)
        this.i=this.stories.length-1
      else
        this.i=this.i-1
    },
    Next(){
      if((this.i+1)<this.stories.length)
          this.i=this.i+1
      else
        this.i=0
    },
     PreviousH(){
      if((this.j-1)<0)
        this.j=this.storiesHighlight.length-1
      else
        this.j=this.j-1
    },
    NextH(){
      if((this.j+1)<this.storiesHighlight.length)
          this.j=this.j+1
      else
        this.j=0
    },
    CreateNewCollection(){
      if(this.newCollection != null){
          let c ={
            name:this.newCollection,
            userId: parseInt(localStorage.getItem('userId')),
            posts:[parseInt(this.post.id)]
          }
           axios
                .post("http://localhost:8080/api/post/createCollection",c,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
							}
				})
                .then(response => {
                  if (response.status==201){
                    location.reload()
                  }
              })
               .catch(error => {
                // print(error.status == 417)
                if(error == "Error: Request failed with status code 400"){
                   alert("Error")
                  }
                })
      } 
      else
        alert("Name of collection is empty")
    },
    DeletePostFromCollection(id){
      for(let i = 0; i < this.collections.length; i++){
        if(this.collections[i].posts != null){
          for(let j = 0; j < this.collections[i].posts.length; j++){
            if(this.collections[i].posts[j].id == id){
              let c = {
                name:this.collections[i].name,
                postId:parseInt(id)
              }
              axios
                        .post("http://localhost:8080/api/post/removeFromCollection",c,
                {
                      headers: {
                        'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
                      }
                })
                        .then(response => {
                          if (response.status==201){
                            location.reload()
                          }
                      })
                      .catch(error => {
                        // print(error.status == 417)
                        if(error == "Error: Request failed with status code 400"){
                          alert("Error")
                          location.reload()
                          }
                        })
            }
          }
        }
      }
    },
    IsInCollection(id){
      if(this.collections != null){
        for(let i = 0; i < this.collections.length; i++){
          if(this.collections[i].posts != null){
            for(let j = 0; j < this.collections[i].posts.length; j++){
              if(this.collections[i].posts[j].id == id)
                return false
            }
          }
        }
      }
      return true
    },
    AddInCollection(n){
      let c = {
        name:n,
        postId:parseInt(this.post.id)
      }
       axios
                .post("http://localhost:8080/api/post/addIntoCollection",c,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
							}
				})
                .then(response => {
                  if (response.status==201){
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
    ShowCollections(p){
      this.post = p
       axios
                .get("http://localhost:8080/api/post/GetCollectionsByUserId/" + parseInt(localStorage.getItem('userId')),
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
				})
                .then(response => {
                  if (response.status==200){
                    this.collections = response.data
                    this.showModalCollections = true
                  
                }
              })
    },
    IsLiked(likes){
      if(likes != null){
        for(let i = 0; i < likes.length; i++){
          if(likes[i] == localStorage.getItem('username'))
            return false
        }
      }
      return true
    },
    Dislike(postId){
      this.like.username = localStorage.getItem('username')
      this.like.userId = parseInt(localStorage.getItem('userId'))
      this.like.postId = postId
      this.like.linkType = "DISLIKE"
      axios
                .post("http://localhost:8080/api/post/addLike", this.like,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
							}
				})
                .then(response => {
                  if (response.status==201){
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
    Like(postId){
      this.like.username = localStorage.getItem('username')
      this.like.userId = parseInt(localStorage.getItem('userId'))
      this.like.postId = postId
      this.like.linkType = "LIKE"
      axios
                .post("http://localhost:8080/api/post/addLike", this.like,
				{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token'),
							}
				})
                .then(response => {
                  if (response.status==201){
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
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
}

.modal-default-button {
  float: right;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}

#slideLandingPage {
    font-family: Impact;
    font-size: 24px;
    background-color: lightblue;
    background-image: none;
    opacity : 0.9;
}






#slides .inner {
  width: 400%;
}

#slides .inner {
  -webkit-transform: translateZ(0);
  -moz-transform: translateZ(0);
  -o-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);

  -webkit-transition: all 800ms cubic-bezier(0.770, 0.000, 0.175, 1.000);
  -moz-transition: all 800ms cubic-bezier(0.770, 0.000, 0.175, 1.000);
  -o-transition: all 800ms cubic-bezier(0.770, 0.000, 0.175, 1.000);
  -ms-transition: all 800ms cubic-bezier(0.770, 0.000, 0.175, 1.000);
  transition: all 800ms cubic-bezier(0.770, 0.000, 0.175, 1.000);

  -webkit-transition-timing-function: cubic-bezier(0.770, 0.000, 0.175, 1.000);
  -moz-transition-timing-function: cubic-bezier(0.770, 0.000, 0.175, 1.000);
  -o-transition-timing-function: cubic-bezier(0.770, 0.000, 0.175, 1.000);
  -ms-transition-timing-function: cubic-bezier(0.770, 0.000, 0.175, 1.000);
  transition-timing-function: cubic-bezier(0.770, 0.000, 0.175, 1.000);
}

#slides article {
  width: 25%;
  float: left;
}

#slide1:checked ~ #slides .inner {
  margin-left: 0;
}

#slide2:checked ~ #slides .inner {
  margin-left: -100%;
}

#slide3:checked ~ #slides .inner {
  margin-left: -200%;
}

#slide4:checked ~ #slides .inner {
  margin-left: -300%;
}

input[type="radio"] {
  display: none;
}

label {
  background: #CCC;
  display: inline-block;
  cursor: pointer;
  width: 10px;
  height: 10px;
  border-radius: 5px;
}

#slide1:checked ~ label[for="slide1"],
#slide2:checked ~ label[for="slide2"],
#slide3:checked ~ label[for="slide3"],
#slide4:checked ~ label[for="slide4"] {
  background: #333;
}
</style>
