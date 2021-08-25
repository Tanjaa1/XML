<template>
  <div id="Profile">
      
            <div class="col-md-12">
                    <div class="bg-white shadow rounded overflow-hidden divs">                    
                        <div>
                            <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
                            <ul class="nav nav-tabs">
                                <li class="active"><a data-toggle="tab" href="#home">Liked</a></li>
                                <li><a data-toggle="tab" href="#menu1">Disliked</a></li>
                            </ul>
                            <div class="tab-content">
                                <div id="home" class="tab-pane fade in active">
                                    <div class="container">
                                        <div class="row">
                                            <div class="col-sm-6" v-for="p in pict" :key="p">
                                                <img src="https://api.time.com/wp-content/uploads/2016/05/gettyimages-502649088.jpg?w=824&quality=70" width="60%" height="60%">
                                                <h4 id="cardId"><b>Vacation</b></h4>
                                                <div class="d-flex text-center">
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
                                    <div class="container">
                                        <div class="row">
                                            <div class="col-sm-6" v-for="p in pict" :key="p">
                                                <img src="https://api.time.com/wp-content/uploads/2016/05/gettyimages-502649088.jpg?w=824&quality=70" width="60%" height="60%">
                                                <h4 id="cardId"><b>Vacation</b></h4>
                                                <div class="d-flex text-center">
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
                            </div>
                        </div>
                    </div>
                </div>
        </div>			
</template>

<script>

export default {
  name: 'LikesAndDislikes',
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
