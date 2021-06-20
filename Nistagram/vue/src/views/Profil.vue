<template>
  <div id="Profile">
            <div class="col-md-12">
                    <div class="bg-white shadow rounded overflow-hidden divs">
                        <div class="px-4 pt-0 pb-4 cover">
                                <div class=""><img  src="https://i.pinimg.com/originals/11/5f/4f/115f4f233582670e085966ee8250e75f.png" alt="..." width="130"  class="rounded-circle">
                                </div>
                                <input type="submit" class="btnRegister"  value="Add story" v-on:click="AddStory()"/>
                                <div class="media-body mb-5">     
                                    <div class="p-4 d-flex justify-content-end text-center">
                                        <ul class="list-inline mb-0">
                                            <li class="list-inline-item">
                                                <h5 class="font-weight-bold mb-0 d-block">215</h5><small class="text-muted"><i class="fa fa-image mr-1"></i>Photos</small>
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
                                                <img src="https://api.time.com/wp-content/uploads/2016/05/gettyimages-502649088.jpg?w=824&quality=70" width="60%" height="60%">
                                                <h4 id="cardId"><b>Vacation</b></h4>
                                                <div class="bg-light d-flex text-center">
                                                    <ul class="list-inline mb-0">
                                                        <li class="list-inline-item">
                                                            <h5 class="font-weight-bold mb-0 d-block">215</h5><small class="text-muted"><i class="fa fa-thumbs-up"></i>Likes</small>
                                                        </li>
                                                        <li class="list-inline-item">
                                                            <h5 class="font-weight-bold mb-0 d-block">745</h5><small class="text-muted"><i class="fa fa-thumbs-down"></i>Dislikes</small>
                                                        </li>
                                                        <li class="list-inline-item">
                                                            <h5 class="font-weight-bold mb-0 d-block">340</h5><small class="text-muted"><i class="fa fa-comments"></i>Comments</small>
                                                        </li>
                                                    </ul>
                                                </div><br><br>
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

export default {
  name: 'Profil',
  components: {
      NewPhoto
  },
  data: function () {
		return {
            usernameText: null,
			passwordText: null,
			loginResponse: null,
			token: null,
			role:null,
			jwtAuthenticationRequest:{},
            pict:[1,2,3,4]
		}
	},beforeMount() {
		localStorage.setItem('isLogged', false);
		localStorage.setItem('userId', "");


	},
  methods: {
      Edit(){
		this.$router.push('/');
    }
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
