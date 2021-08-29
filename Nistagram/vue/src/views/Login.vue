<template>
  <div class="register container divs">
        <h3 class="word">Log in</h3><br><br>
      <div class="row">
        <div class="col-md-6">
          <div class="form-group">
            <input id="logusername" type="text" class="form-control" placeholder="Username *" value="" /><br>
            <input id="logpassword" type="password" class="form-control" placeholder="Password *" value="" />
          </div><br>
              <a style="color:red" v-for="e in error" :key="e">
									{{e}}
							</a><br>
          <input type="submit" class="btnRegister"  value="Log in" v-on:click="LogIn()"/><br><br>
          <label><small>Don't have an account?<a href="">Registration</a></small></label>  
      </div>        
      </div>
  </div>
</template>

<script>

const axios=require('axios')

export default {
  name: 'Login',
  components: {
  },
  data: function () {
		return {
			error:[],
			log:false,
            usernameText: null,
			passwordText: null,
			loginResponse: null,
			token: null,
			role:null,
			jwtAuthenticationRequest:{},
			firstlog:null,
            user:null,
        
		}
	},
    beforeMount(){
    },
  methods: {
		LogIn(){
			if(document.getElementById("logusername").value!="" && document.getElementById("logpassword").value!=""){
				//poziv
                this.username = document.getElementById("logusername").value
                this.password = document.getElementById("logpassword").value

                axios
				.get('http://localhost:8080/api/user/login/' + this.username + '/' + this.password)
				.then(response => {
					this.token = response.data
					//localStorage.setItem('userId', response.data.id)
                    alert(this.token)
					if(this.token==undefined){
						alert("Username or password are wrong, please try again for token")
					}
					localStorage.setItem('token', this.token);
					localStorage.setItem('isLogged', true);
					localStorage.setItem('username',this.username)
					//localStorage.setItem('role',response.data.role)
					localStorage.setItem('firstlog',response.data.firstTimeLogin)

                    axios
						.get('http://localhost:8080/api/user/getUserByUsername/'+this.username,
							{
							headers: {
								'Authorization': 'Bearer' + " " + localStorage.getItem('token')
							}
						})
						.then(response => {
							this.user = response.data
							localStorage.setItem('userId', this.user.id)
                            localStorage.setItem('username', this.user.username)
							
						})
						.catch(error => {
                            if(error == "Error: Request failed with status code 400")
                                alert("Username or password are wrong, please try again after this.user!")
                        })
                })
                .catch(error => {
                // print(error.status == 417)
                if(error == "Error: Request failed with status code 417"){
                   alert("Usename and email mast be unique")
                  }
                })

			}else{
				this.error.push('Username/password is incorrect!');
			}
            //alert(localStorage.getItem('userId'))
		}
	}
}
</script>

<style scoped>
.register{
    margin-top: 3%;
    padding: 3%;
    min-height: 750px;
    text-align: center;
}
.divs{
    border-radius: 1.5rem;
    background: -webkit-linear-gradient(left, #eeeef5, #fcfeff);
}
.register-left{
    text-align: center;
    color: #fff;
    margin-top: 4%;
}
.register-left input{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    width: 60%;
    background: #f8f9fa;
    font-weight: bold;
    color: #383d41;
    margin-top: 10%;
    margin-bottom: 3%;
    cursor: pointer;
}
.register-right{
    background: #f8f9fa;
    border-top-left-radius: 10% 50%;
    border-bottom-left-radius: 10% 50%;
}
.register-left img{
    margin-top: 15%;
    margin-bottom: 5%;
    width: 25%;
    -webkit-animation: mover 2s infinite  alternate;
    animation: mover 1s infinite  alternate;
}
@-webkit-keyframes mover {
    0% { transform: translateY(0); }
    100% { transform: translateY(-20px); }
}
@keyframes mover {
    0% { transform: translateY(0); }
    100% { transform: translateY(-20px); }
}
.register-left p{
    font-weight: lighter;
    padding: 12%;
    margin-top: -9%;
}
.register .register-form{
    padding: 5%;
    margin-top: 5%;
}
.register .edit-form{
    padding: 10%;
    margin-top: 10%;
}
.register-form2{
    padding: 10%;
    margin-top: 10%;
    margin-left: 25%;
}
.btnRegister{
    border: none;
    border-radius: 1.5rem;
    padding: 2%;
    background: #55556a;
    color: #fff;
    width: 20%;
    cursor: pointer;
}
.register .nav-tabs{
    margin-top: 3%;
    border: none;
    background: #55556a;
    border-radius: 1.5rem;
    width: 28%;
    float: right;
}
.register .nav-tabs .nav-link{
    padding: 2%;
    height: 34px;
    font-weight: 600;
    color: #fff;
    border-top-right-radius: 1.5rem;
    border-bottom-right-radius: 1.5rem;
}
.register .nav-tabs .nav-link:hover{
    border: none;
}
.register .nav-tabs .nav-link.active{
    width: 100px;
    color: #7559dc;
    border: 2px solid #9679e0;
    border-top-left-radius: 1.5rem;
    border-bottom-left-radius: 1.5rem;
}
.register-heading{
    text-align: center;
    margin-top: 8%;
    margin-bottom: -15%;
    color: #495057;
}

.word{
    color: #55556a;;
    font-weight: bold;
}

.row{
  margin-left: 40%;
}

</style>
