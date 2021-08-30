<template>

  <div id="Profile"><br><br>
  <h3 style="text-align:center">Collections</h3><br>
    <div v-for="c in collections" :key="c">
       <button type="submit" class="btnRegister" style="width:150px; margin-left:40px" v-on:click="SaveCollectionPosts(c.posts,c.name)">{{c.name}}</button>
    </div><br><br>
            <div class="col-md-12">
                    <div class="bg-white shadow rounded overflow-hidden divs">
                       <div class="container">
                                        <div class="row">
                                            <div class="col-sm-6" v-for="p in collectionsPosts" :key="p">
                                                <div  v-for="im in p.images" :key="im">
                                                    <video v-if="im.img.includes('video')" controls>
                                                    <source v-bind:src="im.img" type="video/mp4">
                                                    Your browser does not support the video tag.
                                                    </video>
                                                    <img :src="im.img" v-if="im.img.includes('image')"/>
                                                </div>
                                                <h4 id="cardId"><b>{{p.description}}</b></h4>
                                                <img src="../images/bookmark.png" style="margin-left:210px;cursor: pointer;" v-if="IsInCollection(p.id)"  v-on:click="ShowCollections(p)">
                                                <img src="../images/bookmark2.png" style="margin-left:210px;cursor: pointer;" v-else v-on:click="DeletePostFromCollection(p.id)">
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
                                                        <li class="list-inline-item">
                                                            <img src="../images/like.png" v-if="IsLiked(p.likes)" style="margin-left:20px;cursor: pointer;" v-on:click="Like(p.id)">
                                                            <img src="../images/like2.png" v-else style="margin-left:20px;cursor: pointer;" v-on:click="Like(p.id)">
                                                        </li>
                                                         <li class="list-inline-item">
                                                            <img src="../images/dislike.png" v-if="IsLiked(p.dislikes)" style="margin-right:110px;cursor: pointer;" v-on:click="Dislike(p.id)">
                                                            <img src="../images/dislike2.png" v-else style="margin-right:110px;cursor: pointer;" v-on:click="Dislike(p.id)">
                                                        </li>
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
                                                 <input id="addComment" type="text" v-model="comment.content"/>
                                                 <button  v-on:click="AddComment(p.id)">Add comment</button>
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
                  </div>
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

    <!--<div id="app">
      <button id="show-modal" @click="showModal = true">Show Modal</button>

      <modal v-if="showModal" @close="showModal = false">-->

        <!--<h3 slot="header">custom header</h3>-->
     <!-- </modal>
    </div>-->
</div>			
</template>

<script>

const axios=require('axios')

export default {
  name: 'Collections',
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
             collectionsPosts:[],
             collectionName:null
		}
	},beforeMount() {
    alert(parseInt(localStorage.getItem('userId')))

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
                    for(let c = 0; c < this.collections.length; c++){
                      if(this.collections[c].posts != null){
                          for(let j = 0; j < this.collections[c].posts.length; j++){
                              for (let i = 0; i < this.collections[c].posts[j].images.length; i++) {
                                      if(this.collections[c].posts[j].images[i].filename.includes('mp4')){
                                          this.collections[c].posts[j].images[i].img = 'data:video/mp4;base64,' + this.collections[c].posts[j].images[i].img
                                      }else{
                                          this.collections[c].posts[j].images[i].img = 'data:image/png;base64,' + this.collections[c].posts[j].images[i].img
                                      }
                              }
                          }
                      }
                    }
                  
                }
              })
	},
  methods: {
      Edit(){
		this.$router.push('/');
    },
    SaveCollectionPosts(posts, name){
        this.collectionsPosts = posts
        this.collectionName = name
    },
    DeletePostFromCollection(id){
              let c = {
                name:this.collectionName,
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

</style>
